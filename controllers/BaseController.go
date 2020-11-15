/**********************************************
** @Des: base
** @Author: EighthRoute
** @Date:   2020/10/24 10:34
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/libs"
	"gushici/models"
	"gushici/servers"
	"path"
	"strconv"
	"strings"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

//基础控制器，其他控制器都从该控制器继承
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.UcAdminModel
	userId         int
	userName       string
	loginName      string
	pageSize       int
	allowUrl       string
	noLayout       bool
}

//前期准备：重写父类方法，实现数据初始化
func (this *BaseController) Prepare() {
	this.pageSize = 20
	//获取控制器名称和对应的方法
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = libs.LcFirst(controllerName[0 : len(controllerName)-10])
	this.actionName = libs.LcFirst(actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["curRoute"] = this.controllerName + this.actionName //拼装路由
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	noAuth := "ads,wxApi,infoList"
	isNoAuth := strings.Contains(noAuth, this.controllerName)
	if isNoAuth == false {
		this.auth()
	}
	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
}

//登录权限验证
func (this *BaseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	this.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := (&servers.AdminServer{}).GetById(userId)
			if err == nil && password == libs.Md5([]byte(this.getClientIp()+"|"+user.Password+user.Salt)) {
				this.userId = user.Id
				this.loginName = user.LoginName
				this.userName = user.RealName
				this.user = user
				this.AdminAuth()
			}
			url := this.controllerName + "/" + this.actionName
			//isHasAuth := strings.Contains(this.allowUrl, strings.ToLower(url))
			isHasAuth := strings.Contains(this.allowUrl, url)
			noAuth := "ajaxSave/ajaxDel/table/login/logout/getnodes/start/show/ajaxApiSave"
			isNoAuth := strings.Contains(noAuth, this.actionName)
			beego.Info(1111, url, this.allowUrl, isHasAuth, noAuth, isNoAuth)
			if isHasAuth == false && isNoAuth == false {
				//this.Ctx.WriteString("没有权限")
				this.ajaxMsg("没有权限", MSG_ERR)
				return
			}
		}
	}

	if this.userId == 0 && (this.controllerName != "login" && this.actionName != "login") {
		this.redirect(beego.URLFor("LoginController.Login"))
	}
}

//管理员验证
func (this *BaseController) AdminAuth() {
	// 左侧导航栏
	filters := make(map[string]interface{})
	filters["status"] = 1
	if this.userId != 1 {
		//普通管理员
		adminAuthIds, _ := (&servers.RoleAuthServer{}).GetByIds(this.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters["id__in"] = adminAuthIdArr
	}
	result, _ := (&servers.AuthServer{}).GetList(1, 1000, filters)
	list := make([]map[string]interface{}, len(result))
	list2 := make([]map[string]interface{}, len(result))
	allow_url := ""
	i, j := 0, 0
	for _, v := range result {
		if v.AuthUrl != " " || v.AuthUrl != "/" {
			allow_url += v.AuthUrl
		}
		row := make(map[string]interface{})
		if v.Pid == 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list[i] = row
			i++
		}
		if v.Pid != 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list2[j] = row
			j++
		}
	}

	this.Data["SideMenu1"] = list[:i]  //一级菜单
	this.Data["SideMenu2"] = list2[:j] //二级菜单
	this.allowUrl = allow_url + "/home/index"
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

//利用规则，找到对应的视图，然后显示
func (this *BaseController) display() {
	tplname := this.controllerName + "/" + this.actionName + ".html"
	if !this.noLayout {
		if this.Layout == "" {
			this.Layout = "public/layout.html"
		}
	}
	this.TplName = tplname
	beego.Info(22222, tplname)
}

//ajax返回信息给前端
func (this *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//ajax返回 列表数据
func (this *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//上传图片
func (this *BaseController) UploadFile(filename string, filepath string) {
	f, h, err := this.GetFile(filename)
	out := make(map[string]interface{})
	if err != nil {
		out["msg"] = "文件读取错误"
	}
	var fileSuffix, newFile string
	fileSuffix = path.Ext(h.Filename)
	newFile = libs.GetRandomString(8) + fileSuffix
	err = this.SaveToFile("upfile", filepath+newFile)
	if err != nil {
		out["msg"] = "文件保存错误"
	}
	defer f.Close()
	out["state"] = "SUCCESS"
	out["url"] = filepath + newFile
	out["title"] = newFile
	out["original"] = h.Filename
	out["size"] = h.Size
	out["msg"] = "ok"
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}
