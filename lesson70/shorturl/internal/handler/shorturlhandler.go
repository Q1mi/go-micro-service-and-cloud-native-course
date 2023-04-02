package handler

import (
	"net/http"

	"shorturl/internal/logic"
	"shorturl/internal/svc"
	"shorturl/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShorturlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShorturlLogic(r.Context(), svcCtx)
		resp, err := l.Shorturl(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// 之前是返回响应数据
			// httpx.OkJsonCtx(r.Context(), w, resp)

			// 返回重定向的HTTP响应
			// w.Header().Set("location", resp.LongURL) // location
			// w.WriteHeader(http.StatusFound)          // status code

			// 另外的一种方式，使用http包里的重定向方法
			http.Redirect(w, r, resp.LongURL, http.StatusFound)
		}
	}
}
