/**********************************************
** @Des: wwwcontroller
** @Author: EighthRoute
** @Date:   2020/10/23 22:28
***********************************************/

package controllers

import (
	"gushici/servers"
)

type InfoListController struct {
	BaseController
}

//首页
func (this *InfoListController) Index() {
	infoListServer := new(servers.InfoListServer)
	//开心儿歌
	filters := make(map[string]interface{})
	filters["status"] = 1
	filters["class_id"] = 5
	result, _ := infoListServer.GetList(1, 6, filters)
	list := infoListServer.DealListData(result)

	//儿童古诗
	filters2 := make(map[string]interface{})
	filters2["status"] = 1
	filters2["class_id"] = 3
	result2, _ := infoListServer.GetList(1, 6, filters2)
	list2 := infoListServer.DealListData(result2)

	//国学生活
	filters3 := make(map[string]interface{})
	filters3["status"] = 1
	filters3["class_id"] = 1
	result3, _ := infoListServer.GetList(1, 16, filters3)
	list3 := infoListServer.DealListData(result3)

	//诗词古韵
	filters4 := make(map[string]interface{})
	filters4["status"] = 1
	filters4["class_id"] = 2
	result4, _ := infoListServer.GetList(1, 6, filters4)
	list4 := infoListServer.DealListData(result4)

	//组装返回给前端
	out := make(map[string]interface{})
	out["list"] = list
	out["list2"] = list2
	out["list3"] = list3
	out["list4"] = list4
	out["class_id"] = 0
	this.Data["data"] = out
	this.Layout = "public/www_layout.html"
	this.display()
}

//显示详情
func (this *InfoListController) Show() {
	infoListServer := new(servers.InfoListServer)
	id, _ := this.GetInt(":id")
	//当前详情
	infoListModel, _ := infoListServer.GetOneById(id)
	row := infoListServer.DealOneData(infoListModel)
	row["class_id"] = 0
	if infoListModel != nil {
		row["desc"] = infoListModel.Desc
		row["content"] = infoListModel.Content
		row["posttime"] = infoListModel.PostTime
	}

	//下一条详情
	infoListModelNext, _ := infoListServer.GetNextOneById(id)
	nextRow := make(map[string]interface{})
	if infoListModelNext != nil {
		nextRow["id"] = infoListModelNext.Id
		nextRow["title"] = infoListModelNext.Title
	}

	row["next"] = nextRow
	this.Data["data"] = row
	this.Layout = "public/www_layout.html"
	this.display()
}

//列表
func (this *InfoListController) List() {
	infoListServer := new(servers.InfoListServer)
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 16
	}

	//查询条件
	classId, err := this.GetInt(":class_id")
	filters := make(map[string]interface{})
	filters["status"] = 1
	if err == nil {
		filters["class_id"] = classId
	}
	result, count := infoListServer.GetList(page, limit, filters)
	list := infoListServer.DealListData(result)

	infoClassServer := new(servers.InfoClassServer)
	infoClassResult, _ := infoClassServer.GetList(1, 4, nil)
	classMap := infoClassServer.DealListData(infoClassResult)
	out := make(map[string]interface{})
	out["count"] = count
	out["class_id"] = classId
	out["page"] = page
	out["class_name"] = classMap[classId]
	out["title"] = classMap[classId]
	out["list"] = list
	this.Data["data"] = out

	this.Layout = "public/www_layout.html"
	this.display()
}
