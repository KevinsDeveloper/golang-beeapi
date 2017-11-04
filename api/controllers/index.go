package controllers

import (
    "beego/api/models"
)

// Operations about object
type IndexController struct {
    BaseController
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (this *IndexController) GetAll() {
    flt := make([]interface{}, 0)
    filters := append(flt, "Id", 1)

    obs := new(models.AuthUser)
    data, _ := obs.One(filters...)
    this.toJsonData(data, "success")
}
