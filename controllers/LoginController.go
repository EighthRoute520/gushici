/**********************************************
** @Des: LoginController
** @Author: EighthRoute
** @Date:   2020/10/25 10:34
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/libs"
	"gushici/servers"
	"strconv"
	"strings"
	"time"
)

type LoginController struct {
	BaseController
}

//后台登录
func (this *LoginController) Login() {
	adminServer := new(servers.AdminServer)
	if this.userId > 0 {
		this.redirect(beego.URLFor("HomeController.Index"))
	}
	beego.ReadFromRequest(&this.Controller) //将request中的数据填充到底层beego.Controller中
	if this.isPost() {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))

		if username != "" && password != "" {
			user, err := adminServer.GetByName(username)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now()
				adminServer.Update(user)
				authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
				this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)

				this.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("LoginController.Login"))
		}
	}
	this.TplName = "login/login.html"
}

//后台登出
func (this *LoginController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("LoginController.Login"))
}

//无权限
func (this *LoginController) NoAuth() {
	this.Ctx.WriteString("没有权限")
}
