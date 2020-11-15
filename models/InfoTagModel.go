/**********************************************
** @Des: InfoTagModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:38
***********************************************/

package models

type InfoTagModel struct {
	Id        int
	TagName   string
	TagPinyin string
}

//获取表名称
func (self *InfoTagModel) TableName() string {
	return TableName("info_tag")
}
