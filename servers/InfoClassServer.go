/**********************************************
** @Des: InfoClassServer
** @Author: EighthRoute
** @Date:   2020/10/24 11:42
***********************************************/

package servers

import (
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type InfoClassServer struct{}

//获取多条InfoClassModel
func (this *InfoClassServer) GetList(page int, pageSize int, filters map[string]interface{}) ([]*models.InfoClassModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.InfoClassModel, 0)
	query := orm.NewOrm().QueryTable((&models.InfoClassModel{}).TableName())
	if len(filters) > 0 {
		for k, v := range filters {
			query = query.Filter(k, v)
		}
	}
	count, _ := query.Count()
	query.OrderBy("-orderid", "-id").Limit(pageSize, offset).All(&list)
	return list, count
}

//处理多条InfoClassModel返回调用者,返回形式为 [key=>value]形式map
func (this *InfoClassServer) DealListDataMap(models []*models.InfoClassModel) map[int]string {
	classMap := make(map[int]string)
	for _, v := range models {
		classMap[v.Id] = v.ClassName
	}
	return classMap
}

//处理多条InfoClassModel返回调用者
func (this *InfoClassServer) DealListData(models []*models.InfoClassModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (this *InfoClassServer) DealOneData(model *models.InfoClassModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["class_name"] = model.ClassName
	row["linkurl"] = model.LinkUrl
	row["desc"] = model.Desc
	row["orderid"] = model.OrderId

	return row
}

//根据id获取InfoClassModel
func (this *InfoClassServer) GetById(id int) (*models.InfoClassModel, error) {
	model := new(models.InfoClassModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//新增一条InfoClassModel
func (this *InfoClassServer) Add(model *models.InfoClassModel) (int64, error) {
	id, err := orm.NewOrm().Insert(model)
	if err != nil {
		return 0, err
	}
	return id, err
}

//更新InfoClassModel
func (this *InfoClassServer) Update(model *models.InfoClassModel) error {
	_, err := orm.NewOrm().Update(model)
	if err != nil {
		return err
	}
	return nil
}
