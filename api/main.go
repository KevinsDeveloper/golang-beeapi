package main

import (
    _ "beego/api/routers"
    "github.com/astaxie/beego"
    "beego/api/controllers"
    "beego/api/models"
)

func main() {
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }
    Bootstrap()
    models.Init()
    beego.ErrorController(&controllers.ErrorController{})
    beego.Run()
}
