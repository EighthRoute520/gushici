/**********************************************
** @Des: AuthController
** @Author: EighthRoute
** @Date:   2020/11/1 19:25
***********************************************/

package controllers

import (
	"gushici/models"
	"gushici/servers"
	"strings"
	"time"
)

type AuthController struct {
	BaseController
}

//首页
func (this *AuthController) Index() {

	this.Data["pageTitle"] = "权限因子"
	this.display()
}

//权限树列表
func (this *AuthController) List() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "权限因子"
	this.display()
}

//获取全部节点
func (this *AuthController) GetNodes() {
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.AuthServer{}).GetList(1, 1000, filters)
	list := (&servers.AuthServer{}).DealListData(result)

	this.ajaxList("成功", MSG_OK, count, list)
}

//获取一个节点
func (this *AuthController) GetNode() {
	id, _ := this.GetInt("id")
	result, _ := (&servers.AuthServer{}).GetById(id)
	row := (&servers.AuthServer{}).DealOneData(result)

	this.ajaxList("成功", MSG_OK, 0, row)
}

//新增或修改
func (this *AuthController) AjaxSave() {
	id, _ := this.GetInt("id")
	auth := new(models.UcAuthModel)
	if id != 0 {
		auth, _ = (&servers.AuthServer{}).GetById(id)
	}

	auth.UserId = this.userId
	auth.Pid, _ = this.GetInt("pid")
	auth.AuthName = strings.TrimSpace(this.GetString("auth_name"))
	auth.AuthUrl = strings.TrimSpace(this.GetString("auth_url"))
	auth.Sort, _ = this.GetInt("sort")
	auth.IsShow, _ = this.GetInt("is_show")
	auth.Icon = strings.TrimSpace(this.GetString("icon"))
	auth.Status = 1

	if id == 0 {
		//新增
		auth.CreateTime = time.Now()
		auth.CreateId = this.userId
		if _, err := (&servers.AuthServer{}).Add(auth); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
	}

	//更新
	auth.Id = id
	auth.UpdateId = this.userId
	auth.UpdateTime = time.Now()
	if err := (&servers.AuthServer{}).Update(auth); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}

	this.ajaxMsg("", MSG_OK)
}

//删除
func (this *AuthController) AjaxDel() {
	id, _ := this.GetInt("id")
	auth, _ := (&servers.AuthServer{}).GetById(id)
	auth.Id = id
	auth.UpdateTime = time.Now()
	auth.Status = 0
	if err := (&servers.AuthServer{}).Update(auth); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
