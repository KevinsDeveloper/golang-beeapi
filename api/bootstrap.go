package main

import "github.com/astaxie/beego/logs"

func Bootstrap() {
    logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/app.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
    logs.Async()
}
