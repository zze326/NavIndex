package main

import (
	"embed"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

//go:embed static/*
var staticDir embed.FS

//go:embed views
var viewsDir embed.FS

type Page struct {
	Title            string   `yaml:"title"`
	MainTitle        string   `yaml:"mainTitle"`
	InputPlaceholder string   `yaml:"inputPlaceholder"`
	Favicon          string   `yaml:"favicon"`
	Groups           []*Group `yaml:"groups"`
}

type Group struct {
	Group string  `yaml:"group"`
	Items []*Item `yaml:"items"`
}

type Item struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

func index(ctx iris.Context) {
	if err := ctx.View("index", page); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

var page = new(Page)

func main() {
	dataBytes, err := os.ReadFile("./data.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(dataBytes, page); err != nil {
		log.Fatal(err)
	}

	app := iris.New()
	htmlEngine := iris.HTML(viewsDir, ".html")
	htmlEngine.RootDir("views")
	htmlEngine.AddFunc("ext", func(val string) string {
		return path.Ext(val)
	})
	app.RegisterView(htmlEngine)
	app.HandleDir("/static", staticDir)
	app.Favicon(page.Favicon)
	app.Get("/", index)

	app.Listen("0.0.0.0:8080")
}
