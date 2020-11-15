/**********************************************
** @Des: SetCodeModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:38
***********************************************/

package models

type SetCodeModel struct {
	Id         int
	Code       string
	Desc       string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime string
	UpdateTime string
}

//获取表名称
func (this *SetCodeModel) TableName() string {
	return TableName("set_code")
}
