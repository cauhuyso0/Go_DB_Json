package controllers

import (
	"go-module/database"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"body"`
}

//Phương thức đọc dữ liệu từ mysql bằng thư viện gin-gonic
func Read(c *gin.Context) {
	db := database.DBConn()
	rows, err := db.Query("SELECT title, body FROM posts ")
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Story not found",
		})
	}

	post := Post{}

	for rows.Next() {
		//var id int
		var title, body string

		err = rows.Scan(&title, &body)
		if err != nil {
			panic(err.Error())
		}

		//post.Id = id
		post.Title = title
		post.Content = body
		c.JSON(200, post)
		defer db.Close()
	}

}

//Phương thức tạo mới bảng post và insert dữ liệu vào bảng.
func Create(c *gin.Context) {
	db := database.DBConn()

	type CreatePost struct {
		Title string `form:"title" json:"title" binding:"required"`
		Body  string `form:"body" json:"body" binding:"required"`
	}

	var json CreatePost

	if err := c.ShouldBindJSON(&json); err == nil {
		insPost, err := db.Prepare("INSERT INTO posts(title, body) VALUES(?,?)")
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		insPost.Exec(json.Title, json.Body)
		c.JSON(200, gin.H{
			"messages": "inserted",
		})

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer db.Close()
}
