/**********************************************
** @Des: ApiParamServer
** @Author: EighthRoute
** @Date:   2020/10/31 19:54
***********************************************/

package servers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gushici/models"
	"time"
)

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type ApiParamServer struct{}

//根据ApiDetailModel的Id获取ApiParamModel列表
func (this *ApiParamServer) GetById(id int) ([]*models.ApiParamModel, error) {
	list := make([]*models.ApiParamModel, 0)
	query := orm.NewOrm().QueryTable((&models.ApiParamModel{}).TableName())
	query.Filter("detail_id", id).Filter("status", 1).All(&list)
	return list, nil
}

//新增一条ApiParamModel
func (this *ApiParamServer) Add(model *models.ApiParamModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//删除一条ApiParamModel
func (this *ApiParamServer) Delete(id int64, update_id int) (int64, error) {
	sql := "UPDATE pp_api_param SET status=0,update_id=?,update_time=? WHERE detail_id=?"
	res, err := orm.NewOrm().Raw(sql, update_id, beego.Date(time.Now(), "Y-m-d H:i:s"), id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num, nil
	}
	return 0, err
}
