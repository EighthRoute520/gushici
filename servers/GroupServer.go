/**********************************************
** @Des: GroupServer
** @Author: EighthRoute
** @Date:   2020/11/1 19:30
***********************************************/

package servers

import (
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

type GroupServer struct{}

//获取多条Group
func (self *GroupServer) GetList(page, pageSize int, filters map[string]interface{}) ([]*models.SetGroupModel, int64) {
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
func (self *GroupServer) DealListData(models []*models.SetGroupModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := self.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (self *GroupServer) DealOneData(model *models.SetGroupModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["group_name"] = model.GroupName
	row["detail"] = model.Detail
	row["create_time"] = model.CreateTime
	row["update_time"] = model.UpdateTime
	return row
}

//新增SetGroupModel
func (self *GroupServer) Add(model *models.SetGroupModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//根据id获取SetGroupModel
func (self *GroupServer) GetById(id int) (*models.SetGroupModel, error) {
	model := new(models.SetGroupModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//根据groupName获取SetGroupModel
func (self *GroupServer) GetByName(groupName string) (*models.SetGroupModel, error) {
	model := new(models.SetGroupModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("group_name", groupName).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//更新SetGroupModel
func (self *GroupServer) Update(model *models.SetGroupModel) error {
	if _, err := orm.NewOrm().Update(model); err != nil {
		return err
	}
	return nil
}
