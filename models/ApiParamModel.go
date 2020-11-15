/**********************************************
** @Des: ApiParamModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:37
***********************************************/

package models

import "time"

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
	CreateTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
}

//获取表名称
func (this *ApiParamModel) TableName() string {
	return TableName("api_param")
}
