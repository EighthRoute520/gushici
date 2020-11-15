/**********************************************
** @Des: HomeController
** @Author: EighthRoute
** @Date:   2020/10/25 13:13
***********************************************/

package controllers

type HomeController struct {
	BaseController
}

//后台首页
func (self *HomeController) Index() {
	self.Data["pageTitle"] = "系统首页"
	self.TplName = "public/main.html"
}

//后台控制面板
func (self *HomeController) Start() {
	self.Data["pageTitle"] = "控制面板"
	self.display()
}
