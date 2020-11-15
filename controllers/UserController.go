/**********************************************
** @Des: UserController
** @Author: EighthRoute
** @Date:   2020/11/1 19:26
***********************************************/

package controllers

import (
	"gushici/libs"
	"gushici/servers"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

//编辑
func (this *UserController) Edit() {
	this.Data["pageTitle"] = "资料修改"
	id := this.userId
	admin, _ := (&servers.AdminServer{}).GetById(id)
	row := (&servers.AdminServer{}).DealOneData(admin)
	this.Data["admin"] = row
	this.display()
}

//保存
func (this *UserController) AjaxSave() {
	Admin_id, _ := this.GetInt("id")
	Admin, _ := (&servers.AdminServer{}).GetById(Admin_id)
	//修改
	Admin.Id = Admin_id
	Admin.UpdateTime = time.Now()
	Admin.UpdateId = this.userId
	Admin.LoginName = strings.TrimSpace(this.GetString("login_name"))
	Admin.RealName = strings.TrimSpace(this.GetString("real_name"))
	Admin.Phone = strings.TrimSpace(this.GetString("phone"))
	Admin.Email = strings.TrimSpace(this.GetString("email"))
	resetPwd := this.GetString("reset_pwd")
	if resetPwd == "1" {
		pwdOld := strings.TrimSpace(this.GetString("password_old"))
		pwdOldMd5 := libs.Md5([]byte(pwdOld + Admin.Salt))
		if Admin.Password != pwdOldMd5 {
			this.ajaxMsg("旧密码错误", MSG_ERR)
		}

		pwdNew1 := strings.TrimSpace(this.GetString("password_new1"))
		pwdNew2 := strings.TrimSpace(this.GetString("password_new2"))

		if pwdNew1 != pwdNew2 {
			this.ajaxMsg("两次密码不一致", MSG_ERR)
		}

		pwd, salt := libs.Password(4, pwdNew1)
		Admin.Password = pwd
		Admin.Salt = salt
	}

	Admin.Status = 1
	if err := (&servers.AdminServer{}).Update(Admin); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
