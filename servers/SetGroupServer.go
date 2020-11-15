/**********************************************
** @Des: SetGroupServer
** @Author: EighthRoute
** @Date:   2020/10/31 18:50
***********************************************/

package servers

import (
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

type SetGroupServer struct{}

//获取分组列表
func (this *SetGroupServer) GetList(page int, pageSize int, filters map[string]interface{}) ([]*models.SetGroupModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.SetGroupModel, 0)
	query := orm.NewOrm().QueryTable((&models.SetGroupModel{}).TableName())
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
func (this *SetGroupServer) DealListData(models []*models.SetGroupModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (this *SetGroupServer) DealOneData(model *models.SetGroupModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["group_name"] = model.GroupName
	return row
}
