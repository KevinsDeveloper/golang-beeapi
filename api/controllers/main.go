package controllers

import (
    "beego/api/models"
)

// Operations about object
type MainController struct {
    BaseController
}

// @Title Index
// @Description 默认访问页
// @router / [get]
func (this *MainController) GetAll() {
    this.toJson(200, "Success")
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @router /lists [get]
func (this *MainController) PostLists() {
    flt := make([]interface{}, 0)
    filters := append(flt, "Id", 1)

    obs := new(models.AuthUser)
    data, _ := obs.One(filters...)
    this.toJsonData(data, "Success")
}
