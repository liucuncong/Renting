package handler

import (
	"context"

		example "renting/PostOrders/proto/example"
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
	"time"
	"renting/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostOrders(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("PostOrders  发布订单 /api/v1.0/orders")

	//1.初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//2.获取用户请求数据并校验
	body := make(map[string]interface{})
	err := json.Unmarshal(req.Body,&body)
	if err != nil || body["house_id"].(string) == "" || body["start_date"].(string) == "" || body["end_date"].(string) == ""{
		rsp.Error  =  utils.RECODE_REQERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}
	//3.获取user_id
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

	//4.确定end_date在start_data之后
	start_date_time,_ := time.Parse("2006-01-02 15:04:05",body["start_date"].(string)+" 00:00:00")
	end_date_time,_ := time.Parse("2006-01-02 15:04:05",body["end_date"].(string)+" 00:00:00")
	if end_date_time.Before(start_date_time) {
		rsp.Error  =  utils.RECODE_ROLEERR
		rsp.Errmsg  = "结束时间在开始时间之前"
		return nil
	}

	//5.得到一共入住的天数
	beego.Info(start_date_time,end_date_time)
	days := end_date_time.Sub(start_date_time).Hours()/24 + 1
	beego.Info( days)

	//6.根据house_id获取房屋信息
	house_id_str := body["house_id"].(string)
	house_id,_ := strconv.Atoi(house_id_str)
	house := models.House{Id:house_id}

	o := orm.NewOrm()
	err = o.Read(&house)
	if err != nil {
		rsp.Error  =  utils.RECODE_NODATA
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	//7.确保当前的uers_id不是房源信息所关联的user_id
	o.LoadRelated(&house,"User")
	if user_id == house.User.Id{
		rsp.Error  =  utils.RECODE_ROLEERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)

		return nil
	}

	//8.确保用户选择的房屋未被预定,日期没有冲突


	//9.添加征信步骤


	//10.封装order订单
	amount := days * float64(house.Price)
	order := models.OrderHouse{}
	order.House = &house
	user := models.User{Id:user_id}
	order.User = &user
	order.Amount = int(amount)
	order.Begin_date = start_date_time
	order.Days = int(days)
	order.End_date = end_date_time
	order.House_price = house.Price
	order.Status = models.ORDER_STATUS_WAIT_ACCEPT
	//征信
	order.Credit = false

	beego.Info(order)

	//11.将订单信息入库表中
	_,err = o.Insert(&order)
	if err != nil {
		rsp.Error  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}
	//12.返回order_id
	bm.Put(req.SessionId + "user_id", user_id_str ,time.Second*7200)


	id := strconv.Itoa(order.Id)

	rsp.OrderId = id



	return nil
}

