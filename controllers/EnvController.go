/**********************************************
** @Des: EnvController
** @Author: EighthRoute
** @Date:   2020/10/28 21:25
***********************************************/

package controllers

import (
	"gushici/models"
	"gushici/servers"
	"strings"
	"time"
)

type EnvController struct {
	BaseController
}

//环境设置首页
func (this *EnvController) List() {
	this.Data["pageTitle"] = "环境设置"
	this.display()
}

//新增环境设置
func (this *EnvController) Add() {
	this.Data["pageTitle"] = "新增环境"
	this.display()
}

//编辑环境设置
func (this *EnvController) Edit() {
	this.Data["pageTitle"] = "编辑环境"

	id, _ := this.GetInt("id", 0)
	setEnvModel, _ := (&servers.EnvServer{}).GetById(id)
	row := (&servers.EnvServer{}).DealOneData(setEnvModel)
	this.Data["env"] = row
	this.display()
}

//环境设置数据列表
func (this *EnvController) Table() {
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
	result, count := (&servers.EnvServer{}).GetList(page, limit, filters)
	list := (&servers.EnvServer{}).DealListData(result)

	this.ajaxList("成功", MSG_OK, count, list)
}

//新增或者编辑环境设置时的保存
func (this *EnvController) AjaxSave() {
	Env_id, _ := this.GetInt("id")
	setEnvModel := new(models.SetEnvModel)
	if Env_id != 0 {
		setEnvModel, _ = (&servers.EnvServer{}).GetById(Env_id)
	}
	setEnvModel.EnvName = strings.TrimSpace(this.GetString("env_name"))
	setEnvModel.EnvHost = strings.TrimSpace(this.GetString("env_host"))
	setEnvModel.Detail = strings.TrimSpace(this.GetString("detail"))
	setEnvModel.UpdateId = this.userId
	setEnvModel.UpdateTime = time.Now()
	setEnvModel.Status = 1

	//新增
	if Env_id == 0 {
		setEnvModel.CreateId = this.userId
		setEnvModel.CreateTime = time.Now()
		_, err := (&servers.EnvServer{}).GetByName(setEnvModel.EnvName)

		if err == nil {
			this.ajaxMsg("环境名称已经存在", MSG_ERR)
		}

		if _, err := (&servers.EnvServer{}).Add(setEnvModel); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	// 修改
	if err := (&servers.EnvServer{}).Update(setEnvModel); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

//删除环境设置
func (this *EnvController) AjaxDel() {
	Env_id, _ := this.GetInt("id")
	setEnvModel, _ := (&servers.EnvServer{}).GetById(Env_id)
	setEnvModel.UpdateTime = time.Now()
	setEnvModel.UpdateId = this.userId
	setEnvModel.Status = 0
	setEnvModel.Id = Env_id

	if err := (&servers.EnvServer{}).Update(setEnvModel); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
