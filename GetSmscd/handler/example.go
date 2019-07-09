package handler

import (
	"context"

	"github.com/micro/go-log"

	example "renting/GetSmscd/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"renting/IhomeWeb/models"
	"encoding/json"
	//redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"reflect"
	"github.com/garyburd/redigo/redis"

	"submail_go_sdk/submail/sms"
	"strconv"
	"fmt"
	"math/rand"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetSmscd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("获取短信验证码 GetSmscd /api/v1.0/smscode/:mobile")

	//1.初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//2.验证手机号是否存在
	//创建数据库orm句柄
	o := orm.NewOrm()
	var user models.User
	//使用手机号作为查询条件
	user.Mobile = req.Mobile
	err := o.Read(&user)
	//如果能查找到用户，说明手机号已经注册了，返回错误
	if err == nil {
		beego.Info("用户已存在")
		rsp.Error = utils.RECODE_MOBILEERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//3.验证图片验证码是否正确
	//连接redis
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
	//通过uuid查找图片验证码的值进行对比
	value := bm.Get(req.Uuid)
	if value == nil{
		beego.Info("redis获取数据失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//reflect.TypeOf(value)会返回当前数据的数据类型
	beego.Info(reflect.TypeOf(value),value)
	//格式转换
	value_str,_ := redis.String(value,nil)
	if value_str != req.Imagestr {
		beego.Info("数据不匹配，图片验证码值错误")
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}


	//4.调用短信接口发送短信
	//创建随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Intn(9999) + 1001
	
	//发送短信的配置信息
	messageconfig := make(map[string]string)
	//预先创建好的appid
	messageconfig["appid"] = "29672"
	//预先获得的app的key
	messageconfig["appkey"] = "89d90165cbea8cae80137d7584179bdb"
	//加密方式默认
	messageconfig["signtype"] = "md5"

	//创建 短信 Send 接口
	submail := sms.CreateSend(messageconfig)
	//设置联系人 手机号码
	submail.SetTo(req.Mobile)
	//设置短信正文，请注意：国内短信需要强制添加短信签名，并且需要使用全角大括号 “【签名】”标识，并放在短信正文的最前面
	submail.SetContent("【IHOME】您的验证码是："+strconv.Itoa(size)+"，请在5分钟输入")
	//执行 Send 方法发送短信
	send := submail.Send()
	fmt.Println("短信 Send 接口:",send)


	//5.将短信验证码存入缓存数据库
	err = bm.Put(req.Mobile,size,time.Second*300)
	if err != nil {
		beego.Info("redis缓存短信验证码失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
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
