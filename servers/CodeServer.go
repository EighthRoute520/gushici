/**********************************************
** @Des: CodeServer
** @Author: EighthRoute
** @Date:   2020/10/28 21:43
***********************************************/

package servers

import (
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type CodeServer struct{}

//获取多条SetCodeModel
func (this *CodeServer) GetList(page int, pageSize int, filters map[string]interface{}) ([]*models.SetCodeModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.SetCodeModel, 0)
	query := orm.NewOrm().QueryTable((&models.SetCodeModel{}).TableName())
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
func (this *CodeServer) DealListData(models []*models.SetCodeModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (this *CodeServer) DealOneData(model *models.SetCodeModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["code"] = model.Code
	row["detail"] = model.Detail
	row["desc"] = model.Desc
	row["create_time"] = model.CreateTime
	row["update_time"] = model.UpdateTime
	return row
}

//根据id获取SetCodeModel
func (this *CodeServer) GetById(id int) (*models.SetCodeModel, error) {
	model := new(models.SetCodeModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//根据code获取SetCodeModel
func (this *CodeServer) GetByName(code string) (*models.SetCodeModel, error) {
	model := new(models.SetCodeModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("code", code).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//新增一条SetCodeModel
func (this *CodeServer) Add(model *models.SetCodeModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//更新一条SetCodeModel
func (this *CodeServer) Update(model *models.SetCodeModel) error {
	if _, err := orm.NewOrm().Update(model); err != nil {
		return err
	}
	return nil
}
