package handler

import (
	"context"

		example "renting/GetUserInfo/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"encoding/json"
	//redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"renting/IhomeWeb/models"
	"strconv"
	"github.com/garyburd/redigo/redis"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("获取用户信息 GetUserInfo /api/v1.0/user")

	//1.初始化错误码
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

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
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//3.获取userId
	userId := bm.Get(req.SessionId + "user_id")
	userIdStr,_ := redis.String(userId,nil)
	id,_ := strconv.Atoi(userIdStr)

	//4.获取用户表信息
	o := orm.NewOrm()
	var user models.User
	user.Id = id
	err = o.Read(&user)
	if err != nil {
		beego.Info("数据获取失败",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}


	//5.将信息返回
	rsp.UserId = strconv.Itoa(user.Id)
	rsp.Name = user.Name
	rsp.IdCard = user.Id_card
	rsp.RealName = user.Real_name
	rsp.AvatarUrl = user.Avatar_url
	rsp.Mobile = user.Mobile

	return nil
}

