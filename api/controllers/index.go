package controllers

import (
    "beego/api/models"
    "github.com/astaxie/beego/logs"
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

    logs.Debug("my book is bought in the year of ", 2016)
    logs.Info("this %s cat is %v years old", "yellow", 3)
    logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
    logs.Error(1024, "is a very", "good game")
    logs.Critical("oh,crash")

    obs := new(models.AuthUser)
    data, _ := obs.One(filters...)
    this.toJsonData(data, "Success")
}
