/**********************************************
** @Des: EnvServer
** @Author: EighthRoute
** @Date:   2020/10/28 21:45
***********************************************/

package servers

import (
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

type EnvServer struct{}

//获取多条SetEnvModel
func (this *EnvServer) GetList(page int, pageSize int, filters map[string]interface{}) ([]*models.SetEnvModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.SetEnvModel, 0)
	query := orm.NewOrm().QueryTable((&models.SetEnvModel{}).TableName())
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
func (this *EnvServer) DealListData(models []*models.SetEnvModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (this *EnvServer) DealOneData(model *models.SetEnvModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["env_name"] = model.EnvName
	row["env_host"] = model.EnvHost
	row["detail"] = model.Detail
	row["create_time"] = model.CreateTime
	row["update_time"] = model.UpdateTime
	return row
}

//新增一条SetEnvModel
func (this *EnvServer) Add(model *models.SetEnvModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//根据id获取SetEnvModel
func (this *EnvServer) GetById(id int) (*models.SetEnvModel, error) {
	model := new(models.SetEnvModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//根据env_name获取SetEnvModel
func (this *EnvServer) GetByName(EnvName string) (*models.SetEnvModel, error) {
	model := new(models.SetEnvModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("env_name", EnvName).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//更新一条SetEnvModel
func (this *EnvServer) Update(model *models.SetEnvModel) error {
	if _, err := orm.NewOrm().Update(model); err != nil {
		return err
	}
	return nil
}
