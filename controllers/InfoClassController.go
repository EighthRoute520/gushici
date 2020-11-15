/**********************************************
** @Des: InfoClassController
** @Author: EighthRoute
** @Date:   2020/10/25 13:52
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"gushici/models"
	"gushici/servers"
	"os"
	"strings"
	"time"
)

type InfoClassController struct {
	BaseController
}

//资讯管理首页
func (this *InfoClassController) List() {
	class_id, _ := this.GetInt("class_id")
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.InfoClassServer{}).GetList(1, 10, filters)
	classList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["class_name"] = v.ClassName
		row["linkurl"] = v.LinkUrl
		row["desc"] = v.Desc
		row["orderid"] = v.OrderId
		row["count"] = count
		classList[k] = row
	}
	this.Data["pageTitle"] = "资讯管理"
	this.Data["news_class"] = classList
	this.Data["class_id"] = class_id
	this.display()
}

//新增资讯
func (this *InfoClassController) Add() {
	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.InfoClassServer{}).GetList(1, 20, filters)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["class_name"] = v.ClassName
		row["desc"] = v.Desc
		row["linkurl"] = v.LinkUrl
		row["orderid"] = v.OrderId
		row["count"] = count
		list[k] = row
	}
	this.Data["pageTitle"] = "新增资讯"
	this.Data["news_class"] = list
	this.display()
}

//修改资讯
func (this *InfoClassController) Edit() {
	id, _ := this.GetInt("id")
	infoListModel, _ := (&servers.InfoListServer{}).GetOneById(id)
	infoRow := make(map[string]interface{})
	infoRow["id"] = infoListModel.Id
	infoRow["title"] = infoListModel.Title
	infoRow["class_id"] = infoListModel.ClassId
	infoRow["orderid"] = infoListModel.OrderId
	infoRow["keywords"] = infoListModel.Keywords
	infoRow["used"] = infoListModel.Used
	infoRow["desc"] = infoListModel.Desc
	infoRow["content"] = infoListModel.Content
	infoRow["pic_url"] = infoListModel.PicUrl
	infoRow["media"] = infoListModel.Media
	infoRow["author"] = infoListModel.Author
	infoRow["posttime"] = infoListModel.PostTime

	filters := make(map[string]interface{})
	filters["status"] = 1
	result, count := (&servers.InfoClassServer{}).GetList(1, 10, filters)
	classList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["class_name"] = v.ClassName
		row["linkurl"] = v.LinkUrl
		row["desc"] = v.Desc
		row["orderid"] = v.OrderId
		row["count"] = count
		classList[k] = row
	}

	this.Data["pageTitle"] = "编辑资讯"
	this.Data["news_class"] = classList
	this.Data["news"] = infoRow
	this.display()
}

//Ajax删除
func (this *InfoClassController) AjaxDel() {
	id, _ := this.GetInt("id")
	infoListModel, _ := (&servers.InfoListServer{}).GetOneById(id)
	infoListModel.Status = 2
	infoListModel.Id = id
	infoListModel.UpdateTime = beego.Date(time.Now(), "Y-m-d H:i:s")
	err := (&servers.InfoListServer{}).Update(infoListModel)
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("删除成功", MSG_OK)
}

//Ajax新增保存
func (this *InfoClassController) AjaxSave() {
	id, _ := this.GetInt("id")
	classId, _ := this.GetInt("class_id")
	orderid, _ := this.GetInt("orderid")

	infoListModel := new(models.InfoListModel)
	if id != 0 {
		infoListModel, _ = (&servers.InfoListServer{}).GetOneById(id)
	}
	infoListModel.Title = strings.TrimSpace(this.GetString("title"))
	infoListModel.Author = strings.TrimSpace(this.GetString("author"))
	infoListModel.Keywords = strings.TrimSpace(this.GetString("keywords"))
	infoListModel.Used, _ = this.GetInt("used")
	infoListModel.Desc = strings.TrimSpace(this.GetString("desc"))
	infoListModel.Content = strings.TrimSpace(this.GetString("content"))
	infoListModel.ClassId = classId
	infoListModel.OrderId = orderid
	infoListModel.UpdateTime = beego.Date(time.Now(), "Y-m-d H:i:s")
	infoListModel.PicUrl = strings.TrimSpace(this.GetString("pic_url"))
	infoListModel.Media = strings.TrimSpace(this.GetString("media"))
	infoListModel.Status = 1
	if id == 0 {
		infoListModel.PostTime = beego.Date(time.Now(), "Y-m-d H:i:s")
		if _, err := (&servers.InfoListServer{}).Add(infoListModel); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	if err := (&servers.InfoListServer{}).Update(infoListModel); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("操作成功", MSG_OK)
}

//表格显示列表信息
func (this *InfoClassController) Table() {
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}
	filters := make(map[string]interface{})
	filters["status"] = 1
	class_id, _ := this.GetInt("class_id")
	if class_id > 0 {
		filters["class_id"] = class_id
	}
	result, count := (&servers.InfoListServer{}).GetList(page, limit, filters)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["class_id"] = v.ClassId
		row["orderid"] = v.OrderId
		row["keywords"] = v.Keywords
		row["used"] = v.Used
		row["desc"] = v.Desc
		row["pic_url"] = v.PicUrl
		row["author"] = v.Author
		row["posttime"] = v.PostTime
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

//上传图片
func (this *InfoClassController) Upload() {
	filepath := "static/upload/" + beego.Date(time.Now(), "Y-m-d") + "/"
	_, err := os.Stat(filepath)
	if err != nil {
		err = os.Mkdir(filepath, 0777)
		if err != nil {
			this.ajaxMsg("上传失败", MSG_ERR)
		}
	}
	this.UploadFile("upfile", filepath)
}
