package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Blog struct {
	ID         int
	Title      string
	Content    string
	Image      string
	Author     string
	PostDate   time.Time
	FormatDate string
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

func main() {
	// Connect to database
	connection.DatabaseConnect()

	// Create new Echo instance
	e := echo.New()

	// Middleware, logger for logging, recover is handling when it's panic
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Serve static files from "/public" directory
	e.Static("/public", "public")

	// To use sessions using echo
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// Routing
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/form-blog", formAddBlog)
	e.POST("/add-blog", addBlog)
	e.POST("/blog-delete/:id", deleteBlog)

	e.GET("/form-register", formRegister)
	e.POST("/register", register)

	e.GET("/form-login", formLogin)
	e.POST("/login", login)

	e.POST("/logout", logout)

	// Start server
	println("Server running on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func home(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, content, image, post_date FROM tb_blog ORDER BY id DESC")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Content, &each.Image, &each.PostDate)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		each.FormatDate = each.PostDate.Format("2 January 2006")
		each.Author = "Surya Elidanto"

		result = append(result, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	blogs := map[string]interface{}{
		"Blogs":       result,
		"DataSession": userData,
	}

	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), blogs)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, content, image, post_date FROM tb_blog WHERE id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Title, &BlogDetail.Content, &BlogDetail.Image, &BlogDetail.PostDate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	BlogDetail.FormatDate = BlogDetail.PostDate.Format("2 January 2006")
	BlogDetail.Author = "Surya Elidanto"

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/blog-detail.html")

	if errTemplate != nil {
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

	image := "astronaut.jpg"

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (title, content, image, post_date) VALUES ($1, $2, $3, $4)", title, content, image, time.Now())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func formRegister(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func register(c echo.Context) error {
	// parseForm is best practice to make sure request body is form data format, not JSON, XML, etc.
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("inputName")
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success!", true, "/form-login")
}

func formLogin(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/form-login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email or password incorrect!", false, "/form-login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Email or password incorrect!", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}
