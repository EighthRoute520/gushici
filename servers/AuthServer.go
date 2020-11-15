/**********************************************
** @Des: AuthServer
** @Author: EighthRoute
** @Date:   2020/11/1 19:48
***********************************************/

package servers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gushici/models"
)

type AuthServer struct{}

//获取列表
func (self *AuthServer) GetList(page, pageSize int, filters map[string]interface{}) ([]*models.UcAuthModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.UcAuthModel, 0)
	query := orm.NewOrm().QueryTable((&models.UcAuthModel{}).TableName())
	if len(filters) > 0 {
		for k, v := range filters {
			query = query.Filter(k, v)
		}
	}
	total, _ := query.Count()
	query.OrderBy("pid", "sort").Limit(pageSize, offset).All(&list)

	return list, total
}

//处理多条数据返回给前端
func (self *AuthServer) DealListData(models []*models.UcAuthModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := self.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (self *AuthServer) DealOneData(model *models.UcAuthModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["pId"] = model.Pid
	row["name"] = model.AuthName
	row["open"] = true
	row["auth_name"] = model.AuthName
	row["auth_url"] = model.AuthUrl
	row["sort"] = model.Sort
	row["is_show"] = model.IsShow
	row["icon"] = model.Icon
	return row
}

//根据id获取UcAuthModel
func (self *AuthServer) AuthGetListByIds(authIds string, userId int) ([]*models.UcAuthModel, error) {
	list1 := make([]*models.UcAuthModel, 0)
	var list []orm.Params
	var err error
	if userId == 1 {
		//超级管理员
		_, err = orm.NewOrm().Raw("select id,auth_name,auth_url,pid,icon,is_show from pp_uc_auth where status=? order by pid asc,sort asc", 1).Values(&list)
	} else {
		_, err = orm.NewOrm().Raw("select id,auth_name,auth_url,pid,icon,is_show from pp_uc_auth where status=1 and id in("+authIds+") order by pid asc,sort asc", authIds).Values(&list)
	}

	for k, v := range list {
		fmt.Println(k, v)
	}

	fmt.Println(list)
	return list1, err
}

//新增一条UcAuthModel
func (self *AuthServer) Add(model *models.UcAuthModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//根据id获取UcAuthModel
func (self *AuthServer) GetById(id int) (*models.UcAuthModel, error) {
	model := new(models.UcAuthModel)

	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

//更新一条UcAuthModel
func (self *AuthServer) Update(model *models.UcAuthModel) error {
	if _, err := orm.NewOrm().Update(model); err != nil {
		return err
	}
	return nil
}
