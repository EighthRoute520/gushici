/**********************************************
** @Des: RoleController
** @Author: EighthRoute
** @Date:   2020/11/1 19:25
***********************************************/

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gushici/models"
	"gushici/servers"
	"strconv"
	"strings"
	"time"
)

type RoleController struct {
	BaseController
}

//首页
func (self *RoleController) List() {
	self.Data["pageTitle"] = "角色管理"
	self.display()
}

//新增页
func (self *RoleController) Add() {
	self.Data["zTree"] = true //引入ztreecss
	self.Data["pageTitle"] = "新增角色"
	self.display()
}

//编辑页
func (self *RoleController) Edit() {
	self.Data["zTree"] = true //引入ztreecss
	self.Data["pageTitle"] = "编辑角色"

	id, _ := self.GetInt("id", 0)
	role, _ := (&servers.RoleServer{}).GetById(id)
	row := (&servers.RoleServer{}).DealOneData(role)
	self.Data["role"] = row

	//获取选择的树节点
	roleAuth, _ := (&servers.RoleAuthServer{}).GetById(id)
	authId := make([]int, 0)
	for _, v := range roleAuth {
		authId = append(authId, v.AuthId)
	}
	self.Data["auth"] = authId
	fmt.Println(authId)
	self.display()
}

//列表
func (self *RoleController) Table() {
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
	result, count := (&servers.RoleServer{}).GetList(page, limit, filters)
	list := (&servers.RoleServer{}).DealListData(result)
	self.ajaxList("成功", MSG_OK, count, list)
}

//新增或者更新真正保存
func (self *RoleController) AjaxSave() {
	role_id, _ := self.GetInt("id")
	role := new(models.UcRoleModel)
	if role_id != 0 {
		role, _ = (&servers.RoleServer{}).GetById(role_id)
	}
	role.RoleName = strings.TrimSpace(self.GetString("role_name"))
	role.Detail = strings.TrimSpace(self.GetString("detail"))
	role.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	role.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	role.Status = 1
	auths := strings.TrimSpace(self.GetString("nodes_data"))

	if role_id == 0 {
		//新增
		role.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		role.CreateId = self.userId
		if id, err := (&servers.RoleServer{}).Add(role); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		} else {
			self.batchAddRoleAuth(auths, int(id))
		}
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	role.Id = role_id
	if err := (&servers.RoleServer{}).Update(role); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		// 删除该角色权限
		(&servers.RoleAuthServer{}).Delete(role_id)
		self.batchAddRoleAuth(auths, role_id)
	}
	self.ajaxMsg("", MSG_OK)
}

//批量增加UcRoleAuthModel
func (self *RoleController) batchAddRoleAuth(auths string, roleId int) {
	ra := new(models.UcRoleAuthModel)
	authsSlice := strings.Split(auths, ",")
	for _, v := range authsSlice {
		aid, _ := strconv.Atoi(v)
		ra.AuthId = aid
		ra.RoleId = roleId
		(&servers.RoleAuthServer{}).Add(ra)
	}
}

//删除
func (self *RoleController) AjaxDel() {
	role_id, _ := self.GetInt("id")
	role, _ := (&servers.RoleServer{}).GetById(role_id)
	role.Status = 0
	role.Id = role_id
	role.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")

	if err := (&servers.RoleServer{}).Update(role); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	// 删除该角色权限
	//models.RoleAuthDelete(role_id)
	self.ajaxMsg("", MSG_OK)
}
