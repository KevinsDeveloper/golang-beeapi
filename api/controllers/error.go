package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
    beego.Controller
}

// 返回JSON提示信息
func (this *ErrorController) toJson(ret int, msg string) {
    data := make(map[string]interface{})
    data["ret"] = ret
    data["msg"] = msg

    this.Data["json"] = data
    this.ServeJSON()
    this.StopRun()
}

func (this *ErrorController) Error404() {
    this.toJson(404, "url not found")
}

func (this *ErrorController) Error500() {
    this.toJson(404, "server error")
}
