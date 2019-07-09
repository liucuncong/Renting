package handler

import (
	"context"

	"github.com/micro/go-log"

	example "renting/PostRet/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"renting/IhomeWeb/models"
			"time"
)

type Example struct{}


// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostRet(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("注册 PostRet /api/v1.0/users")
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)
	//1.验证短信验证码
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

	//通过手机号获取到短信验证码
	sms_code := bm.Get(req.Mobile)
	if sms_code == nil {
		beego.Info("获取数据失败")
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//进行短信验证码对比
	sms_code_str,_ := redis.String(sms_code,nil)
	if sms_code_str != req.SmsCode {
		beego.Info("获取短信验证码错误")
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//2.数据存入mysql数据库
	o := orm.NewOrm()
	var user models.User
	user.Mobile = req.Mobile
	user.Password_hash = utils.Md5String(req.Password)
	user.Name = req.Mobile
	id,err := o.Insert(&user)
	if err != nil {
		beego.Info("注册数据失败")
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	beego.Info("user_id",id)

	//3.创建sessionId  (唯一，随机)
	sessionId := utils.Md5String(req.Mobile + req.Password)
	rsp.SessionId = sessionId

	//4.以sessionId为key的一部分创建session
	//name
	bm.Put(sessionId+"name",user.Mobile,time.Second*3600)
	//user_id
	bm.Put(sessionId+"user_id",user.Id,time.Second*3600)
	//手机号
	bm.Put(sessionId+"mobile",user.Mobile,time.Second*3600)

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
