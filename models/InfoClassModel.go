/**********************************************
** @Des: InfoClassModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:37
***********************************************/

package models

type InfoClassModel struct {
	Id        int
	ClassName string
	LinkUrl   string `orm:"column(linkurl)"`
	Desc      string
	OrderId   int `orm:"column(orderid)"`
	Status    int
}

//获取表名称
func (self *InfoClassModel) TableName() string {
	return TableName("info_class")
}
