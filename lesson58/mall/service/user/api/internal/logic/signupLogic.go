package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"mall/service/user/model"
	"time"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var secret = []byte("夏天夏天悄悄过去")

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignupLogic) Signup(req *types.SignupRequest) (resp *types.SignupResponse, err error) {
	// 参数校验
	if req.RePassword != req.Password {
		return nil, errors.New("两次输入的密码不一致")
	}
	// 你的业务逻辑写在这里
	fmt.Printf("req:%#v\n", req)
	// 把用户的注册信息保存到数据库中
	// 0. 查询username是否已经被注册
	// https://github.com/go-sql-driver/mysql#system-variables
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	// 0.1 查询数据库失败了
	if err != nil && err != sqlx.ErrNotFound {
		fmt.Printf("FindOneByUsername err:%v\n", err)
		return nil, errors.New("内部错误")
	}
	// 0.2 查到记录,表示该用户名已经被注册
	if u != nil {
		return nil, errors.New("用户名已存在")
	}
	// 0.3 没查到记录
	// 1. 生成userId（雪花算法）
	// 2. 加密密码（加盐 md5）
	h := md5.New()
	h.Write([]byte(req.Password)) // 密码计算md5
	h.Write(secret)               // 加盐
	passwordStr := hex.EncodeToString(h.Sum(nil))

	user := &model.User{
		UserId:   time.Now().Unix(), // 这里简化，后面再讲雪花算法生成userid
		Username: req.Username,
		Password: passwordStr, // 不能存明文
		Gender:   int64(req.Gender),
	}
	if _, err := l.svcCtx.UserModel.Insert(context.Background(), user); err != nil {
		return nil, err
	}
	return &types.SignupResponse{Message: "success"}, nil
}
