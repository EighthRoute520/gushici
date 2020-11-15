/**********************************************
** @Des: ApiController
** @Author: EighthRoute
** @Date:   2020/10/28 21:22
***********************************************/

package controllers

import (
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
func (this *ApiController) List() {
	this.Data["pageTitle"] = "API接口"
	this.Data["ApiCss"] = true

	//获取分组列表
	group_id, _ := this.GetInt("gid", 0)
	//构造分组查询条件
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, _ := (&servers.SetGroupServer{}).GetList(1, 1000, filters)
	list := (&servers.SetGroupServer{}).DealListData(result)
	this.Data["Groups"] = list

	//获取Api_source列表
	filters_source := make(map[string]interface{})
	filters_source["status__in"] = []int{1, 2, 3}
	if group_id != 0 {
		filters_source["group_id"] = group_id
	}
	result_source, _ := (&servers.ApiSourceServer{}).GetList(1, 1000, filters_source)
	list_source := (&servers.ApiSourceServer{}).DealListData(result_source)
	this.Data["Source"] = list_source

	this.Data["Gid"] = group_id
	this.display()
}

//显示所有的接口详情
func (this *ApiController) Show() {
	this.Data["ApiCss"] = true

	id, _ := this.GetInt("id", 0)
	detail_result, _ := (&servers.ApiDetailServer{}).GetExtensionById(id)
	list := (&servers.ApiDetailServer{}).DealListData(detail_result)

	this.Data["Detail"] = list
	this.TplName = "api/info.html"
}

//新增资源
func (this *ApiController) Add() {
	this.Data["pageTitle"] = "新增资源"

	//查询条件
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, _ := (&servers.SetGroupServer{}).GetList(1, 1000, filters)
	list := (&servers.SetGroupServer{}).DealListData(result)

	this.Data["Groups"] = list
	this.display()
}

//编辑资源
func (this *ApiController) Edit() {
	this.Data["pageTitle"] = "编辑API"

	//获取Source
	id, _ := this.GetInt("id", 0)
	Api, err := (&servers.ApiSourceServer{}).GetById(id)
	if err != nil {
		this.Ctx.WriteString("数据不存在")
		return
	}
	row := (&servers.ApiSourceServer{}).DealOneData(Api)
	this.Data["Source"] = row

	//获取Groups
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, _ := (&servers.SetGroupServer{}).GetList(1, 1000, filters)
	list := (&servers.SetGroupServer{}).DealListData(result)
	this.Data["Groups"] = list

	this.display()
}

//新增或者编辑资源时，真正的保存
func (this *ApiController) AjaxSave() {
	Api_id, _ := this.GetInt("id")
	Api := new(models.ApiSourceModel)
	if Api_id != 0 {
		Api, _ = (&servers.ApiSourceServer{}).GetById(Api_id)
		Api.UpdateId = this.userId
	}
	Api.SourceName = strings.TrimSpace(this.GetString("source_name"))
	Api.GroupId, _ = this.GetInt("group_id")
	Api.UpdateTime = time.Now()
	Api.Status = 2

	if Api_id == 0 {
		Api.CreateId = this.userId
		Api.CreateTime = time.Now()

		// 检查登录名是否已经存在
		_, err := (&servers.ApiSourceServer{}).GetByName(Api.SourceName)
		if err == nil {
			this.ajaxMsg("资源名已经存在", MSG_ERR)
		}

		if _, err := (&servers.ApiSourceServer{}).Add(Api); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	if err := (&servers.ApiSourceServer{}).Update(Api); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

//删除资源
func (this *ApiController) AjaxDel() {
	Api_id, _ := this.GetInt("id")
	Api, _ := (&servers.ApiSourceServer{}).GetById(Api_id)
	Api.UpdateTime = time.Now()
	Api.UpdateId = this.userId
	Api.Status = 0
	Api.Id = Api_id

	if err := (&servers.ApiSourceServer{}).Update(Api); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

//新增接口(API)
func (this *ApiController) AddApi() {
	this.Data["pageTitle"] = ""

	source_id, _ := this.GetInt("sid")
	this.Data["Sid"] = source_id

	//查询条件
	this.display()
}

//编辑接口(API)
func (this *ApiController) EditApi() {
	id, _ := this.GetInt("id", 0)
	detail, _ := (&servers.ApiDetailServer{}).GetById(id)
	params, _ := (&servers.ApiParamServer{}).GetById(detail.Id)
	this.Data["Detail"] = detail
	this.Data["Params"] = params
	this.Data["ParamsCount"] = len(params)
	this.display()
}

//新增或者编辑接口(API)时，真正的保存
func (this *ApiController) AjaxApiSave() {
	Api_id, _ := this.GetInt("id")
	ApiDetail := new(models.ApiDetailModel)
	if Api_id != 0 {
		ApiDetail, _ = (&servers.ApiDetailServer{}).GetById(Api_id)
	}

	ApiDetail.SourceId, _ = this.GetInt("source_id")
	ApiDetail.ProtocolType, _ = this.GetInt("protocol_type")
	ApiDetail.Method, _ = this.GetInt("method")
	ApiDetail.ApiName = strings.TrimSpace(this.GetString("api_name"))
	ApiDetail.ApiUrl = strings.TrimSpace(this.GetString("api_url"))
	ApiDetail.Result = strings.TrimSpace(this.GetString("result"))
	ApiDetail.Example = strings.TrimSpace(this.GetString("example"))
	ApiDetail.Detail = strings.TrimSpace(this.GetString("detail"))
	ApiDetail.UpdateId = this.userId
	ApiDetail.UpdateTime = time.Now()
	ApiDetail.Status = 1

	//新增
	if Api_id == 0 {
		ApiDetail.CreateId = this.userId
		ApiDetail.CreateTime = time.Now()
		detail_id, err := (&servers.ApiDetailServer{}).Add(ApiDetail)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}

		//批量增加ApiParamModel
		err = this.batchAdd4Form(int(detail_id))
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}

		this.ajaxMsg("", MSG_OK)
	}

	//修改
	ApiDetail.Id = Api_id
	ApiDetail.Status, _ = this.GetInt("status")
	if err := (&servers.ApiDetailServer{}).Update(ApiDetail); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}

	//先删除后新增参数
	if _, err := (&servers.ApiParamServer{}).Delete(int64(Api_id), this.userId); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}

	//批量增加ApiParamModel
	err := this.batchAdd4Form(Api_id)
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}

	this.ajaxMsg("", MSG_OK)
}

//批量增加ApiParamModel
func (this *ApiController) batchAdd4Form(Api_id int) error {
	params := make(map[int]map[string]string)
	for k, v := range this.Ctx.Request.Form {
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
		apiParam.CreateId = this.userId
		apiParam.UpdateId = this.userId
		apiParam.CreateTime = time.Now()
		apiParam.UpdateTime = time.Now()
		apiParam.Status = 1

		if _, err := (&servers.ApiParamServer{}).Add(apiParam); err != nil {
			return err
		}
	}
	return nil
}
