/**********************************************
** @Des: DatabaseInit
** @Author: EighthRoute
** @Date:   2020/10/23 22:18
***********************************************/

package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

//必须首字母大写，才能让别人访问
func Init() {
	//获取配置文件中的配置信息
	dbhost := beego.AppConfig.String("db.host")
	dbusername := beego.AppConfig.String("db.username")
	dbpassword := beego.AppConfig.String("db.password")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.dbname")
	dbcharset := beego.AppConfig.String("db.charset")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbusername + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + dbcharset

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	// 设置数据库连接信息
	orm.RegisterDataBase("default", "mysql", dsn, 30)

	// 注册模型
	orm.RegisterModel(new(ApiDetailModel), new(ApiParamModel), new(ApiSourceModel), new(InfoClassModel), new(InfoIdiomModel),
		new(InfoListModel), new(InfoTagModel), new(SetCodeModel), new(SetEnvModel), new(SetGroupModel), new(UcAdminModel),
		new(UcAuthModel), new(UcRoleAuthModel), new(UcRoleModel), new(UserInfoModel))

	//根据runmode配置参数决定是否数据库开启Debug模式，Debug模式将会打印SQL语句
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

//所有表在调用该方法之后，会自动加上前缀信息
func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
