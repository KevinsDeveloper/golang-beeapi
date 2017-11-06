package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
    beego.Controller
}

// 返回JSON提示信息
func (this *ErrorController) toJson(ret int, msg string, types int) {
    data := make(map[string]interface{})
    data["errType"] = types
    data["retCode"] = ret
    data["retMsg"] = msg

    this.Data["json"] = data
    this.ServeJSON()
    this.StopRun()
}

// 404错误
func (this *ErrorController) Error404() {
    this.toJson(404, "url not found", 404)
}

// 500错误
func (this *ErrorController) Error500() {
    this.toJson(404, "server error", 500)
}
