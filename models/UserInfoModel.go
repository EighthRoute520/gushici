/**********************************************
** @Des: UserInfoModel
** @Author: EighthRoute
** @Date:   2020/10/24 10:40
***********************************************/

package models

type UserInfoModel struct {
	Id            int
	OpenId        string
	UserName      string
	UserAvatar    int
	UserPassword  string
	FromType      string
	CreateTime    string
	LastLoginTime string
	UserNickname  string
}

//获取表名称
func (self *UserInfoModel) TableName() string {
	return TableName("user_info")
}
