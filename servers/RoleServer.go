/**********************************************
** @Des: RoleServer
** @Author: EighthRoute
** @Date:   2020/11/1 20:04
***********************************************/

package servers

import (
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type RoleServer struct{}

//获取多条UcRoleModel
func (this *RoleServer) GetList(page, pageSize int, filters map[string]interface{}) ([]*models.UcRoleModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.UcRoleModel, 0)
	query := orm.NewOrm().QueryTable((&models.UcRoleModel{}).TableName())
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
func (this *RoleServer) DealListData(models []*models.UcRoleModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (this *RoleServer) DealOneData(model *models.UcRoleModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["role_name"] = model.RoleName
	row["detail"] = model.Detail
	row["create_time"] = model.CreateTime
	row["update_time"] = model.UpdateTime
	return row
}

//根据ID获取一条UcRoleModel数据
func (this *RoleServer) GetById(id int) (*models.UcRoleModel, error) {
	model := new(models.UcRoleModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//增加一条UcRoleModel数据
func (this *RoleServer) Add(model *models.UcRoleModel) (int64, error) {
	id, err := orm.NewOrm().Insert(model)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//更新一条UcRoleModel数据
func (this *RoleServer) Update(model *models.UcRoleModel) error {
	if _, err := orm.NewOrm().Update(model); err != nil {
		return err
	}
	return nil
}
