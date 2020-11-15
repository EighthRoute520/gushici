/**********************************************
** @Des: UcAdminServer
** @Author: EighthRoute
** @Date:   2020/10/25 11:20
***********************************************/

package servers

import (
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type AdminServer struct{}

//获取多条UcAdminModel
func (this *AdminServer) GetList(page, pageSize int, filters map[string]interface{}) ([]*models.UcAdminModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.UcAdminModel, 0)
	query := orm.NewOrm().QueryTable((&models.UcAdminModel{}).TableName())
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
func (this *AdminServer) DealListData(models []*models.UcAdminModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (this *AdminServer) DealOneData(model *models.UcAdminModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["login_name"] = model.LoginName
	row["real_name"] = model.RealName
	row["phone"] = model.Phone
	row["email"] = model.Email
	row["role_ids"] = model.RoleIds
	row["create_time"] = model.CreateTime
	row["update_time"] = model.UpdateTime
	return row
}

//根据id获取AdminModel
func (this *AdminServer) GetById(id int) (*models.UcAdminModel, error) {
	model := new(models.UcAdminModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//根据login_name获取UcAdminModel
func (this *AdminServer) GetByName(name string) (*models.UcAdminModel, error) {
	model := new(models.UcAdminModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("login_name", name).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//新增UcAdminModel
func (this *AdminServer) Add(model *models.UcAdminModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//更新UcAdminModel
func (this *AdminServer) Update(model *models.UcAdminModel) error {
	_, err := orm.NewOrm().Update(model)
	if err != nil {
		return err
	}

	return nil
}
