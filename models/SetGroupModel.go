/**********************************************
** @Des: SetGroupModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:39
***********************************************/

package models

import "time"

type SetGroupModel struct {
	Id         int
	GroupName  string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
}

//获取表名称
func (this *SetGroupModel) TableName() string {
	return TableName("set_group")
}
