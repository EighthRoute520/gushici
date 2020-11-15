/**********************************************
** @Des: SetGroupModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:39
***********************************************/

package models

type SetGroupModel struct {
	Id         int
	GroupName  string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime string
	UpdateTime string
}

//获取表名称
func (self *SetGroupModel) TableName() string {
	return TableName("set_group")
}
