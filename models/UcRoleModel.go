/**********************************************
** @Des: UcRoleModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:40
***********************************************/

package models

type UcRoleModel struct {
	Id         int
	RoleName   string
	Detail     string
	CreateId   int
	UpdateId   int
	Status     int
	CreateTime string
	UpdateTime string
}

//获取表名称
func (self *UcRoleModel) TableName() string {
	return TableName("uc_role")
}
