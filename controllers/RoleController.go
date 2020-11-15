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
func (this *RoleController) List() {
	this.Data["pageTitle"] = "角色管理"
	this.display()
}

//新增页
func (this *RoleController) Add() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "新增角色"
	this.display()
}

//编辑页
func (this *RoleController) Edit() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "编辑角色"

	id, _ := this.GetInt("id", 0)
	role, _ := (&servers.RoleServer{}).GetById(id)
	row := (&servers.RoleServer{}).DealOneData(role)
	this.Data["role"] = row

	//获取选择的树节点
	roleAuth, _ := (&servers.RoleAuthServer{}).GetById(id)
	authId := make([]int, 0)
	for _, v := range roleAuth {
		authId = append(authId, v.AuthId)
	}
	this.Data["auth"] = authId
	fmt.Println(authId)
	this.display()
}

//列表
func (this *RoleController) Table() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}
	//查询条件
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.RoleServer{}).GetList(page, limit, filters)
	list := (&servers.RoleServer{}).DealListData(result)
	this.ajaxList("成功", MSG_OK, count, list)
}

//新增或者更新真正保存
func (this *RoleController) AjaxSave() {
	role_id, _ := this.GetInt("id")
	role := new(models.UcRoleModel)
	if role_id != 0 {
		role, _ = (&servers.RoleServer{}).GetById(role_id)
	}
	role.RoleName = strings.TrimSpace(this.GetString("role_name"))
	role.Detail = strings.TrimSpace(this.GetString("detail"))
	role.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	role.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	role.Status = 1
	auths := strings.TrimSpace(this.GetString("nodes_data"))

	if role_id == 0 {
		//新增
		role.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		role.CreateId = this.userId
		if id, err := (&servers.RoleServer{}).Add(role); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		} else {
			this.batchAddRoleAuth(auths, int(id))
		}
		this.ajaxMsg("", MSG_OK)
	}

	//修改
	role.Id = role_id
	if err := (&servers.RoleServer{}).Update(role); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		// 删除该角色权限
		(&servers.RoleAuthServer{}).Delete(role_id)
		this.batchAddRoleAuth(auths, role_id)
	}
	this.ajaxMsg("", MSG_OK)
}

//批量增加UcRoleAuthModel
func (this *RoleController) batchAddRoleAuth(auths string, roleId int) {
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
func (this *RoleController) AjaxDel() {
	role_id, _ := this.GetInt("id")
	role, _ := (&servers.RoleServer{}).GetById(role_id)
	role.Status = 0
	role.Id = role_id
	role.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")

	if err := (&servers.RoleServer{}).Update(role); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	// 删除该角色权限
	//models.RoleAuthDelete(role_id)
	this.ajaxMsg("", MSG_OK)
}
