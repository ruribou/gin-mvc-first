package controller

import (
	"fmt"
	"strconv"
	"tutorial/model"

	"github.com/gin-gonic/gin"
)


func ShowAllBlog(c *gin.Context) {
	data := model.GetAll()
	c.HTML(200, "index.html", gin.H{"data": data})
}

func ShowOneBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.HTML(200, "show.html", gin.H{"data": data})
}

func ShowCreateBlog(c *gin.Context) {
	c.HTML(200, "create.html", gin.H{})
}

func CreateBlog(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	data := model.Blog{Title: title, Content: content}
	data.Create()
	c.Redirect(302, "/")
}

func ShowEditBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.HTML(200, "edit.html", gin.H{"data": data})
}

func EditBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	data := model.GetOne(id)
	title := c.PostForm("title")
	data.Title = title
	content := c.PostForm("content")
	data.Content = content
	data.Update()
	c.Redirect(302, "/")
}

func ShowCheckDeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.HTML(200, "delete.html", gin.H{"data": data})
}

func DeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Println("delete:", id)
	data := model.GetOne(id)
	data.Delete()
	c.Redirect(302, "/")
}
