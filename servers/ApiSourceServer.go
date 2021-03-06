/**********************************************
** @Des: ApiServer
** @Author: EighthRoute
** @Date:   2020/10/28 21:39
***********************************************/

package servers

import (
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type ApiSourceServer struct{}

//获取列表
func (this *ApiSourceServer) GetList(page int, pageSize int, filters map[string]interface{}) ([]*models.ApiSourceModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.ApiSourceModel, 0)
	query := orm.NewOrm().QueryTable((&models.ApiSourceModel{}).TableName())
	if len(filters) > 0 {
		for k, v := range filters {
			query = query.Filter(k, v)
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

//处理多条数据返回给前端
func (this *ApiSourceServer) DealListData(models []*models.ApiSourceModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (this *ApiSourceServer) DealOneData(model *models.ApiSourceModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["source_name"] = model.SourceName
	row["group_id"] = model.GroupId
	row["status"] = model.Status
	return row
}

//根据id获取ApiSourceModel
func (this *ApiSourceServer) GetById(id int) (*models.ApiSourceModel, error) {
	model := new(models.ApiSourceModel)
	err := orm.NewOrm().QueryTable((&models.ApiSourceModel{}).TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//根据source_name查找ApiSourceModel是否存在
func (this *ApiSourceServer) GetByName(ApiName string) (*models.ApiSourceModel, error) {
	model := new(models.ApiSourceModel)
	err := orm.NewOrm().QueryTable((&models.ApiSourceModel{}).TableName()).Filter("source_name", ApiName).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//新增一条ApiSourceModel
func (this *ApiSourceServer) Add(model *models.ApiSourceModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//更新ApiSourceModel
func (this *ApiSourceServer) Update(model *models.ApiSourceModel) error {
	if _, err := orm.NewOrm().Update(model); err != nil {
		return err
	}
	return nil
}
