package controllers

import "time"

type MainController struct {
    BaseController
}

func (self *MainController) Get() {
    self.Data["time"] = time.Now().Unix()
    self.display()
}
