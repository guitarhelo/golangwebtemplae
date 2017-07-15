package web

import (
	"bytes"
	"fmt"
	"golangidev/domain"
	"golangidev/service"
	"html/template"
	"log"
	"math"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var userService *service.UserServiceImpl = new(service.UserServiceImpl)

type loginForum struct {
	User    domain.User
	Message string
}
type greetings struct {
	Intro    string
	Messages []string
}
type pagejumplink struct {
	PageNum int
	PageUrl string
}
type pageinfo struct {
	Users        []domain.User
	Total        int
	Current_Page int
	PerPageNum   int
	TotalPages   int
	PageNavgator template.HTML
	PageJumpLink []pagejumplink
	IsFirstPage  bool
	IsLastPage   bool
	Next_page    int
	Last_page    int
}

func StartUp() {
	log.Println("start...Register User Controller route")
	//Register(router * gin.Engine)

}
func Register(router *gin.Engine) {

	//router.Static("/assets", "./assets")

	//router.LoadHTMLGlob("html/*")

	router.GET("/user", userIndexHandler)
	router.GET("/user/:id", userShowHandler)
	router.GET("/form", newFormHandler)
	router.POST("/create", userCreateHandler)
	router.POST("/update", userUpdateHandler)
	router.POST("/search/:page", userSearchHandler)
	router.GET("/list/:page", userListHandler)

	router.GET("/edit/:id", userEditHandler)
	router.GET("/test/ping", userTestPingHandler)
	router.GET("/users.html", userUserListHandler)
	router.GET("/test/test", userTestCmdHandler)

}
func userIndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello world,user")
}
func newFormHandler(c *gin.Context) {
	//c.String(http.StatusOK, "New Form")
	model := loginForum{
		User:    domain.User{},
		Message: "create user form",
	}

	c.HTML(http.StatusOK, "form.html", model)

}
func userCreateHandler(c *gin.Context) {

	c.Request.ParseForm()
	name := c.Request.Form.Get("name")
	password := c.Request.Form.Get("password")
	address := c.Request.Form.Get("address")
	age := c.Request.Form.Get("age")
	var myage int
	var user_save_result int64
	myage, err := strconv.Atoi(age)
	if err != nil {
		panic(err.Error())
	}
	//age1 := strconv.Atoi("33")
	createtime := time.Now().Format("2006-01-02 15:04:05")
	user := domain.User{
		Name:       name,
		Password:   password,
		Age:        myage,
		Address:    address,
		CreateTime: createtime,
	}

	user_save_result = userService.Save(user)
	fmt.Println("================begin add_user")
	fmt.Println(user_save_result)
	fmt.Println("===============end add_user")

	if user_save_result > 0 {
		Intro := "create user success"
		c.String(http.StatusOK, Intro)

	} else {
		Intro := "create user failed"
		c.String(http.StatusOK, Intro)
	}

}
func userUpdateHandler(c *gin.Context) {

	c.Request.ParseForm()
	name := c.Request.Form.Get("name")
	password := c.Request.Form.Get("password")
	address := c.Request.Form.Get("address")
	age := c.Request.Form.Get("age")
	id := c.Request.Form.Get("user_id")
	fmt.Println("id===========================begin")
	fmt.Println(id)
	fmt.Println("id===========================end")
	var myage int
	var user_update_result int64
	myage, err := strconv.Atoi(age)
	if err != nil {
		panic(err.Error())
	}
	my_id, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	//age1 := strconv.Atoi("33")
	createtime := time.Now().Format("2006-01-02 15:04:05")
	user := domain.User{
		Id:         my_id,
		Name:       name,
		Password:   password,
		Age:        myage,
		Address:    address,
		CreateTime: createtime,
	}

	user_update_result = userService.Update(user)
	fmt.Println("================begin update_user")
	fmt.Println(user)
	fmt.Printf(id)
	fmt.Println(user_update_result)
	fmt.Println("===============end update_user")

	if user_update_result > 0 {
		Intro := "update user success"
		c.String(http.StatusOK, Intro)
	} else {
		Intro := "update user failed"
		c.String(http.StatusOK, Intro)
	}

}

