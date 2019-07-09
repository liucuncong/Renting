package handler

import (
	"context"

	"github.com/micro/go-log"

	example "renting/GetArea/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"renting/IhomeWeb/models"
	//redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"encoding/json"

	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("请求地区信息 GetArea /api/v1.0/areas")
	//初始化错误码
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//1.从缓存中获取数据，如果有就发给前端
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

	//获取数据,这里我们需要定制一个key，就是用来作area查询的  area_info
	area_value := bm.Get("area_info")
	if area_value != nil{
		beego.Info("获取到地域信息缓存")

		area_map := []map[string]interface{}{}
		//将获取到的数据进行json解码操作
		json.Unmarshal(area_value.([]byte),&area_map)


		beego.Info("得到从缓存中提取的数据:",area_map)

		for _, value := range area_map {
			tmp := example.Response_Areas{Aid:int32(value["aid"].(float64)),Aname:value["aname"].(string)}
			rsp.Data = append(rsp.Data,&tmp)
		}

		//将数据发送给前段
		return nil

	}


	//2.如果没有就从mysql中查找数据
	o := orm.NewOrm()
	var areas []models.Area
	num,err := o.QueryTable("Area").All(&areas)
	if err != nil{
		beego.Info("数据库查询失败",err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	if num == 0{
		beego.Info("数据库没有数据")
		rsp.Error = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//3.将查找的数据存到缓存中
	//需要将获取到的数据转化为json
	area_json,_ := json.Marshal(areas)
	//操作redis将数据存入
	err = bm.Put("area_info",area_json,time.Second*3600)
	if err != nil{
		beego.Info("数据缓存失败",err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
	}

	//4.将查找到的数据发送给前端
	//将查询到的数据按照protobuf格式发送给web服务
	for _, value := range areas {
		tmp := example.Response_Areas{Aid:int32(value.Id),Aname:value.Name}
		rsp.Data = append(rsp.Data,&tmp)
	}

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
