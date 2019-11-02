package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	
	timeString := time.Now().Format("2006-01-02 15:04:05")
	c.Ctx.WriteString("hello world "+timeString)
	fmt.Println("hello "+timeString)
}
