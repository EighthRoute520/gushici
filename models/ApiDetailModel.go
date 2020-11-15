/**********************************************
** @Des: ApiDetailModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:36
***********************************************/

package models

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
	AuditTime    string
	AuditId      int
	Status       int
	CreateId     int
	UpdateId     int
	CreateTime   string
	UpdateTime   string
}

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
