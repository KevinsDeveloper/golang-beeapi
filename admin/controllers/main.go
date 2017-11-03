package controllers

import (
    "time"
    "fmt"
)

type MainController struct {
    BaseController
}

func (self *MainController) Get() {
    numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Println(numbers)

    self.Data["time"] = time.Now().Unix()
    self.display()
}
