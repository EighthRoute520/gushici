/**********************************************
** @Des: AdminController
** @Author: EighthRoute
** @Date:   2020/11/1 19:26
***********************************************/

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gushici/libs"
	"gushici/models"
	"gushici/servers"
	"strconv"
	"strings"
	"time"
)

type AdminController struct {
	BaseController
}

//首页
func (self *AdminController) List() {
	self.Data["pageTitle"] = "管理员管理"
	self.display()
	//self.TplName = "admin/list.html"
}

//新增页
func (self *AdminController) Add() {
	self.Data["pageTitle"] = "新增管理员"

	// 角色
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, _ := (&servers.RoleServer{}).GetList(1, 1000, filters)
	list := (&servers.RoleServer{}).DealListData(result)
	self.Data["role"] = list
	self.display()
}

//编辑页
func (self *AdminController) Edit() {
	self.Data["pageTitle"] = "编辑管理员"

	id, _ := self.GetInt("id", 0)
	Admin, _ := (&servers.AdminServer{}).GetById(id)
	row := (&servers.AdminServer{}).DealOneData(Admin)
	self.Data["admin"] = row

	role_ids := strings.Split(Admin.RoleIds, ",")

	filters := make(map[string]interface{})
	filters["status"] = 1
	result, _ := (&servers.RoleServer{}).GetList(1, 1000, filters)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(role_ids); i++ {
			role_id, _ := strconv.Atoi(role_ids[i])
			if role_id == v.Id {
				row["checked"] = 1
			}
			fmt.Println(role_ids[i])
		}
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list[k] = row
	}
	self.Data["role"] = list
	self.display()
}

//数据列表
func (self *AdminController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	//查询条件
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.AdminServer{}).GetList(page, limit, filters)
	list := (&servers.AdminServer{}).DealListData(result)
	self.ajaxList("成功", MSG_OK, count, list)
}

//新增或者编辑真正保存
func (self *AdminController) AjaxSave() {
	Admin_id, _ := self.GetInt("id")
	Admin := new(models.UcAdminModel)
	if Admin_id != 0 {
		Admin, _ = (&servers.AdminServer{}).GetById(Admin_id)
	}

	Admin.LoginName = strings.TrimSpace(self.GetString("login_name"))
	Admin.RealName = strings.TrimSpace(self.GetString("real_name"))
	Admin.Phone = strings.TrimSpace(self.GetString("phone"))
	Admin.Email = strings.TrimSpace(self.GetString("email"))
	Admin.RoleIds = strings.TrimSpace(self.GetString("roleids"))
	Admin.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	Admin.UpdateId = self.userId
	Admin.Status = 1

	if Admin_id == 0 {
		// 检查登录名是否已经存在
		_, err := (&servers.AdminServer{}).GetByName(Admin.LoginName)
		if err == nil {
			self.ajaxMsg("登录名已经存在", MSG_ERR)
		}
		//新增
		pwd, salt := libs.Password(4, "")
		Admin.Password = pwd
		Admin.Salt = salt
		Admin.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		if _, err := (&servers.AdminServer{}).Add(Admin); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	Admin.Id = Admin_id
	resetPwd, _ := self.GetInt("reset_pwd")
	if resetPwd == 1 {
		pwd, salt := libs.Password(4, "")
		Admin.Password = pwd
		Admin.Salt = salt
	}
	if err := (&servers.AdminServer{}).Update(Admin); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg(strconv.Itoa(resetPwd), MSG_OK)
}

//删除
func (self *AdminController) AjaxDel() {
	Admin_id, _ := self.GetInt("id")
	Admin, _ := (&servers.AdminServer{}).GetById(Admin_id)
	Admin.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	Admin.Status = 0
	Admin.Id = Admin_id
	if Admin_id == 1 {
		self.ajaxMsg("超级管理员不允许删除", MSG_ERR)
	}

	if err := (&servers.AdminServer{}).Update(Admin); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
