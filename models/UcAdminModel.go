/**********************************************
** @Des: UcAdminModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:39
***********************************************/

package models

type UcAdminModel struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  string
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime string
	UpdateTime string
}

//获取表名称
func (this *UcAdminModel) TableName() string {
	return TableName("uc_admin")
}
