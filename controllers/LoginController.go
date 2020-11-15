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
func (self *LoginController) Login() {
	adminServer := new(servers.AdminServer)
	if self.userId > 0 {
		self.redirect(beego.URLFor("HomeController.Index"))
	}
	beego.ReadFromRequest(&self.Controller) //将request中的数据填充到底层beego.Controller中
	if self.isPost() {
		username := strings.TrimSpace(self.GetString("username"))
		password := strings.TrimSpace(self.GetString("password"))

		if username != "" && password != "" {
			user, err := adminServer.GetByName(username)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = self.getClientIp()
				user.LastLogin = beego.Date(time.Now(), "Y-m-d H:i:s")
				adminServer.Update(user)
				authkey := libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
				self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)

				self.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&self.Controller)
			self.redirect(beego.URLFor("LoginController.Login"))
		}
	}
	self.TplName = "login/login.html"
}

//后台登出
func (self *LoginController) Logout() {
	self.Ctx.SetCookie("auth", "")
	self.redirect(beego.URLFor("LoginController.Login"))
}

//无权限
func (self *LoginController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}