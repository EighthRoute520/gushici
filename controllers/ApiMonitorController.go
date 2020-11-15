/**********************************************
** @Des: ApiMonitorController
** @Author: EighthRoute
** @Date:   2020/11/15 14:00
***********************************************/

package controllers

type ApiMonitorController struct {
	BaseController
}

func (this *ApiMonitorController) List() {
	this.Data["pageTitle"] = "API文档"
	this.display()
}
