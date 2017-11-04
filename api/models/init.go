package models

import (
    "net/url"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func Init() {
    host := beego.AppConfig.String("db.host")
    port := beego.AppConfig.String("db.port")
    user := beego.AppConfig.String("db.user")
    password := beego.AppConfig.String("db.password")
    name := beego.AppConfig.String("db.name")
    timezone := beego.AppConfig.String("db.timezone")
    charset := beego.AppConfig.String("db.charset")

    dns := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=" + charset
    if timezone != "" {
        dns += "&loc=" + url.QueryEscape(timezone)
    }

    // set default database
    orm.RegisterDataBase("default", "mysql", dns, 30)

    // register model
    orm.RegisterModel(new(AuthUser))

    // create table
    orm.RunSyncdb("default", false, true)

    // debug
    if beego.AppConfig.String("runmode") == "dev" {
        orm.Debug = true
    }

}

func TableName(name string) string {
    return beego.AppConfig.String("db.prefix") + name
}
