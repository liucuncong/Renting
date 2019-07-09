package handler

import (
	"context"

		example "renting/PostUserAuth/proto/example"
	"github.com/astaxie/beego"
	//redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"renting/IhomeWeb/utils"
	"encoding/json"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"renting/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostUserAuth(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("实名认证 PostUserAuth /api/v1.0/user/auth")

	//1.初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//2.获取sessionId
	sessionId := req.SessionId
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

	//4.拼接key,获取user_id
	sessionId_user := sessionId + "user_id"
	user_id := bm.Get(sessionId_user)
	user_id_str,_ := redis.String(user_id,nil)
	id,_ :=strconv.Atoi(user_id_str)
	//5.通过usr_id更新表,将身份证号和姓名更新到表上
	var user models.User
	user.Id = id
	o := orm.NewOrm()
	err = o.Read(&user)
	if err != nil {
		beego.Info("根据id查询用户失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	user.Id_card = req.IdCard
	user.Real_name = req.RealName
	_,err = o.Update(&user,"real_name","id_card")
	if err != nil {
		beego.Info("更新身份信息失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//6.刷新下sessionId时间
	bm.Put(sessionId_user,user_id_str,time.Second*600)


	return nil
}

