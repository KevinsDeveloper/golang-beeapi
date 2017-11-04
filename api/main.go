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
    beego.ErrorController(&controllers.ErrorController{})
    models.Init()
    beego.Run()
}
