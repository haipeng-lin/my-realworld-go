package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"         // Gin框架
	_ "github.com/go-sql-driver/mysql" // 加载MySQL驱动
)

// 常量
const (
	dbUsername = "root"
	dbPassword = "haipeng123"
	dbHost     = "localhost"
	dbPort     = 3306
	dbName     = "my-realworld"
)

type user struct {
	Id       int    `json:"id"`
	Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
	Email    string `form:"email" json:"email" binding:"exists,email"`
	Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	Bio      string `form:"bio" json:"bio" binding:"max=1024"`
	Image    string `form:"image" json:"image" binding:"omitempty,url"`
}

// 连接数据库
func connectToDatabase() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := connectToDatabase() // 连接数据库
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default() // gin路由
	r.GET("/user/:id", func(c *gin.Context) {
		var (
			user   user  // 用户结构体
			result gin.H // Map集合类型
		)
		id := c.Param("id")                                                                     // 获取参数id
		row := db.QueryRow("SELECT id, username, email, image, bio FROM user WHERE id = ?", id) // 执行 SQL 查询
		err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Image, &user.Bio)           // 将行数据扫描到user 结构体内
		if err != nil {
			result = gin.H{
				"message": "没找到该用户",
			}
		} else {
			result = gin.H{
				"data": user, //返回查询结果
			}
		}
		c.JSON(200, result)
	})
	r.Run(":8080")
}
