package controllers

import (
    "time"
    "fmt"
)

type MainController struct {
    BaseController
}

func (self *MainController) Get() {
    numbers := []string{}
    fmt.Println(numbers)

    self.Data["time"] = time.Now().Unix()
    self.display()
}
