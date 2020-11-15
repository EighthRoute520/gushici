/**********************************************
** @Des: ApiSourceModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:37
***********************************************/

package models

type ApiSourceModel struct {
	Id         int
	GroupId    int
	SourceName string
	Status     int
	AuditId    int
	CreateId   int
	UpdateId   int
	CreateTime string
	UpdateTime string
	AuditTime  string
}

//获取表名称
func (self *ApiSourceModel) TableName() string {
	return TableName("api_source")
}
