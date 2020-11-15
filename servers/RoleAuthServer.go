/**********************************************
** @Des: RoleAuthServer
** @Author: EighthRoute
** @Date:   2020/11/1 20:14
***********************************************/

package servers

import (
	"bytes"
	"github.com/astaxie/beego/orm"
	"gushici/models"
	"strconv"
	"strings"
)

type RoleAuthServer struct{}

//新增一个UcRoleAuthModel
func (this *RoleAuthServer) Add(model *models.UcRoleAuthModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//根据roleId获取UcRoleAuthModel
func (this *RoleAuthServer) GetById(id int) ([]*models.UcRoleAuthModel, error) {
	list := make([]*models.UcRoleAuthModel, 0)
	query := orm.NewOrm().QueryTable((&models.UcRoleAuthModel{}).TableName())
	_, err := query.Filter("role_id", id).All(&list, "AuthId")
	if err != nil {
		return nil, err
	}
	return list, nil
}

//根据RoleId删除UcRoleAuthModel
func (this *RoleAuthServer) Delete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable((&models.UcRoleAuthModel{}).TableName())
	return query.Filter("role_id", id).Delete()
}

//根据RoleIds获取多个Authids
func (this *RoleAuthServer) GetByIds(RoleIds string) (Authids string, err error) {
	list := make([]*models.UcRoleAuthModel, 0)
	query := orm.NewOrm().QueryTable((&models.UcRoleAuthModel{}).TableName())
	ids := strings.Split(RoleIds, ",")
	_, err = query.Filter("role_id__in", ids).All(&list, "AuthId")
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.Itoa(v.AuthId))
			b.WriteString(",")
		}
	}
	Authids = strings.TrimRight(b.String(), ",")
	return Authids, nil
}

//批量增加UcRoleAuthModel
func (this *RoleAuthServer) MultiAdd(list []*models.UcRoleAuthModel) (n int, err error) {
	query := orm.NewOrm().QueryTable((&models.UcRoleAuthModel{}).TableName())
	i, _ := query.PrepareInsert()
	for _, ra := range list {
		_, err := i.Insert(ra)
		if err == nil {
			n = n + 1
		}
	}
	i.Close() // 别忘记关闭 statement
	return n, err
}
