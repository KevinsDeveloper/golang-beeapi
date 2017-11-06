package main

import (
    _ "beego/api/routers"
    "github.com/astaxie/beego"
    "beego/api/libs"
)

func main() {
    Bootstrap()
    libs.Init()
    beego.Run()
}
