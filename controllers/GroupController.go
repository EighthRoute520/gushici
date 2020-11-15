/**********************************************
** @Des: GroupController
** @Author: EighthRoute
** @Date:   2020/11/1 19:24
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/models"
	"gushici/servers"
	"strings"
	"time"
)

type GroupController struct {
	BaseController
}

//首页
func (self *GroupController) List() {
	self.Data["pageTitle"] = "分组设置"
	self.display()
}

//新增页
func (self *GroupController) Add() {
	self.Data["pageTitle"] = "新增分组"
	self.display()
}

//编辑页
func (self *GroupController) Edit() {
	self.Data["pageTitle"] = "编辑分组"

	id, _ := self.GetInt("id", 0)
	group, _ := (&servers.GroupServer{}).GetById(id)
	row := (&servers.GroupServer{}).DealOneData(group)
	self.Data["group"] = row
	self.display()
}

//列表
func (self *GroupController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	self.pageSize = limit
	//查询条件
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.GroupServer{}).GetList(page, self.pageSize, filters)
	list := (&servers.GroupServer{}).DealListData(result)
	self.ajaxList("成功", MSG_OK, count, list)
}

//新增或编辑保存
func (self *GroupController) AjaxSave() {
	Group_id, _ := self.GetInt("id")
	Group := new(models.SetGroupModel)
	if Group_id == 0 {
		Group, _ = (&servers.GroupServer{}).GetById(Group_id)
	}
	Group.GroupName = strings.TrimSpace(self.GetString("group_name"))
	Group.Detail = strings.TrimSpace(self.GetString("detail"))
	Group.UpdateId = self.userId
	Group.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	Group.Status = 1

	if Group_id == 0 {
		Group.CreateId = self.userId
		Group.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")

		// 检查登录名是否已经存在
		_, err := (&servers.GroupServer{}).GetByName(Group.GroupName)
		if err == nil {
			self.ajaxMsg("分组名已经存在", MSG_ERR)
		}

		if _, err := (&servers.GroupServer{}).Add(Group); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	// 修改
	if err := (&servers.GroupServer{}).Update(Group); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

//删除
func (self *GroupController) AjaxDel() {

	Group_id, _ := self.GetInt("id")
	Group, _ := (&servers.GroupServer{}).GetById(Group_id)
	Group.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	Group.UpdateId = self.userId
	Group.Status = 0
	Group.Id = Group_id

	if err := (&servers.GroupServer{}).Update(Group); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
