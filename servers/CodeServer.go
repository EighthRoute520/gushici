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

type CodeServer struct{}

//获取多条SetCodeModel
func (self *CodeServer) GetList(page int, pageSize int, filters map[string]interface{}) ([]*models.SetCodeModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.SetCodeModel, 0)
	query := orm.NewOrm().QueryTable((&models.SetCodeModel{}).TableName())
	if len(filters) > 0 {
		for k, v := range filters {
			query.Filter(k, v)
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

//处理多条数据返回给前端
func (self *CodeServer) DealListData(models []*models.SetCodeModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := self.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (self *CodeServer) DealOneData(model *models.SetCodeModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["code"] = model.Code
	row["detail"] = model.Detail
	row["desc"] = model.Desc
	row["create_time"] = model.CreateTime
	row["update_time"] = model.UpdateTime
	return row
}

//新增一条SetCodeModel
func (self *CodeServer) Add(model *models.SetCodeModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//根据id获取SetCodeModel
func (self *CodeServer) GetById(id int) (*models.SetCodeModel, error) {
	model := new(models.SetCodeModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//根据env_name获取SetCodeModel
func (self *CodeServer) GetByName(EnvName string) (*models.SetCodeModel, error) {
	model := new(models.SetCodeModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("env_name", EnvName).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//更新一条SetCodeModel
func (self *CodeServer) Update(model *models.SetCodeModel) error {
	if _, err := orm.NewOrm().Update(model); err != nil {
		return err
	}
	return nil
}
