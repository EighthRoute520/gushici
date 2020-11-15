/**********************************************
** @Des: EnvController
** @Author: EighthRoute
** @Date:   2020/10/28 21:25
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/models"
	"gushici/servers"
	"strings"
	"time"
)

type EnvController struct {
	BaseController
}

//环境设置首页
func (self *EnvController) List() {
	self.Data["pageTitle"] = "环境设置"
	self.display()
}

//新增环境设置
func (self *EnvController) Add() {
	self.Data["pageTitle"] = "新增环境"
	self.display()
}

//编辑环境设置
func (self *EnvController) Edit() {
	self.Data["pageTitle"] = "编辑环境"

	id, _ := self.GetInt("id", 0)
	setEnvModel, _ := (&servers.EnvServer{}).GetById(id)
	row := (&servers.EnvServer{}).DealOneData(setEnvModel)
	self.Data["env"] = row
	self.display()
}

//环境设置数据列表
func (self *EnvController) Table() {
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
	result, count := (&servers.EnvServer{}).GetList(page, limit, filters)
	list := (&servers.EnvServer{}).DealListData(result)

	self.ajaxList("成功", MSG_OK, count, list)
}

//新增或者编辑环境设置时的保存
func (self *EnvController) AjaxSave() {
	Env_id, _ := self.GetInt("id")
	setEnvModel := new(models.SetEnvModel)
	if Env_id != 0 {
		setEnvModel, _ = (&servers.EnvServer{}).GetById(Env_id)
	}
	setEnvModel.EnvName = strings.TrimSpace(self.GetString("env_name"))
	setEnvModel.EnvHost = strings.TrimSpace(self.GetString("env_host"))
	setEnvModel.Detail = strings.TrimSpace(self.GetString("detail"))
	setEnvModel.UpdateId = self.userId
	setEnvModel.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	setEnvModel.Status = 1

	//新增
	if Env_id == 0 {
		setEnvModel.CreateId = self.userId
		setEnvModel.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		_, err := (&servers.EnvServer{}).GetByName(setEnvModel.EnvName)

		if err == nil {
			self.ajaxMsg("环境名称已经存在", MSG_ERR)
		}

		if _, err := (&servers.EnvServer{}).Add(setEnvModel); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	// 修改
	if err := (&servers.EnvServer{}).Update(setEnvModel); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

//删除环境设置
func (self *EnvController) AjaxDel() {
	Env_id, _ := self.GetInt("id")
	setEnvModel, _ := (&servers.EnvServer{}).GetById(Env_id)
	setEnvModel.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	setEnvModel.UpdateId = self.userId
	setEnvModel.Status = 0
	setEnvModel.Id = Env_id

	if err := (&servers.EnvServer{}).Update(setEnvModel); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
