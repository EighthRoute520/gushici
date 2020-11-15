/**********************************************
** @Des: UcAdminModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:39
***********************************************/

package models

import "time"

type UcAdminModel struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  time.Time
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
}

//获取表名称
func (this *UcAdminModel) TableName() string {
	return TableName("uc_admin")
}
