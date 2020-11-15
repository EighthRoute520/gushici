package main

import (
	"github.com/astaxie/beego"
	"gushici/models"
	_ "gushici/routers"
)

func main() {
	//数据库初始化
	models.Init()
	//静态资源设置
	beego.SetStaticPath("/uploads", "static/upload")
	beego.Run()
}
