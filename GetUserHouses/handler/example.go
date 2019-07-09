package handler

import (
	"context"

		example "renting/GetUserHouses/proto/example"
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
func (e *Example) GetUserHouses(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("获取当前用户已发布房源信息 GetUserHouses /api/v1.0/user/houses")

	//1.初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//2.根据sessionId拼接key
	sessionId_user := req.SessionId + "user_id"

	//3.连接redis
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

	//4.查询user_id
	user_id := bm.Get(sessionId_user)

	//5.user_id转换格式
	user_id_str,_ :=redis.String(user_id,nil)
	id,_ := strconv.Atoi(user_id_str)

	//6.关联查询数据库，获得当前用户的房屋信息
	var houses []models.House
	o := orm.NewOrm()
	_,err = o.QueryTable("house").Filter("user__id",id).All(&houses)
	if err != nil {
		beego.Info("查询房屋数据失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}


	//7.编码成为二进制流返回
	houses_lists,_ := json.Marshal(houses)
	rsp.Mix = houses_lists

	return nil
}
