package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	Title    string
	Content  string
	Author   string
	PostDate string
}

var dataBlog = []Blog{
	{
		Title:    "Hallo Title",
		Content:  "Hallo Content",
		Author:   "Surya Elidanto",
		PostDate: "07/02/2023",
	},
	{
		Title:    "Hallo Title 2",
		Content:  "Hallo Content 2",
		Author:   "Surya Elidanto",
		PostDate: "08/02/2023",
	},
}

func main() {
	// Create new Echo instance
	e := echo.New()

	// Middleware, logger for logging, recover is handling when it's panic
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Serve static files from "/public" directory
	e.Static("/public", "public")

	// Routing
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/form-blog", formAddBlog)
	e.POST("/blog-delete/:id", deleteBlog)
	e.POST("/add-blog", addBlog)

	// Start server
	println("Server running on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogs := map[string]interface{}{
		"Blogs": dataBlog,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	for i, data := range dataBlog {
		if id == i {
			BlogDetail = Blog{
				Title:    data.Title,
				Content:  data.Content,
				PostDate: data.PostDate,
				Author:   data.Author,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddBlog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("inputTitle")
	content := c.FormValue("inputContent")

	println("Title : " + title)
	println("Content : " + content)

	var newBlog = Blog{
		Title:    title,
		Content:  content,
		Author:   "Surya Elidanto",
		PostDate: time.Now().String(),
	}

	dataBlog = append(dataBlog, newBlog)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
