/**********************************************
** @Des: UcAuthModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:40
***********************************************/

package models

import "time"

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
	CreateTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
}

//获取表名称
func (this *UcAuthModel) TableName() string {
	return TableName("uc_auth")
}
