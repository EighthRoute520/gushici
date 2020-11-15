/**********************************************
** @Des: ApiDetailServer
** @Author: EighthRoute
** @Date:   2020/10/31 19:35
***********************************************/

package servers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

var (
	AUDIT_STATUS   = [4]string{"暂停使用", "正在开发", "正在审核", "审核通过"}
	PROTOCOL_TYPE  = [3]string{"HTTP/HTTPS", "HTTP", "HTTPS"}
	REQUEST_METHOD = [6]string{"未知", "GET", "POST", "PUT", "PATCH", "DELETE"}
)

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type ApiDetailServer struct{}

//根据ID获取ApiDetailModel
func (this *ApiDetailServer) GetById(id int) (*models.ApiDetailModel, error) {
	model := new(models.ApiDetailModel)
	err := orm.NewOrm().QueryTable((&models.ApiDetailModel{}).TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//根据ID获取ApiDetailExtensionModel
func (this *ApiDetailServer) GetExtensionById(id int) ([]*models.ApiDetailExtensionModel, error) {
	list := make([]*models.ApiDetailExtensionModel, 0)
	sql := "SELECT pp_api_detail.*,a.real_name as create_name,b.real_name as update_name,c.real_name as audit_name FROM pp_api_detail LEFT JOIN pp_uc_admin as a ON pp_api_detail.create_id=a.id LEFT JOIN pp_uc_admin as b ON pp_api_detail.update_id=b.id LEFT JOIN pp_uc_admin as c ON pp_api_detail.audit_id=c.id WHERE pp_api_detail.source_id=?"
	orm.NewOrm().Raw(sql, id).QueryRows(&list)
	fmt.Println(list)
	return list, nil
}

//处理多条数据返回给前端
func (this *ApiDetailServer) DealListData(models []*models.ApiDetailExtensionModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据
func (this *ApiDetailServer) DealOneData(model *models.ApiDetailExtensionModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.ApiDetailModel.Id
	row["source_id"] = model.ApiDetailModel.SourceId
	row["api_url"] = model.ApiDetailModel.ApiUrl
	row["api_name"] = model.ApiDetailModel.ApiName
	row["example"] = model.ApiDetailModel.Example
	row["detail"] = model.ApiDetailModel.Detail
	row["result"] = model.ApiDetailModel.Result
	row["status"] = model.ApiDetailModel.Status
	row["create_name"] = model.CreateName
	row["update_name"] = model.UpdateName
	row["audit_name"] = model.AuditName
	row["protocol_type"] = PROTOCOL_TYPE[model.ApiDetailModel.ProtocolType]
	row["audit_status"] = AUDIT_STATUS[model.ApiDetailModel.Status]
	row["method"] = REQUEST_METHOD[model.ApiDetailModel.Method]
	row["audit_time"] = model.ApiDetailModel.AuditTime
	row["update_time"] = model.ApiDetailModel.UpdateTime
	//参数
	row["Params"], _ = (&ApiParamServer{}).GetById(model.ApiDetailModel.Id)
	return row
}

//新增一条ApiDetailModel
func (this *ApiDetailServer) Add(model *models.ApiDetailModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//更新一条ApiDetailModel
func (this *ApiDetailServer) Update(model *models.ApiDetailModel) error {
	_, err := orm.NewOrm().Update(model)
	if err != nil {
		return err
	}
	return nil
}
