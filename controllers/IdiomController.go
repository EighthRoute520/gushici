/**********************************************
** @Des: IdiomController
** @Author: EighthRoute
** @Date:   2020/11/15 17:02
***********************************************/

package controllers

type IdiomController struct {
	BaseController
}

//后台首页
func (this *IdiomController) List() {
	this.Data["pageTitle"] = "成语大全首页"
	this.display()
}
