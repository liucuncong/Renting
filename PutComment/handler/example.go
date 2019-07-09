package handler

import (
	"context"

		example "renting/PutComment/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"encoding/json"
	"strconv"
	//redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/garyburd/redigo/redis"
	"renting/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PutComment(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("PutComment  用户评价 /api/v1.0/orders/:id/comment")

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

	comment := req.Comment
	//检验评价信息是否合法 确保不为空
	if comment == "" {

		rsp.Error  =  utils.RECODE_PARAMERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	//3.根据order_id找到所关联的房源信息
	//查询数据库，订单必须存在，订单状态必须为WAIT_COMMENT待评价状态
	order := models.OrderHouse{}
	o := orm.NewOrm()
	if err := o.QueryTable("order_house").Filter("id", order_id).Filter("status", models.ORDER_STATUS_WAIT_COMMENT).One(&order); err != nil {
		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	//4.关联查询order订单所关联的user信息
	if _, err := o.LoadRelated(&order, "User"); err != nil {

		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)

		return nil
	}

	//5.确保订单所关联的用户和该用户是同一个人
	if userId != order.User.Id {

		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	//6.关联查询order订单所关联的House信息
	if _, err := o.LoadRelated(&order, "House"); err != nil {

		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}
	house := order.House

	//7.更新订单和房屋数据
	//更新order的status为COMPLETE
	order.Status = models.ORDER_STATUS_COMPLETE
	order.Comment = comment

	//将房屋订单成交量+1
	house.Order_count++

	//8.将order和house更新到数据库
	if _, err := o.Update(&order, "status", "comment"); err != nil {
		beego.Error("update order status, comment error, err = ", err)


		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	if _, err := o.Update(house, "order_count"); err != nil {
		beego.Error("update house order_count error, err = ", err)


		rsp.Error  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	//9.将house_info_[house_id]的缓存key删除 （因为已经修改订单数量）
	house_info_key := "house_info_" + strconv.Itoa(house.Id)
	if err := bm.Delete(house_info_key); err != nil {
		beego.Error("delete ", house_info_key, "error , err = ", err)
	}


	return nil
}
