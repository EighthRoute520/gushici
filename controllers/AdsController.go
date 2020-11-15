/**********************************************
** @Des: AdsController
** @Author: EighthRoute
** @Date:   2020/10/28 21:07
***********************************************/

package controllers

type AdsController struct {
	BaseController
}

//首页
func (self *AdsController) Index() {
	self.Data["siteName"] = "测试使用Index"
	self.Data["pageTitle"] = "测试使用Index"
	self.display()
}

//显示详情
func (self *AdsController) Show() {
	self.Data["siteName"] = "测试使用Show"
	self.Data["pageTitle"] = "测试使用Show"
	self.TplName = "ads/show.html"
	self.display()
}

//显示图片
func (self *AdsController) ImageShow() {
	self.Data["siteName"] = "测试使用ImageShow"
	self.Data["pageTitle"] = "测试使用ImageShow"
	self.TplName = "ads/imagesshow.html"
	self.display()
}
