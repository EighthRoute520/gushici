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

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type RoleAuthServer struct{}

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

//根据roleIds获取多个Authids
func (this *RoleAuthServer) GetByIds(roleIds string) (Authids string, err error) {
	list := make([]*models.UcRoleAuthModel, 0)
	query := orm.NewOrm().QueryTable((&models.UcRoleAuthModel{}).TableName())
	ids := strings.Split(roleIds, ",")
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

//新增一个UcRoleAuthModel
func (this *RoleAuthServer) Add(model *models.UcRoleAuthModel) (int64, error) {
	return orm.NewOrm().Insert(model)
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

//根据RoleId删除UcRoleAuthModel
func (this *RoleAuthServer) Delete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable((&models.UcRoleAuthModel{}).TableName())
	return query.Filter("role_id", id).Delete()
}
