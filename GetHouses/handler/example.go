package handler

import (
	"context"

		example "renting/GetHouses/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"strconv"
	"renting/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetHouses(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("搜索房源 GetHouses /api/v1.0/houses")

	//初始化返回值
	rsp.Error  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Error)

	//1.获取数据
	var aid int //地区
	aid, _ = strconv.Atoi(req.Aid)
	var sd string //起始时间
	sd  = req.Sd
	var ed string //结束时间
	ed = req.Ed
	var sk string //第三栏的信息
	sk = req.Sk
	var page int // 页
	page ,_ = strconv.Atoi(req.P)
	beego.Info(aid, sd, ed, sk, page)


	//2.根据aid查找传入地区的所有房屋
	var houses []models.House
	o := orm.NewOrm()
	num,err := o.QueryTable("house").Filter("area_id", aid).All(&houses)
	if err != nil {
		rsp.Error  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	//3.//计算以下所有房屋/一页显示的数量
	totalPage := int(num)/models.HOUSE_LIST_PAGE_CAPACITY + 1


	//4.设置当前页数为1
	housePage := 1

	//5.创建房屋列表容器（[]interface{}{}）
	houseList := []interface{}{}

	//6.关联查询该地区每一个房子对应的地区，发布者，图片设施等信息，处理后存入房屋列表容器
	for _, house := range houses {
		o.LoadRelated(&house,"Area")
		o.LoadRelated(&house,"User")
		o.LoadRelated(&house,"Images")
		o.LoadRelated(&house,"Facilities")
		houseList = append(houseList,house.To_house_info())
	}
	beego.Info("houseList",houseList)

	//7.返回总页数，当前页数，房屋信息数据
	rsp.TotalPage = int64(totalPage)
	rsp.CurrentPage = int64(housePage)
	rsp.Houses,_ = json.Marshal(houseList)

	return nil
}

