/**********************************************
** @Des: UcRoleAuthModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:40
***********************************************/

package models

type UcRoleAuthModel struct {
	AuthId int `orm:"pk"`
	RoleId int
}

//获取表名称
func (self *UcRoleAuthModel) TableName() string {
	return TableName("uc_role_auth")
}
