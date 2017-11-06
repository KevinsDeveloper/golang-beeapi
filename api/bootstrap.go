package main

import (
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
    "beego/api/controllers"
)

func Bootstrap() {
    // 开发模式
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }
    // 日志设置
    logs.SetLogger(logs.AdapterMultiFile, `{
            "filename":"logs/app.log",
            "separate":[
                "emergency",
                "alert",
                "critical",
                "error",
                "warning",
                "notice",
                "info",
                "debug"
            ]
        }`)
    logs.Async()
    
    // 错误处理控制器
    beego.ErrorController(&controllers.ErrorController{})
}
