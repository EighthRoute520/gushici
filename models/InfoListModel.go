/**********************************************
** @Des: InfoListModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:38
***********************************************/

package models

import "time"

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
	OrderId    int       `orm:"column(orderid)"`
	PostTime   time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
	Status     int
}

//获取表名称
func (this *InfoListModel) TableName() string {
	return TableName("info_list")
}
