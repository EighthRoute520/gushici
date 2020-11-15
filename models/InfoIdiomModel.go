/**********************************************
** @Des: InfoIdiomModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:37
***********************************************/

package models

import "time"

type InfoIdiomModel struct {
	Id          int
	Image       string
	Title       string
	Keywords    string
	Description string
	Question    string
	Content     string
	Sort        int
	Status      int
	time.Time   `orm:"auto_now_add;type(timestamp)"`
}

//获取表名称
func (this *InfoIdiomModel) TableName() string {
	return TableName("info_idiom")
}
