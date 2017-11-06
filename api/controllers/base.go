package controllers

import (
    "github.com/astaxie/beego"
    "strings"
)

type BaseController struct {
    beego.Controller

    controllerName string
    actionName     string
    auth           interface{}
    auth_id        int
}

// Method 方法之前执行
func (this *BaseController) Prepare() {
    controllerName, actionName := this.GetControllerAndAction()
    this.controllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
    this.actionName = strings.ToLower(actionName)

    this.authValidate()
    this.Data["auth"] = this.auth
    this.Data["auth_id"] = this.auth_id
}

// 登录权限校验
func (this *BaseController) authValidate() {
    //arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
}

// 返回JSON提示信息
func (this *BaseController) toJson(code int, msg string, types int) {
    data := make(map[string]interface{})
    data["errType"] = types
    data["retCode"] = code
    data["retMsg"] = msg

    this.Data["json"] = data
    this.ServeJSON()
    this.StopRun()
}

// 返回JSON数据
func (this *BaseController) toJsonData(datas interface{}, msg string) {
    data := make(map[string]interface{})
    data["retCode"] = 0
    data["retMsg"] = msg
    data["retData"] = datas

    this.Data["json"] = data
    this.ServeJSON()
}