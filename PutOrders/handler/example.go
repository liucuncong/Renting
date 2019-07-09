package handler

import (
	"context"

		example "renting/PutOrders/proto/example"
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
	"github.com/astaxie/beego/orm"
	"renting/IhomeWeb/models"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PutOrders(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("/api/v1.0/orders/:id/status  PutOrders 房东同意/拒绝订单 ")

	//1.初始化返回值
	rsp.Error  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Error)

	//2.接收数据,获取userId,orderId,action
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
	}

	user_id := bm.Get(req.SessionId + "user_id")
	user_id_str,_ := redis.String(user_id,nil)
	userId,_ := strconv.Atoi(user_id_str)

	order_id_str := req.OrderId
	order_id,_ := strconv.Atoi(order_id_str)

	action := req.Action

	//3.查找订单表,找到该订单并确定当前订单状态是wait_accept
	o := orm.NewOrm()
	order := models.OrderHouse{}
	err = o.QueryTable("order_house").Filter("id",order_id).Filter("status",models.ORDER_STATUS_WAIT_ACCEPT).One(&order)
	if err != nil {

		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	//4.关联查询house
	_,err = o.LoadRelated(&order,"House")
	if err != nil {
		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	house := order.House
	//5.校验该订单的user_id是否是当前用户的user_id
	if house.User.Id != userId {
		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = "订单用户不匹配,操作无效"
		return nil
	}

	//6.action为accept或reject
	if action == "accept"{
		//如果是接受订单,将订单状态变成待评价状态
		order.Status = models.ORDER_STATUS_WAIT_COMMENT
		beego.Debug("action = accpet!")

	} else if action == "reject" {
		//如果是拒绝接单, 尝试获得拒绝原因,并把拒单原因保存起来
		order.Status = models.ORDER_STATUS_REJECTED
		//更换订单状态为status为reject
		reason := req.Action
		//添加评论
		order.Comment = reason
		beego.Debug("action = reject!, reason is ", reason)
	}

	_, err = o.Update(&order)
	if err != nil {
		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}


	return nil
}

