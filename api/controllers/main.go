package controllers

import (
    "beego/api/models"
    "beego/api/libs"
    "fmt"
)

// Operations about object
type MainController struct {
    BaseController
}

// @Title Index
// @Description 默认访问页
// @router / [get]
func (this *MainController) GetAll() {
    //logs.Error
    rc := libs.Redis().Set("test", "aaa", 0).Err()
    fmt.Println(rc)
    this.toJson(100011, "Success", 500)
}

// @Title 获取列表
// @Description 获取管理员列表
// @Success 200 {object} models.Object
// @router /lists [get]
func (this *MainController) PostLists() {
    flt := make([]interface{}, 0)
    filters := append(flt, "Id", 1)
    filters = append(filters, "RoleId", 1)

    obs := new(models.AuthUser)
    data, _ := obs.One(filters...)
    this.toJsonData(data, "Success")
}
