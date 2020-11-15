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
func (this *HomeController) Index() {
	this.Data["pageTitle"] = "系统首页"
	this.TplName = "public/main.html"
}

//后台控制面板
func (this *HomeController) Start() {
	this.Data["pageTitle"] = "控制面板"
	this.display()
}
