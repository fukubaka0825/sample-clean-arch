package main

import (
	"fmt"
	"io"
	"html/template"
	"sample-clean-arch/const"
	"sample-clean-arch/infrastructure/web/validater"
	"sample-clean-arch/infrastructure/datastore"
	"sample-clean-arch/registry"
	"sample-clean-arch/conf"
	"sample-clean-arch/infrastructure/web/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// レイアウト適用済のテンプレートを保存するmap
var templates map[string]*template.Template

// Template はHTMLテンプレートを利用するためのRenderer Interfaceです。
type Template struct {
	templates *template.Template
}

// Render はHTMLテンプレートにデータを埋め込んだ結果をWriterに書き込みます。
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return templates[name].ExecuteTemplate(w, "layout.html", data)
}

// 初期化を行います。
func init() {
	loadTemplates()
	dao.MigrateMembers()
	//dao.CreateMember()
	dao.MigratePosts()
}

// 各HTMLテンプレートに共通レイアウトを適用した結果を保存します（初期化時に実行）。
func loadTemplates() {
	templates = make(map[string]*template.Template)
	templates["update_post_form"] = template.Must(
		template.ParseFiles(_const.BaseTemplate, _const.UpdataPost))
	templates["posts_all"] = template.Must(
		template.ParseFiles(_const.BaseTemplate, _const.PostsAll))
}

func MethodOverride(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			method := c.Request().PostFormValue("_method")
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				c.Request().Method = method
			}
		}
		//fmt.Println(c.Request())
		return next(c)
	}
}

func main() {
	// conf読み取り
	conf.ReadConf()

	// DB取得
	conn := datastore.NewMySqlDB()

	// interactor
	r := registry.NewInteractor(conn)

	// 依存解決
	h := r.NewAppHandler()

	// echo
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// テンプレートを利用するためのRendererの設定
	t := &Template{
		templates: template.Must(template.ParseGlob(_const.BaseTemplate)),
	}
	e.Renderer = t

	// validate
	e.Validator = &validater.CustomValidator{Validator:validator.New()}

	// router
	router.NewRouter(e, h)

	// DB stop
	defer func() {
		if err := conn.Close(); err != nil {
			e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()

	// server start
	e.Logger.Fatal(e.Start(":" + conf.C.Server.Address))
}
