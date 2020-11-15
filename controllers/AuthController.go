/**********************************************
** @Des: AuthController
** @Author: EighthRoute
** @Date:   2020/11/1 19:25
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/models"
	"gushici/servers"
	"strings"
	"time"
)

type AuthController struct {
	BaseController
}

//首页
func (self *AuthController) Index() {

	self.Data["pageTitle"] = "权限因子"
	self.display()
}

//权限树列表
func (self *AuthController) List() {
	self.Data["zTree"] = true //引入ztreecss
	self.Data["pageTitle"] = "权限因子"
	self.display()
}

//获取全部节点
func (self *AuthController) GetNodes() {
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.AuthServer{}).GetList(1, 1000, filters)
	list := (&servers.AuthServer{}).DealListData(result)

	self.ajaxList("成功", MSG_OK, count, list)
}

//获取一个节点
func (self *AuthController) GetNode() {
	id, _ := self.GetInt("id")
	result, _ := (&servers.AuthServer{}).GetById(id)
	row := (&servers.AuthServer{}).DealOneData(result)

	self.ajaxList("成功", MSG_OK, 0, row)
}

//新增或修改
func (self *AuthController) AjaxSave() {
	id, _ := self.GetInt("id")
	auth := new(models.UcAuthModel)
	if id != 0 {
		auth, _ = (&servers.AuthServer{}).GetById(id)
	}

	auth.UserId = self.userId
	auth.Pid, _ = self.GetInt("pid")
	auth.AuthName = strings.TrimSpace(self.GetString("auth_name"))
	auth.AuthUrl = strings.TrimSpace(self.GetString("auth_url"))
	auth.Sort, _ = self.GetInt("sort")
	auth.IsShow, _ = self.GetInt("is_show")
	auth.Icon = strings.TrimSpace(self.GetString("icon"))
	auth.Status = 1

	if id == 0 {
		//新增
		auth.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		auth.CreateId = self.userId
		if _, err := (&servers.AuthServer{}).Add(auth); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}

	//更新
	auth.Id = id
	auth.UpdateId = self.userId
	auth.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	if err := (&servers.AuthServer{}).Update(auth); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	self.ajaxMsg("", MSG_OK)
}

//删除
func (self *AuthController) AjaxDel() {
	id, _ := self.GetInt("id")
	auth, _ := (&servers.AuthServer{}).GetById(id)
	auth.Id = id
	auth.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	auth.Status = 0
	if err := (&servers.AuthServer{}).Update(auth); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
