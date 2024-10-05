package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	bin, err := os.ReadFile("/run/secrets/db-password")
	if err != nil {
		return nil, err
	}
	return sql.Open("mysql", fmt.Sprintf("root:%s@tcp(db:3306)/example", string(bin)))
}

func blogHandler(c *gin.Context) {
	db, err := connect()
	if err != nil {
		c.JSON(500, "connect db fail")
	}
	defer db.Close()

	rows, err := db.Query("SELECT title FROM blog")
	if err != nil {
		c.JSON(500, "query db fail")
	}
	var titles []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		titles = append(titles, title)
	}
	c.JSON(200, titles)
}

func main() {
	log.Print("Prepare db...")
	if err := prepare(); err != nil {
		log.Fatal(err)
	}

	log.Print("Listening 8080")

	r := gin.Default()
	r.GET("/", blogHandler)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func prepare() error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	for i := 0; i < 60; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if _, err := db.Exec("DROP TABLE IF EXISTS blog"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS blog (id int NOT NULL AUTO_INCREMENT, title varchar(255), PRIMARY KEY (id))"); err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		if _, err := db.Exec("INSERT INTO blog (title) VALUES (?);", fmt.Sprintf("Q1mi's Blog post #%d", i)); err != nil {
			return err
		}
	}
	return nil
}
