package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()
    app.Get("/", index)
    app.Listen(":8080")
}

func index(ctx iris.Context) {
    ctx.Writef("ok")
}
