package handler

import (
	"context"

		example "renting/PostAvatar/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"path"
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
func (e *Example) PostAvatar(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("上传头像 PostAvatar /api/v1.0/user/avatar")
	//1.初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//校验size
	size := len(req.Avatar)

	if int64(size) != req.FileSize{
		beego.Info("传输数据丢失")
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//2.获取文件的后缀名
	ext := path.Ext(req.FileExt)

	//3.上传图片到fasfDFS服务器
	fileId,err := utils.UploadByBuffer(req.Avatar,ext[1:])
	if err != nil {
		beego.Info("上传失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//4.得到fileId
	beego.Info(fileId)
	//5.获取sessionId,连接redis
	sessionId := req.SessionId
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


	//6.拼接key，获取当前用户的user_id
	id := bm.Get(sessionId + "user_id")
	userId1,_ := redis.String(id,nil)
	userId,_ := strconv.Atoi(userId1)

	//7.将图片的存储地址更新到user表
	var user models.User
	o := orm.NewOrm()
	user.Id =userId
	o.Read(&user)
	user.Avatar_url = fileId
	_,err = o.Update(&user,"avatar_url")
	if err != nil {
		beego.Info("数据库更新数据失败")
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//8.回传fileId
	rsp.AvatarUrl = fileId


	return nil
}

