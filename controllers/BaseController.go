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
func (self *BaseController) Prepare() {
	self.pageSize = 20
	//获取控制器名称和对应的方法
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = libs.LcFirst(controllerName[0 : len(controllerName)-10])
	self.actionName = libs.LcFirst(actionName)
	self.Data["version"] = beego.AppConfig.String("version")
	self.Data["siteName"] = beego.AppConfig.String("site.name")
	self.Data["curRoute"] = self.controllerName + self.actionName //拼装路由
	self.Data["curController"] = self.controllerName
	self.Data["curAction"] = self.actionName
	noAuth := "ads,wxApi,infoList"
	isNoAuth := strings.Contains(noAuth, self.controllerName)
	if isNoAuth == false {
		self.auth()
	}
	self.Data["loginUserId"] = self.userId
	self.Data["loginUserName"] = self.userName
}

//登录权限验证
func (self *BaseController) auth() {
	arr := strings.Split(self.Ctx.GetCookie("auth"), "|")
	self.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := (&servers.AdminServer{}).GetById(userId)
			if err == nil && password == libs.Md5([]byte(self.getClientIp()+"|"+user.Password+user.Salt)) {
				self.userId = user.Id
				self.loginName = user.LoginName
				self.userName = user.RealName
				self.user = user
				self.AdminAuth()
			}
			url := self.controllerName + "/" + self.actionName
			isHasAuth := strings.Contains(self.allowUrl, url)
			noAuth := "ajaxSave/ajaxDel/table/login/logout/getnodes/start/show/ajaxapisave"
			isNoAuth := strings.Contains(noAuth, self.actionName)
			beego.Info(1111, url, self.allowUrl, isHasAuth, noAuth, isNoAuth)
			if isHasAuth == false && isNoAuth == false {
				self.Ctx.WriteString("没有权限")
				self.ajaxMsg("没有权限", MSG_ERR)
				return
			}
		}
	}

	if self.userId == 0 && (self.controllerName != "login" && self.actionName != "login") {
		self.redirect(beego.URLFor("LoginController.Login"))
	}
}

//管理员验证
func (self *BaseController) AdminAuth() {
	// 左侧导航栏
	filters := make(map[string]interface{})
	filters["status"] = 1
	if self.userId != 1 {
		//普通管理员
		adminAuthIds, _ := (&servers.RoleAuthServer{}).GetByIds(self.user.RoleIds)
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

	self.Data["SideMenu1"] = list[:i]  //一级菜单
	self.Data["SideMenu2"] = list2[:j] //二级菜单
	self.allowUrl = allow_url + "/home/index"
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

//利用规则，找到对应的视图，然后显示
func (self *BaseController) display() {
	tplname := self.controllerName + "/" + self.actionName + ".html"
	if !self.noLayout {
		if self.Layout == "" {
			self.Layout = "public/layout.html"
		}
	}
	self.TplName = tplname
	beego.Info(22222, tplname)
}

//ajax返回信息给前端
func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//ajax返回 列表数据
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//上传图片
func (self *BaseController) UploadFile(filename string, filepath string) {
	f, h, err := self.GetFile(filename)
	out := make(map[string]interface{})
	if err != nil {
		out["msg"] = "文件读取错误"
	}
	var fileSuffix, newFile string
	fileSuffix = path.Ext(h.Filename)
	newFile = libs.GetRandomString(8) + fileSuffix
	err = self.SaveToFile("upfile", filepath+newFile)
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
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}
