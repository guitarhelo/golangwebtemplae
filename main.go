// golangidev project main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangidev/domain"
	"golangidev/web"
	"mime"
	"net/http"
)

type indexInfo struct {
	Title   string
	Version string
}
type pageInfo struct {
	Result       []domain.User
	Title        string
	TotalRecords int
}

//var userDao *respository.UserRepoImpl = new(respository.UserRepoImpl)
//var userService *service.UserServiceImpl = new(service.UserServiceImpl)

func main() {
	mime.AddExtensionType(".css", "text/css")
	fmt.Println("Hello World!")
	route := gin.Default()

	route.Static("/css", "./assets/css")
	route.Static("/img", "./assets/img")
	route.Static("/js", "./assets/js")
	route.Static("/plugins", "./assets/plugins")
	route.Static("/fonts", "./assets/fonts")
	//route.Static("/font-awesome", "./assets/font-awesome")

	//route.StaticFS("/", http.Dir("./view"))
	route.LoadHTMLGlob("view/*")

	route.GET("/", func(c *gin.Context) {

		obj := indexInfo{
			Title:   "Golang Bootstrap Platform",
			Version: "1.0",
		}
		obj.Version = "2.0"
		//c.Header("Content-Type", "application/x-css")

		//obj := gin.H{"title": "Main website"}
		c.HTML(http.StatusOK, "index.html", obj)

	})

	web.Register(route)

	route.Run(":8080") // listen and serve on 0.0.0.0:8080
}
