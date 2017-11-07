package main

import (
    _ "beego/api/routers"
    "github.com/astaxie/beego"
    "beego/api/models"
)

func main() {
    Bootstrap()
    models.Init()
    beego.Run()
}
