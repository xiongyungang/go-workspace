package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "xyg!")
		ctx.View("hello.html")
	})
	app.Run(iris.Addr(":8088"))
}
