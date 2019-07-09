package handler

import (
	"context"

	"github.com/micro/go-log"

	example "renting/GetImageCd/proto/example"
	"github.com/astaxie/beego"
	"github.com/afocus/captcha"
	"image/color"
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
func (e *Example) GetImageCd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("获取验证码图片 GetImageCd /api/v1.0/imagecode/:uuid")
	
	//生成图片验证码
	//创建图片句柄
	cap := captcha.New()
	
	//设置字体
	if err := cap.SetFont("comic.ttf");err != nil {
		panic(err.Error())
	}
	
	//设置图片大小
	cap.SetSize(91,41)
	//设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//设置背景色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	//生成随机的图片验证码
	img,str := cap.Create(4,captcha.NUM)
	
	//将uuid和图片验证码的值进行缓存
	//配置缓存参数
	redis_conf := map[string]string{
		"key":utils.G_server_name,
		"conn":utils.G_redis_addr + ":" +utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	beego.Info(redis_conf)
	
	//将map转化为json
	redis_conf_js,_ := json.Marshal(redis_conf)
	
	//创建redis句柄
	bm,err := cache.NewCache("redis",string(redis_conf_js))
	if err != nil {
		beego.Info("redis数据库连接错误",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//验证码与uuid进行缓存
	bm.Put(req.Uuid,str,time.Second*300)
	
	//图片解引用
	img1 := *img
	img2 := img1.RGBA
	
	//返回错误信息
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)
	
	//将图片拆分
	rsp.Stride = int64(img2.Stride)
	rsp.Pix = []byte(img2.Pix)
	rsp.Max = &example.Response_Point{X:int64(img2.Rect.Max.X),Y:int64(img2.Rect.Max.Y)}
	rsp.Min = &example.Response_Point{X:int64(img2.Rect.Min.X),Y:int64(img2.Rect.Min.Y)}

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
