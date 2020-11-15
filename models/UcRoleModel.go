/**********************************************
** @Des: UcRoleModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:40
***********************************************/

package models

import "time"

type UcRoleModel struct {
	Id         int
	RoleName   string
	Detail     string
	CreateId   int
	UpdateId   int
	Status     int
	CreateTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
}

//获取表名称
func (this *UcRoleModel) TableName() string {
	return TableName("uc_role")
}
