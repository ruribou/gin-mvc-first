package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.GET("/", ShowAllBlog)
	r.GET("/show/:id", ShowOneBlog)
	r.GET("/create", ShowCreateBlog)
	r.POST("/create", CreateBlog)
	r.GET("/edit/:id", ShowEditBlog)
	r.POST("/edit", EditBlog)
	r.GET("/delete/:id", ShowCheckDeleteBlog)
	r.POST("/delete", DeleteBlog)
	return r
}