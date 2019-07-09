package handler

import (
	"context"

		example "renting/GetUserOrder/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"encoding/json"
	//redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"renting/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserOrder(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("/api/v1.0/user/orders   GetUserOrder 获取订单 ")
	//1.初始化返回值
	rsp.Error  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Error)

	//2.根据sessionId得到当前用户的user_id
	//准备连接redis信息
	//{"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
	redis_conf := map[string]string{
		"key":utils.G_server_name,
		"conn":utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	beego.Info(redis_conf)

	//将map转为json
	redis_conf_json,_ := json.Marshal(redis_conf)


	//创建redis句柄
	bm,err := cache.NewCache("redis",string(redis_conf_json))
	if err != nil {
		beego.Info("redis连接失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	userId := bm.Get(req.SessionId + "user_id")
	user_id_str,_ := redis.String(userId,nil)
	user_id,_ := strconv.Atoi(user_id_str)

	//3.创建存放订单的容器及转化后的容器
	orders := []models.OrderHouse{}
	orderList := []interface{}{}
	o := orm.NewOrm()

	//4.判断role，从数据库获取数据
	if req.Role == "landlord" {
		//角色为房东
		//先找到自己目前已经发布了哪些房子
		var houses []models.House
		_,err = o.QueryTable("house").Filter("user_id",user_id).All(&houses)
		if err != nil {
			beego.Info("数据库查询失败",err)
			rsp.Error = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(rsp.Error)
			return nil
		}
		var houseIds []int
		for _, value := range houses {
			houseIds = append(houseIds,value.Id)
		}
		_,err = o.QueryTable("order_house").Filter("house__id__in", houseIds).OrderBy("ctime").All(&orders)
		if err != nil {
			beego.Info("数据库查询失败",err)
			rsp.Error = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(rsp.Error)
			return nil
		}

	}else {
		//角色为租客
		_,err = o.QueryTable("order_house").Filter("user_id",user_id).OrderBy("ctime").All(&orders)
		if err != nil {
			beego.Info("数据库查询失败",err)
			rsp.Error = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(rsp.Error)
			return nil
		}
	}

	//5.关联查询，将数据处理后存入到容器中
	for _, value := range orders {
		o.LoadRelated(&value,"User")
		o.LoadRelated(&value,"House")
		orderList = append(orderList,value.To_order_info())
	}

	//6.将处理后的数据编码返回给客户端
	rsp.Orders,_ = json.Marshal(orderList)

	return nil
}

