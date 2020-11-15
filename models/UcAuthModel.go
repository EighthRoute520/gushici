/**********************************************
** @Des: UcAuthModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:40
***********************************************/

package models

type UcAuthModel struct {
	Id         int
	Pid        int
	AuthName   string
	AuthUrl    string
	Sort       int
	Icon       string
	IsShow     int
	UserId     int
	CreateId   int
	UpdateId   int
	Status     int
	CreateTime string
	UpdateTime string
}

//获取表名称
func (self *UcAuthModel) TableName() string {
	return TableName("uc_auth")
}
