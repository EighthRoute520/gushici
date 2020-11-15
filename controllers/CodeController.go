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
func (self *CodeController) List() {
	self.Data["pageTitle"] = "状态码设置"
	self.display()
}

//新增状态码
func (self *CodeController) Add() {
	self.Data["pageTitle"] = "新增状态码"
	self.display()
}

//编辑状态码
func (self *CodeController) Edit() {
	self.Data["pageTitle"] = "编辑状态码"

	id, _ := self.GetInt("id", 0)
	setCodeModel, _ := (&servers.CodeServer{}).GetById(id)
	row := (&servers.CodeServer{}).DealOneData(setCodeModel)
	self.Data["code"] = row
	self.display()
}

//状态码设置列表（包括分页）
func (self *CodeController) Table() {
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
	result, count := (&servers.CodeServer{}).GetList(page, limit, filters)
	list := (&servers.CodeServer{}).DealListData(result)

	self.ajaxList("成功", MSG_OK, count, list)
}

//状态码设置真正保存数据
func (self *CodeController) AjaxSave() {
	Code_id, _ := self.GetInt("id")
	setCodeModel := new(models.SetCodeModel)
	if Code_id != 0 {
		setCodeModel, _ = (&servers.CodeServer{}).GetById(Code_id)
	}

	setCodeModel.Code = strings.TrimSpace(self.GetString("code"))
	setCodeModel.Desc = strings.TrimSpace(self.GetString("desc"))
	setCodeModel.Detail = strings.TrimSpace(self.GetString("detail"))
	setCodeModel.UpdateId = self.userId
	setCodeModel.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	setCodeModel.Status = 1

	if Code_id == 0 {
		setCodeModel.CreateId = self.userId
		setCodeModel.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		_, err := (&servers.CodeServer{}).GetByName(setCodeModel.Code)
		if err == nil {
			self.ajaxMsg("状态码已经存在", MSG_ERR)
		}

		if _, err := (&servers.CodeServer{}).Add(setCodeModel); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	// 修改
	if err := (&servers.CodeServer{}).Update(setCodeModel); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

//状态码设置删除
func (self *CodeController) AjaxDel() {
	Code_id, _ := self.GetInt("id")
	setCodeModel, _ := (&servers.CodeServer{}).GetById(Code_id)
	setCodeModel.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	setCodeModel.UpdateId = self.userId
	setCodeModel.Status = 0
	setCodeModel.Id = Code_id

	if err := (&servers.CodeServer{}).Update(setCodeModel); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
