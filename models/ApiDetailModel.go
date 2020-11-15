/**********************************************
** @Des: ApiDetailModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:36
***********************************************/

package models

import "time"

type ApiDetailModel struct {
	Id           int
	SourceId     int
	Method       int
	ApiName      string
	ApiUrl       string
	ProtocolType int
	Result       string
	Example      string
	Detail       string
	AuditTime    time.Time
	AuditId      int
	Status       int
	CreateId     int
	UpdateId     int
	CreateTime   time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime   time.Time `orm:"auto_now;type(timestamp)"`
}

//ApiDetailModel扩展类
type ApiDetailExtensionModel struct {
	ApiDetailModel ApiDetailModel
	CreateName     string
	UpdateName     string
	AuditName      string
}

//获取表名称
func (this *ApiDetailModel) TableName() string {
	return TableName("api_detail")
}
