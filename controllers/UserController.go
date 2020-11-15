/**********************************************
** @Des: UserController
** @Author: EighthRoute
** @Date:   2020/11/1 19:26
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/libs"
	"gushici/servers"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

//编辑
func (self *UserController) Edit() {
	self.Data["pageTitle"] = "资料修改"
	id := self.userId
	admin, _ := (&servers.AdminServer{}).GetById(id)
	row := (&servers.AdminServer{}).DealOneData(admin)
	self.Data["admin"] = row
	self.display()
}

//保存
func (self *UserController) AjaxSave() {
	Admin_id, _ := self.GetInt("id")
	Admin, _ := (&servers.AdminServer{}).GetById(Admin_id)
	//修改
	Admin.Id = Admin_id
	Admin.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	Admin.UpdateId = self.userId
	Admin.LoginName = strings.TrimSpace(self.GetString("login_name"))
	Admin.RealName = strings.TrimSpace(self.GetString("real_name"))
	Admin.Phone = strings.TrimSpace(self.GetString("phone"))
	Admin.Email = strings.TrimSpace(self.GetString("email"))
	resetPwd := self.GetString("reset_pwd")
	if resetPwd == "1" {
		pwdOld := strings.TrimSpace(self.GetString("password_old"))
		pwdOldMd5 := libs.Md5([]byte(pwdOld + Admin.Salt))
		if Admin.Password != pwdOldMd5 {
			self.ajaxMsg("旧密码错误", MSG_ERR)
		}

		pwdNew1 := strings.TrimSpace(self.GetString("password_new1"))
		pwdNew2 := strings.TrimSpace(self.GetString("password_new2"))

		if pwdNew1 != pwdNew2 {
			self.ajaxMsg("两次密码不一致", MSG_ERR)
		}

		pwd, salt := libs.Password(4, pwdNew1)
		Admin.Password = pwd
		Admin.Salt = salt
	}

	Admin.Status = 1
	if err := (&servers.AdminServer{}).Update(Admin); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
