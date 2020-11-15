/**********************************************
** @Des: InfoIdiomModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:37
***********************************************/

package models

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
	CreateTime  string
}

//获取表名称
func (self *InfoIdiomModel) TableName() string {
	return TableName("info_idiom")
}
