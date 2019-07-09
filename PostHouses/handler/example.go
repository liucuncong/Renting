package handler

import (
	"context"

		example "renting/PostHouses/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"encoding/json"
	//redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"renting/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostHouses(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("发布房源 PostHouses api/v1.0/houses")
	//1.初始化默认返回值
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

	//4.拼接key
	sessionId_user := sessionId + "user_id"

	//5.查询user_id
	user_id := bm.Get(sessionId_user)

	//6.转换user_id类型
	user_id_str,_ := redis.String(user_id,nil)
	id,_ :=strconv.Atoi(user_id_str)

	//7.准备插入数据库对象
	//解析对端发送过来的body
	var requset = make(map[string]interface{})
	json.Unmarshal(req.Body,&requset)
	var house models.House

	//"title":"上奥世纪中心",
	house.Title = requset["title"].(string)

	//"price":"666",
	price,_ := strconv.Atoi(requset["price"].(string))
	house.Price = price*100

	//"address":"西三旗桥东建材城1号",
	house.Address = requset["address"].(string)

	//"room_count":"2",
	room_count,_ := strconv.Atoi(requset["room_count"].(string))
	house.Room_count = room_count
	//"acreage":"60",
	acreage,_ := strconv.Atoi(requset["acreage"].(string))
	house.Acreage = acreage

	//"unit":"2室1厅",
	house.Unit = requset["unit"].(string)

	//"capacity":"3",
	capacity,_ := strconv.Atoi(requset["capacity"].(string))
	house.Capacity = capacity

	//"beds":"双人床2张",
	house.Beds = requset["beds"].(string)
	//"deposit":"200",
	deposit,_ := strconv.Atoi(requset["deposit"].(string))
	house.Deposit = deposit*100

	//"min_days":"3",
	min_days,_ := strconv.Atoi(requset["min_days"].(string))
	house.Min_days = min_days

	//"max_days":"0",
	max_days,_ := strconv.Atoi(requset["max_days"].(string))
	house.Max_days = max_days

	//"area_id":"5",
	area_id,_ := strconv.Atoi(requset["area_id"].(string))
	area := models.Area{Id:area_id}
	house.Area = &area

	//"facility":["1","2","3","7","12","14","16","17","18","21","22"]
	facility := []*models.Facility{}
	for _, value := range requset["facility"].([]interface{}) {
		id,_ := strconv.Atoi(value.(string))
		tem := models.Facility{Id:id}
		facility = append(facility,&tem)
	}

	//8.插入数据库
	user := models.User{Id:id}
	house.User = &user

	o := orm.NewOrm()
	_,err = o.Insert(&house)
	if err != nil {
		beego.Info("数据插入失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}


	//9.多对多插入
	num,err := o.QueryM2M(&house,"Facilities").Add(facility)
	if err != nil || num ==0{
		beego.Info("房屋设施多对多插入失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//10.返回user_id
	rsp.HouseId = strconv.Itoa(house.Id)

	return nil
}

