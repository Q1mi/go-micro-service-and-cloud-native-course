package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 基于游标的分页

func bookListHandler(c *gin.Context) {
	pageToken := c.Query("page_token") // 从请求中获取分页令牌
	// 解token
	page := Token(pageToken).Decode()
	// 分页校验
	if page.NextID == "" {
		c.JSON(401, "bad page_token")
		return
	}
	if page.NextTimeAtUTC > time.Now().Unix() || time.Now().Unix()-page.NextTimeAtUTC > int64(time.Hour)*24 {
		c.JSON(401, "bad page_token")
		return
	}

	sql := `select id, title from books where id > ? order by id asc limit ?` // page.NextID page.PageSize

	// 去数据库查询数据
	data := db.Query(sql)

	// 拿到最后一条数据，拼接下一页的page_token
	nextPage := Page{
		NextID:        "20",
		NextTimeAtUTC: time.Now().Unix(),
		PageSize:      page.PageSize,
	}

	nextPageToken := nextPage.Encode()
	c.JSON(200, gin.H{
		"data":       data,
		"page_token": nextPageToken,
	})
}

func main() {
	r := gin.Default()
	r.GET("/api/v1/books", bookListHandler)
	r.Run()
}
