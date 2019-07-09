package handler

import (
	"context"

		example "renting/GetHouseInfo/proto/example"
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
	"fmt"
	"renting/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetHouseInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")

	//1.初始化返回值
	rsp.Error  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Error)

	//2.连接redis
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

	//3.通过sessionId拼接key，获取user_id，用于返回
	user_id := bm.Get(req.SessionId + "user_id")
	user_id_str,_ := redis.String(user_id,nil)
	userId,_ := strconv.Atoi(user_id_str)


	//4.通过房屋id从缓存数据库中获取到当前房屋的数据
	house_id,_ := strconv.Atoi(req.Id)
	house_info_key := fmt.Sprintf("house_info_%s",house_id)
	house_info_value := bm.Get(house_info_key)
	if  house_info_value != nil{
		rsp.UserId = int64(userId)
		rsp.HouseData = house_info_value.([]byte)
	}

	//5.通过房屋id从数据库中获取到当前房屋的数据
	house := models.House{Id:house_id}
	o := orm.NewOrm()
	o.Read(&house)

	//6.关联查询 area user images fac等表
	o.LoadRelated(&house,"Area")
	o.LoadRelated(&house,"User")
	o.LoadRelated(&house,"Images")
	o.LoadRelated(&house,"Facilities")

	//7.将查询到的结果存储到缓存当中
	housemix ,err := json.Marshal(house)
	bm.Put(house_info_key,housemix,time.Second*3600)

	//8.返回数据
	rsp.UserId = int64(userId)
	rsp.HouseData = housemix

	return nil
}

