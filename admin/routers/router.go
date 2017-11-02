package routers

import (
    "beego/admin/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