func userShowHandler(c *gin.Context) {

	model := loginForum{
		User:    domain.User{},
		Message: "show user info",
	}
	temp_id := c.Params.ByName("id")
	var u_id int
	u_id, err := strconv.Atoi(temp_id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	model.User = userService.GetUserById(u_id)
	c.HTML(http.StatusOK, "show.html", model)
	//c.String(http.StatusOK, "show user ")
}
func userListHandler(c *gin.Context) {

	//passedObj := greetings{
	//	Intro:    "Hello from Go!",
	//	Messages: []string{"Hello!", "Hi!", "潘敬平", "¡Hola!", "Bonjour!", "Ciao!", "<script>evilScript()</script>"},
	//}
	var cpage int
	var totalUsers int
	var perPageNum = 4
	var maxpage int
	totalUsers = userService.GetTotalUsers("select count(*) as count from userinfo")

	if math.Mod(float64(totalUsers), float64(perPageNum)) == 0 {
		maxpage = totalUsers / perPageNum
	} else {

		maxpage = (totalUsers / perPageNum) + 1
	}
	//c.Request.ParseForm()
	passedObj := pageinfo{
		Total:      totalUsers,
		PerPageNum: perPageNum,
		TotalPages: maxpage,
	}
	//current_page := c.Request.Form.Get("current_page")
	current_page := c.Params.ByName("page")
	var temp int
	temp, err := strconv.Atoi(current_page)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if temp > maxpage {
		cpage = maxpage
	} else {
		cpage = temp
	}
	passedObj.Current_Page = cpage

	passedObj.Next_page = cpage + 1
	passedObj.Last_page = cpage - 1
	fmt.Println("===========================crrrent paee")
	fmt.Println(cpage)
	fmt.Println("===========================pass object page")
	fmt.Println(passedObj.Current_Page)
	if err != nil {
		panic(err.Error())
	}
	for i := 1; i <= maxpage; i++ {
		pagelinkobj := pagejumplink{
			PageNum: i,
			PageUrl: "/list",
		}
		passedObj.PageJumpLink = append(passedObj.PageJumpLink, pagelinkobj)
	}

	switch {
	case (cpage > 1 && cpage < maxpage):
		{
			passedObj.PageNavgator = "<a href=''>last page </a> | <a href=''>next page </a>"
			passedObj.IsFirstPage = false
			passedObj.IsLastPage = false

		}
	case cpage == maxpage:
		{
			if cpage != 1 {
				passedObj.PageNavgator = "<a href=''>last page </a>"
				passedObj.IsFirstPage = false
				passedObj.IsLastPage = true
			} else {
				passedObj.PageNavgator = ""
				passedObj.IsFirstPage = true
				passedObj.IsLastPage = true
			}
		}

	}
	//passedObj.Users = service.GetAllUserList()

	fmt.Println("=================edd==========pass object page")
	fmt.Println("====================ddd======================")
	fmt.Println(passedObj.IsFirstPage)
	fmt.Println(passedObj.IsLastPage)
	fmt.Println("=====================ddd=====================")
	fmt.Println(passedObj.Current_Page)
	//passedObj.Users = service.GetUserList(cpage-1, perPageNum)
	passedObj.Users = userService.GetTotalUsersByPaging(cpage-1, perPageNum)
	c.HTML(http.StatusOK, "list.html", passedObj)
	//c.String(http.StatusOK, "List Users")

}

func userSearchHandler(c *gin.Context) {

	//passedObj := greetings{
	//	Intro:    "Hello from Go!",
	//	Messages: []string{"Hello!", "Hi!", "潘敬平", "¡Hola!", "Bonjour!", "Ciao!", "<script>evilScript()</script>"},
	//}
	var cpage int
	var totalUsers int
	var perPageNum = 4
	var maxpage int

	totalUsers = userService.GetTotalUsers("select count(*) as count from userinfo where 1 and  (name like '%pan%' or address like '%demo%')  and ENABLEd=1 and  age BETWEEN 1 and 46")

	if math.Mod(float64(totalUsers), float64(perPageNum)) == 0 {
		maxpage = totalUsers / perPageNum
	} else {

		maxpage = (totalUsers / perPageNum) + 1
	}
	//c.Request.ParseForm()
	passedObj := pageinfo{
		Total:      totalUsers,
		PerPageNum: perPageNum,
		TotalPages: maxpage,
	}
	//current_page := c.Request.Form.Get("current_page")
	current_page := c.Params.ByName("page")

	var temp int
	temp, err := strconv.Atoi(current_page)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if temp > maxpage {
		cpage = maxpage
	} else {
		cpage = temp
	}
	passedObj.Current_Page = cpage

	passedObj.Next_page = cpage + 1
	passedObj.Last_page = cpage - 1
	fmt.Println("===========================crrrent paee")
	fmt.Println(cpage)
	fmt.Println("===========================pass object page")
	fmt.Println(passedObj.Current_Page)
	if err != nil {
		panic(err.Error())
	}
	for i := 1; i <= maxpage; i++ {
		pagelinkobj := pagejumplink{
			PageNum: i,
			PageUrl: "/search",
		}
		passedObj.PageJumpLink = append(passedObj.PageJumpLink, pagelinkobj)
	}

	switch {
	case (cpage > 1 && cpage < maxpage):
		{
			passedObj.PageNavgator = "<a href=''>last page </a> | <a href=''>next page </a>"
			passedObj.IsFirstPage = false
			passedObj.IsLastPage = false

		}
	case cpage == maxpage:
		{

			passedObj.PageNavgator = "<a href=''>last page </a>"
			passedObj.IsFirstPage = false
			passedObj.IsLastPage = true

		}
	case cpage == 1:
		{

			passedObj.PageNavgator = "<a href='/list?current_page=1'>next page </a>"
			passedObj.IsFirstPage = true
			passedObj.IsLastPage = false

		}
	default:
		{
			passedObj.PageNavgator = "<a href=''>next page </a>"
			passedObj.IsFirstPage = true
			passedObj.IsLastPage = false

		}

	}
	//passedObj.Users = service.GetAllUserList()

	fmt.Println("=================edd==========pass object page")
	fmt.Println("====================ddd======================")
	fmt.Println(passedObj.IsFirstPage)
	fmt.Println(passedObj.IsLastPage)
	fmt.Println("=====================ddd=====================")
	fmt.Println(passedObj.Current_Page)
	//passedObj.Users = service.GetUserList(cpage-1, perPageNum)
	passedObj.Users = userService.GetTotalUsersByPaging(cpage-1, perPageNum)
	passedObj.Total = len(passedObj.Users)
	c.HTML(http.StatusOK, "search_result.html", passedObj)
	//c.String(http.StatusOK, "List Users")

}
func userEditHandler(c *gin.Context) {
	var passedObj domain.User

	id := c.Params.ByName("id")

	userId, err := strconv.Atoi(id)
	fmt.Println("User id=========%d", userId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app

	}
	passedObj = userService.GetUserById(userId)
	c.HTML(http.StatusOK, "edit.html", passedObj)
	fmt.Println(passedObj)

}
func userTestPingHandler(c *gin.Context) {

	c.String(200, "pong")
}
func userUserListHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "users.html", "demo")
}

func userTestCmdHandler(c *gin.Context) {
	cmd := exec.Command("cmd", "/C", "dir", "c:\\")
	var out bytes.Buffer //缓冲字节
	cmd.Stdout = &out    //标准输出

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	c.String(200, out.String())
}
