package handler

import (
	"context"

	"github.com/micro/go-log"

	example "renting/PostLogin/proto/example"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"renting/IhomeWeb/models"
	"renting/IhomeWeb/utils"
	"encoding/json"
	//redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostLogin(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("登录 PostLogin /api/v1.0/sessions")
	//设置默认返回
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//1.查询数据
	o := orm.NewOrm()
	var user models.User
	err := o.QueryTable("user").Filter("mobile",req.Mobile).One(&user)
	if err != nil {
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//2.密码校验
	if utils.Md5String(req.Password) != user.Password_hash {
		rsp.Error = utils.RECODE_PWDERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//3.创建sessionid
	sessionId := utils.Md5String(req.Mobile+req.Password)
	//4.返回数据
	rsp.SessionId = sessionId

	//5.连接redis,拼接key，将登录信息缓存
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
	//缓存用户数据
	//name
	bm.Put(sessionId+"name",user.Name,time.Second*3600)
	//user_id
	bm.Put(sessionId+"user_id",user.Id,time.Second*3600)
	//mobile
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
