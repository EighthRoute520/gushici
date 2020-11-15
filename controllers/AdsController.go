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
func (this *AdsController) Index() {
	this.Data["siteName"] = "测试使用Index"
	this.Data["pageTitle"] = "测试使用Index"
	this.display()
}

//显示详情
func (this *AdsController) Show() {
	this.Data["siteName"] = "测试使用Show"
	this.Data["pageTitle"] = "测试使用Show"
	this.TplName = "ads/show.html"
	this.display()
}

//显示图片
func (this *AdsController) ImageShow() {
	this.Data["siteName"] = "测试使用ImageShow"
	this.Data["pageTitle"] = "测试使用ImageShow"
	this.TplName = "ads/imagesshow.html"
	this.display()
}
