/**********************************************
** @Des: SetEnvModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:38
***********************************************/

package models

import "time"

type SetEnvModel struct {
	Id         int
	EnvName    string
	EnvHost    string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
}

//获取表名称
func (this *SetEnvModel) TableName() string {
	return TableName("set_env")
}
