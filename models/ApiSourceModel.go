/**********************************************
** @Des: ApiSourceModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:37
***********************************************/

package models

import "time"

type ApiSourceModel struct {
	Id         int
	GroupId    int
	SourceName string
	Status     int
	AuditId    int
	CreateId   int
	UpdateId   int
	CreateTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
	AuditTime  time.Time
}

//获取表名称
func (this *ApiSourceModel) TableName() string {
	return TableName("api_source")
}
