package controllers

import (
    "github.com/astaxie/beego"
    "strings"
)

type BaseController struct {
    beego.Controller

    controllerName string
    actionName     string
    layouts        string
    tplsuffix      string
    auth           interface{}
    auth_id        int
    pagesize       int
}

// Method 方法之前执行
func (self *BaseController) Prepare() {
    self.pagesize = 12
    self.layouts = "layout/layout"
    self.tplsuffix = "tpl"

    controllerName, actionName := self.GetControllerAndAction()
    self.controllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
    self.actionName = strings.ToLower(actionName)

    self.authValidate()
    self.Data["auth"] = self.auth
    self.Data["auth_id"] = self.auth_id
}

// 登录权限校验
func (self *BaseController) authValidate() {
    //arr := strings.Split(self.Ctx.GetCookie("auth"), "|")
}

// 判断是否为POST提交
func (self *BaseController) isPost() bool {
    return self.Ctx.Request.Method == "POST"
}

// 获取用户请求IP
func (self *BaseController) clientIP() string {
    s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
    return s[0]
}

// 重定向
func (self *BaseController) redirect(url string) {
    self.Redirect(url, 302)
    self.StopRun()
}

// 自动加载模板
func (self *BaseController) display(tpl ...string) {
    tplname := ""
    if len(tpl) > 0 {
        tplname = strings.Join([]string{tpl[0], self.tplsuffix}, ".")
    } else {
        tplname = self.controllerName + "/" + self.actionName + "." + self.tplsuffix
    }
    self.Layout = self.layouts + "." + self.tplsuffix
    self.TplName = tplname
}

// 返回JSON提示信息
func (self *BaseController) toJson(ret int, msg interface{}) {
    data := make(map[string]interface{})
    data["ret"] = ret
    data["msg"] = msg

    self.Data["json"] = data
    self.ServeJSON()
    self.StopRun()
}

// 返回JSON数据
func (self *BaseController) toJsonData(datas interface{}, msg interface{}) {
    data := make(map[string]interface{})
    data["ret"] = 200
    data["msg"] = msg
    data["data"] = datas

    self.Data["json"] = data
    self.ServeJSON()
    self.StopRun()
}
