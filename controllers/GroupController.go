/**********************************************
** @Des: GroupController
** @Author: EighthRoute
** @Date:   2020/11/1 19:24
***********************************************/

package controllers

import (
	"gushici/models"
	"gushici/servers"
	"strings"
	"time"
)

type GroupController struct {
	BaseController
}

//首页
func (this *GroupController) List() {
	this.Data["pageTitle"] = "分组设置"
	this.display()
}

//新增页
func (this *GroupController) Add() {
	this.Data["pageTitle"] = "新增分组"
	this.display()
}

//编辑页
func (this *GroupController) Edit() {
	this.Data["pageTitle"] = "编辑分组"

	id, _ := this.GetInt("id", 0)
	group, _ := (&servers.GroupServer{}).GetById(id)
	row := (&servers.GroupServer{}).DealOneData(group)
	this.Data["group"] = row
	this.display()
}

//列表
func (this *GroupController) Table() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}

	this.pageSize = limit
	//查询条件
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.GroupServer{}).GetList(page, this.pageSize, filters)
	list := (&servers.GroupServer{}).DealListData(result)
	this.ajaxList("成功", MSG_OK, count, list)
}

//新增或编辑保存
func (this *GroupController) AjaxSave() {
	Group_id, _ := this.GetInt("id")
	Group := new(models.SetGroupModel)
	if Group_id != 0 {
		Group, _ = (&servers.GroupServer{}).GetById(Group_id)
	}
	Group.GroupName = strings.TrimSpace(this.GetString("group_name"))
	Group.Detail = strings.TrimSpace(this.GetString("detail"))
	Group.UpdateId = this.userId
	Group.UpdateTime = time.Now()
	Group.Status = 1

	if Group_id == 0 {
		Group.CreateId = this.userId
		Group.CreateTime = time.Now()

		// 检查登录名是否已经存在
		_, err := (&servers.GroupServer{}).GetByName(Group.GroupName)
		if err == nil {
			this.ajaxMsg("分组名已经存在", MSG_ERR)
		}

		if _, err := (&servers.GroupServer{}).Add(Group); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	// 修改
	if err := (&servers.GroupServer{}).Update(Group); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

//删除
func (this *GroupController) AjaxDel() {

	Group_id, _ := this.GetInt("id")
	Group, err := (&servers.GroupServer{}).GetById(Group_id)
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	Group.UpdateTime = time.Now()
	Group.UpdateId = this.userId
	Group.Status = 0
	Group.Id = Group_id

	if err := (&servers.GroupServer{}).Update(Group); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
