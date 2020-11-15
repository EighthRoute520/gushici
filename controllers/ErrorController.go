package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (self *ErrorController) Error404() {
	self.Data["content"] = "page not found"
	self.TplName = "error/404.html"
}

func (self *ErrorController) Error501() {
	self.Data["content"] = "server error"
	self.TplName = "error/501.html"
}

func (self *ErrorController) ErrorDb() {
	self.Data["content"] = "database is now down"
	self.TplName = "error/dberror.html"
}
