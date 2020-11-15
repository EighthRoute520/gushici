/**********************************************
** @Des: ApiController
** @Author: EighthRoute
** @Date:   2020/10/28 21:22
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/models"
	"gushici/servers"
	"strconv"
	"strings"
	"time"
)

//该控制器控制资源和资源下面具体的API
type ApiController struct {
	BaseController
}

//显示资源首页
func (self *ApiController) List() {
	self.Data["pageTitle"] = "API接口"
	self.Data["ApiCss"] = true

	//获取分组列表
	group_id, _ := self.GetInt("gid", 0)
	//构造分组查询条件
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, _ := (&servers.SetGroupServer{}).GetList(1, 1000, filters)
	list := (&servers.SetGroupServer{}).DealListData(result)
	self.Data["Groups"] = list

	//获取Api_source列表
	filters_source := make(map[string]interface{})
	filters_source["status__in"] = []int{1, 2, 3}
	if group_id != 0 {
		filters_source["group_id"] = group_id
	}
	result_source, _ := (&servers.ApiSourceServer{}).GetList(1, 1000, filters_source)
	list_source := (&servers.ApiSourceServer{}).DealListData(result_source)
	self.Data["Source"] = list_source

	self.Data["Gid"] = group_id
	self.display()
}

//显示所有的接口详情
func (self *ApiController) Show() {
	self.Data["ApiCss"] = true

	id, _ := self.GetInt("id", 0)
	detail_result, _ := (&servers.ApiDetailServer{}).GetExtensionById(id)
	list := (&servers.ApiDetailServer{}).DealListData(detail_result)

	self.Data["Detail"] = list
	self.TplName = "api/info.html"
}

//新增资源
func (self *ApiController) Add() {
	self.Data["pageTitle"] = "新增资源"

	//查询条件
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, _ := (&servers.SetGroupServer{}).GetList(1, 1000, filters)
	list := (&servers.SetGroupServer{}).DealListData(result)

	self.Data["Groups"] = list
	self.display()
}

//编辑资源
func (self *ApiController) Edit() {
	self.Data["pageTitle"] = "编辑API"

	//获取Source
	id, _ := self.GetInt("id", 0)
	Api, err := (&servers.ApiSourceServer{}).GetById(id)
	if err != nil {
		self.Ctx.WriteString("数据不存在")
		return
	}
	row := (&servers.ApiSourceServer{}).DealOneData(Api)
	self.Data["Source"] = row

	//获取Groups
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, _ := (&servers.SetGroupServer{}).GetList(1, 1000, filters)
	list := (&servers.SetGroupServer{}).DealListData(result)
	self.Data["Groups"] = list

	self.display()
}

//新增或者编辑资源时，真正的保存
func (self *ApiController) AjaxSave() {
	Api_id, _ := self.GetInt("id")
	Api := new(models.ApiSourceModel)
	if Api_id != 0 {
		Api, _ = (&servers.ApiSourceServer{}).GetById(Api_id)
		Api.UpdateId = self.userId
		Api.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	}
	Api.SourceName = strings.TrimSpace(self.GetString("source_name"))
	Api.GroupId, _ = self.GetInt("group_id")
	Api.Status = 2

	// 检查登录名是否已经存在
	_, err := (&servers.ApiSourceServer{}).GetByName(Api.SourceName)
	if err == nil {
		self.ajaxMsg("资源名已经存在", MSG_ERR)
	}

	if Api_id == 0 {
		Api.CreateId = self.userId
		Api.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		if _, err := (&servers.ApiSourceServer{}).Add(Api); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	if err := (&servers.ApiSourceServer{}).Update(Api); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

//删除资源
func (self *ApiController) AjaxDel() {
	Api_id, _ := self.GetInt("id")
	Api, _ := (&servers.ApiSourceServer{}).GetById(Api_id)
	Api.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	Api.UpdateId = self.userId
	Api.Status = 0
	Api.Id = Api_id

	if err := (&servers.ApiSourceServer{}).Update(Api); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

//新增接口(API)
func (self *ApiController) AddApi() {
	self.Data["pageTitle"] = ""

	source_id, _ := self.GetInt("sid")
	self.Data["Sid"] = source_id

	//查询条件
	self.display()
}

//编辑接口(API)
func (self *ApiController) EditApi() {
	id, _ := self.GetInt("id", 0)
	detail, _ := (&servers.ApiDetailServer{}).GetById(id)
	params, _ := (&servers.ApiParamServer{}).GetById(detail.Id)
	self.Data["Detail"] = detail
	self.Data["Params"] = params
	self.Data["ParamsCount"] = len(params)
	self.display()
}

//新增或者编辑接口(API)时，真正的保存
func (self *ApiController) AjaxApiSave() {
	Api_id, _ := self.GetInt("id")
	ApiDetail := new(models.ApiDetailModel)
	if Api_id != 0 {
		ApiDetail, _ = (&servers.ApiDetailServer{}).GetById(Api_id)
	}

	ApiDetail.SourceId, _ = self.GetInt("source_id")
	ApiDetail.ProtocolType, _ = self.GetInt("protocol_type")
	ApiDetail.Method, _ = self.GetInt("method")
	ApiDetail.ApiName = strings.TrimSpace(self.GetString("api_name"))
	ApiDetail.ApiUrl = strings.TrimSpace(self.GetString("api_url"))
	ApiDetail.Result = strings.TrimSpace(self.GetString("result"))
	ApiDetail.Example = strings.TrimSpace(self.GetString("example"))
	ApiDetail.Detail = strings.TrimSpace(self.GetString("detail"))
	ApiDetail.UpdateId = self.userId
	ApiDetail.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
	ApiDetail.Status = 1

	//新增
	if Api_id == 0 {
		ApiDetail.CreateId = self.userId
		ApiDetail.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		detail_id, err := (&servers.ApiDetailServer{}).Add(ApiDetail)
		if err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}

		//批量增加ApiParamModel
		err = self.batchAdd4Form(int(detail_id))
		if err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}

		self.ajaxMsg("", MSG_OK)
	}

	//修改
	ApiDetail.Id = Api_id
	ApiDetail.Status, _ = self.GetInt("status")
	if err := (&servers.ApiDetailServer{}).Update(ApiDetail); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	//先删除后新增参数
	if _, err := (&servers.ApiParamServer{}).Delete(int64(Api_id), self.userId); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	//批量增加ApiParamModel
	err := self.batchAdd4Form(Api_id)
	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	self.ajaxMsg("", MSG_OK)
}

//批量增加ApiParamModel
func (self *ApiController) batchAdd4Form(Api_id int) error {
	params := make(map[int]map[string]string)
	for k, v := range self.Ctx.Request.Form {
		if strings.Contains(k, "attr_") == true {
			ks := strings.Split(k, "_")
			i, _ := strconv.Atoi(ks[1])
			if _, ok := params[i]; ok {
				params[i][ks[2]] = v[0]
			} else {
				param := make(map[string]string)
				param[ks[2]] = v[0]
				params[i] = param
			}
		}
	}
	for _, vv := range params {
		apiParam := new(models.ApiParamModel)
		if vv["key"] == "" {
			break
		}
		apiParam.ApiKey = vv["key"]
		apiParam.ApiType = vv["type"]
		apiParam.ApiValue = vv["value"]
		apiParam.ApiDetail = vv["detail"]
		apiParam.IsNull = vv["isnull"]
		apiParam.DetailId = Api_id
		apiParam.CreateId = self.userId
		apiParam.UpdateId = self.userId
		apiParam.CreateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		apiParam.UpdateTime = beego.DateFormat(time.Now(), "Y-m-d H:i:s")
		apiParam.Status = 1

		if _, err := (&servers.ApiParamServer{}).Add(apiParam); err != nil {
			return err
		}
	}
	return nil
}
