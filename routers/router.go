package routers

import (
	"github.com/astaxie/beego"
	"gushici/controllers"
)

func init() {
	beego.Router("/", &controllers.InfoListController{}, "GET:Index") //指明只能GET方法可以访问
	beego.Router("/show/:class_id/:id", &controllers.InfoListController{}, "GET:Show")
	beego.Router("/list/:class_id", &controllers.InfoListController{}, "*:List") //GET 和 POST 都可以访问

	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/login_out", &controllers.LoginController{}, "*:Logout")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	beego.Router("/infoClass/list", &controllers.InfoClassController{}, "*:List")
	beego.Router("/infoClass/add", &controllers.InfoClassController{}, "*:Add")
	beego.Router("/infoClass/edit", &controllers.InfoClassController{}, "*:Edit")
	beego.Router("/infoClass/ajaxDel", &controllers.InfoClassController{}, "*:AjaxDel")
	beego.Router("/infoClass/ajaxSave", &controllers.InfoClassController{}, "*:AjaxSave")
	beego.Router("/infoClass/table", &controllers.InfoClassController{}, "*:Table")
	beego.Router("/infoClass/upload", &controllers.InfoClassController{}, "*:Upload")

	//下面三个路由好像没有地方调用
	beego.Router("/ads/index", &controllers.AdsController{}, "*:Index")
	beego.Router("/ads/show", &controllers.AdsController{}, "*:Show")
	beego.Router("/ads/image_show", &controllers.AdsController{}, "*:ImageShow")

	//注意：自动路由不能使用驼峰作为rootpath,应该优先使用明确意义的路由，提高可读性
	//例如 ApiSourceController 不能使用自动路由找到 /apiSource/index 路由，他能找到 /apisource/index 路由
	beego.AutoRouter(&controllers.ApiController{})
	beego.AutoRouter(&controllers.ApiMonitorController{})
	beego.AutoRouter(&controllers.EnvController{})
	beego.AutoRouter(&controllers.CodeController{})
	beego.AutoRouter(&controllers.IdiomController{})

	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})

	beego.ErrorController(&controllers.ErrorController{})
}
