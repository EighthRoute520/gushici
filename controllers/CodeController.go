/**********************************************
** @Des: CodeController
** @Author: EighthRoute
** @Date:   2020/10/28 21:26
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/models"
	"gushici/servers"
	"strings"
	"time"
)

type CodeController struct {
	BaseController
}

//状态码设置首页
func (this *CodeController) List() {
	this.Data["pageTitle"] = "状态码设置"
	this.display()
}

//新增状态码
func (this *CodeController) Add() {
	this.Data["pageTitle"] = "新增状态码"
	this.display()
}

//编辑状态码
func (this *CodeController) Edit() {
	this.Data["pageTitle"] = "编辑状态码"

	id, _ := this.GetInt("id", 0)
	setCodeModel, _ := (&servers.CodeServer{}).GetById(id)
	row := (&servers.CodeServer{}).DealOneData(setCodeModel)
	this.Data["code"] = row
	this.display()
}

//状态码设置列表（包括分页）
func (this *CodeController) Table() {
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
	result, count := (&servers.CodeServer{}).GetList(page, limit, filters)
	list := (&servers.CodeServer{}).DealListData(result)

	this.ajaxList("成功", MSG_OK, count, list)
}

//状态码设置真正保存数据
func (this *CodeController) AjaxSave() {
	Code_id, _ := this.GetInt("id")
	setCodeModel := new(models.SetCodeModel)
	if Code_id != 0 {
		setCodeModel, _ = (&servers.CodeServer{}).GetById(Code_id)
	}

	setCodeModel.Code = strings.TrimSpace(this.GetString("code"))
	setCodeModel.Desc = strings.TrimSpace(this.GetString("desc"))
	setCodeModel.Detail = strings.TrimSpace(this.GetString("detail"))
	setCodeModel.UpdateId = this.userId
	setCodeModel.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	setCodeModel.Status = 1

	if Code_id == 0 {
		setCodeModel.CreateId = this.userId
		setCodeModel.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		_, err := (&servers.CodeServer{}).GetByName(setCodeModel.Code)
		if err == nil {
			this.ajaxMsg("状态码已经存在", MSG_ERR)
		}

		if _, err := (&servers.CodeServer{}).Add(setCodeModel); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	// 修改
	if err := (&servers.CodeServer{}).Update(setCodeModel); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

//状态码设置删除
func (this *CodeController) AjaxDel() {
	Code_id, _ := this.GetInt("id")
	setCodeModel, _ := (&servers.CodeServer{}).GetById(Code_id)
	setCodeModel.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	setCodeModel.UpdateId = this.userId
	setCodeModel.Status = 0
	setCodeModel.Id = Code_id

	if err := (&servers.CodeServer{}).Update(setCodeModel); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
