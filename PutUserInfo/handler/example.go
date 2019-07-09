package handler

import (
	"context"

	"github.com/micro/go-log"

	example "renting/PutUserInfo/proto/example"
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
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PutUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info(" 更新用户名 Putuserinfo /api/v1.0/user/name")
	//初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//1.获取sessionId
	sessionId := req.SessionId

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
		return nil
	}

	//3.获取user_id
	user_id := bm.Get(sessionId+"user_id")
	user_id_str,_ := redis.String(user_id,nil)
	id,_ := strconv.Atoi(user_id_str)

	//4.将Name更新到数据库
	var user models.User
	user.Id = id
	user.Name = req.Name
	o := orm.NewOrm()
	_,err = o.Update(&user,"name")
	if err != nil {
		beego.Info("用户名更新失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//5.更新缓存
	bm.Put(sessionId+"user_id",id,time.Second*600)
	bm.Put(sessionId+"name",id,time.Second*600)


	//6.返回数据
	rsp.Name = user.Name


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
