/**********************************************
** @Des: InfoListModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:38
***********************************************/

package models

type InfoListModel struct {
	Id         int
	ClassId    int
	Title      string
	Author     string
	Keywords   string
	Tag        string
	Used       int
	Desc       string
	Content    string
	PicUrl     string `orm:"column(picurl)"`
	LinkUrl    string `orm:"column(linkurl)"`
	Media      string
	Hits       int
	OrderId    int `orm:"column(orderid)"`
	PostTime   string
	UpdateTime string
	Status     int
}

//获取表名称
func (self *InfoListModel) TableName() string {
	return TableName("info_list")
}
