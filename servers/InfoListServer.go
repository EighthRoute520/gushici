/**********************************************
** @Des: InfoListServer
** @Author: EighthRoute
** @Date:   2020/10/24 11:42
***********************************************/

package servers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gushici/models"
	"math/rand"
)

//用一个空的结构体，是为了定义他自己的方法，这样就不会跟同命名空间下其他方法重名，尽量使用对象方法，少使用函数
type InfoListServer struct{}

//获取多条InFoListModel信息
func (this *InfoListServer) GetList(page int, pageSize int, filters map[string]interface{}) ([]*models.InfoListModel, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.InfoListModel, 0)
	query := orm.NewOrm().QueryTable((&models.InfoListModel{}).TableName())
	if len(filters) > 0 {
		for k, v := range filters {
			query = query.Filter(k, v)
		}
	}
	count, _ := query.Count()
	query.OrderBy("-orderid", "-id").Limit(pageSize, offset).All(&list) //主要list必须是引用
	return list, count
}

//处理多条数据返回给前端
func (this *InfoListServer) DealListData(models []*models.InfoListModel) []map[string]interface{} {
	count := len(models)
	list := make([]map[string]interface{}, count)
	for k, v := range models {
		row := this.DealOneData(v)
		list[k] = row
	}
	return list
}

//处理一条数据返回给前端
func (this *InfoListServer) DealOneData(model *models.InfoListModel) map[string]interface{} {
	row := make(map[string]interface{})
	row["id"] = model.Id
	row["title"] = model.Title
	row["class_id"] = model.ClassId
	row["orderid"] = model.OrderId
	row["keywords"] = model.Keywords
	row["used"] = model.Used
	row["posttime"] = model.PostTime
	row["content"] = model.Content

	if model.PicUrl == "" {
		var r = rand.Intn(10)
		model.PicUrl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
	}
	row["pic_url"] = model.PicUrl
	row["media"] = model.Media

	if model.Desc != "" {
		nameRune := []rune(model.Desc)
		lth := len(nameRune)
		if lth > 30 {
			lth = 30
		}
		row["desc"] = string(nameRune[:lth])
	}

	row["linkurl"] = model.LinkUrl
	row["author"] = model.Author
	beego.Info(444444, row)
	return row
}

//获取一条InfoListModel数据
func (this *InfoListServer) GetOneById(id int) (*models.InfoListModel, error) {
	model := new(models.InfoListModel)
	err := orm.NewOrm().QueryTable(model.TableName()).Filter("id", id).One(model) //由于model是new出来的，已经是引用类型
	if err != nil {
		return nil, err
	}
	return model, nil
}

//获取下一条InfoListModel数据
func (this *InfoListServer) GetNextOneById(id int) (*models.InfoListModel, error) {
	model := new(models.InfoListModel)
	sql := fmt.Sprintf("select id,title from "+model.TableName()+
		" where status=1 and id < %d order by id desc limit 1", id)

	err := orm.NewOrm().Raw(sql).QueryRow(model) //使用sql方式查询
	if err != nil {
		return nil, err
	}
	return model, nil
}

//增加一条InfoListModel数据
func (this *InfoListServer) Add(model *models.InfoListModel) (int64, error) {
	return orm.NewOrm().Insert(model)
}

//更新一条InfoListModel数据
func (this *InfoListServer) Update(model *models.InfoListModel) error {
	_, err := orm.NewOrm().Update(model)
	if err != nil {
		return err
	}
	return nil
}
