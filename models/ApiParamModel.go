/**********************************************
** @Des: ApiParamModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:37
***********************************************/

package models

type ApiParamModel struct {
	Id         int
	DetailId   int
	ApiKey     string
	ApiType    string
	ApiValue   string
	ApiDetail  string
	IsNull     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime string
	UpdateTime string
}

//获取表名称
func (this *ApiParamModel) TableName() string {
	return TableName("api_param")
}
