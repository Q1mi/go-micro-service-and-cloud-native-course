package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

// singleflight demo

// getData 取数据函数
func getData(id int64) string {
	fmt.Println("正在查询中...")
	time.Sleep(10 * time.Second) // 模拟一个比较耗时的操作
	return "liwenzhou.com"
}

func doGetData(g *singleflight.Group, id int64) (string, error) {
	v, err, _ := g.Do("getData", func() (interface{}, error) {
		ret := getData(id)
		return ret, nil
	})
	return v.(string), err
}

func doChanGetData(ctx context.Context, g *singleflight.Group, id int64) (string, error) {
	ch := g.DoChan("getData", func() (interface{}, error) {
		ret := getData(id)
		return ret, nil
	})
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case ret := <-ch:
		return ret.Val.(string), ret.Err
	}
}

func main() {

	g := new(singleflight.Group)

	// 第1次调用
	go func() {
		// v, err := doGetData(g, 1)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		v, err := doChanGetData(ctx, g, 1)
		fmt.Printf("1st call: v:%v err:%v \n", v, err)
	}()

	time.Sleep(2 * time.Second) // 确保上面第1次调用先执行

	// 第2次调用
	// v, err, shared := g.Do("getData", func() (interface{}, error) {
	// 	ret := getData(1)
	// 	return ret, nil
	// })
	// fmt.Printf("2nd call: v:%v err:%v shared:%v\n", v, err, shared)

	// do版本
	// v, err := doGetData(g, 1)

	// doChan 版本
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	v, err := doChanGetData(ctx, g, 1)
	fmt.Printf("2nd call: v:%v err:%v \n", v, err)

	// 第3次调用。。。
	// g.Do("getData", func() (interface{}, error) {
	// 	ret := getData(1)
	// 	return ret, nil
	// })

}
