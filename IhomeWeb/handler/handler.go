package handler

import (
	"context"
	"encoding/json"
	"net/http"
	GETAREA "renting/GetArea/proto/example"
	GETIMAGECD "renting/GetImageCd/proto/example"
	GETSMSCD "renting/GetSmscd/proto/example"
	POSTRET "renting/PostRet/proto/example"
	GETSESSION "renting/GetSession/proto/example"
	POSTLOGIN "renting/PostLogin/proto/example"
	DELETESESSION "renting/DeleteSession/proto/example"
	GETUSERINFO "renting/GetUserInfo/proto/example"
	POSTAVATAR "renting/PostAvatar/proto/example"
	POSTUSERAUTH "renting/PostUserAuth/proto/example"
	GETUSERHOUSES "renting/GetUserHouses/proto/example"
	POSTHOUSES "renting/PostHouses/proto/example"
	PUTUSERINFO "renting/PutUserInfo/proto/example"
	POSTHOUSESIMAGE "renting/PostHousesImage/proto/example"
	GETHOUSEINFO "renting/GetHouseInfo/proto/example"
	GETINDEX "renting/GetIndex/proto/example"
	GETHOUSES "renting/GetHouses/proto/example"
	POSTORDERS "renting/PostOrders/proto/example"
	GETUSERORDER "renting/GetUserOrder/proto/example"
	PUTORDERS "renting/PutOrders/proto/example"
	PUTCOMMENT "renting/PutComment/proto/example"


	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"	

	"image"
	"image/png"
	"regexp"
	"renting/IhomeWeb/models"
	"renting/IhomeWeb/utils"

	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
		"io/ioutil"
		"strconv"
)

/*
func ExampleCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
*/
//获取地区信息
func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("请求地区信息 GetArea api/v1.0/areas")

	//创建服务，获取句柄
	server := grpc.NewService()
	//服务初始化
	server.Init()

	// 调用服务，返回句柄
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea", server.Client())
	// 调用服务，返回数据
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接收数据
	//准备接收切片
	area_list := []models.Area{}

	//循环接收数据
	for _, value := range rsp.Data {
		tem := models.Area{Id: int(value.Aid), Name: value.Aname}
		area_list = append(area_list, tem)
	}

	//返回给前端的另外两个数据
	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.Errmsg,
		"data":   area_list,
	}

	//回传数据的时候要设置数据格式
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//获取验证码图片（/api/v1.0/imagecode/:uuid 带冒号的路由用最后一个参数）
func GetImageCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取验证码图片 GetImageCd /api/v1.0/imagecode/:uuid")

	//创建服务
	server := grpc.NewService()
	server.Init()

	// 调用服务
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd", server.Client())

	//获取uuid
	uuid := ps.ByName("uuid")
	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid: uuid,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接收图片信息的图片格式
	var img image.RGBA
	//接收图片
	img.Stride = int(rsp.Stride)
	img.Pix = []uint8(rsp.Pix)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)

	var image captcha.Image
	image.RGBA = &img

	//将图片发送给浏览器
	png.Encode(w, image)

}

//获取短信验证码
func GetSmscd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取短信验证码 GetSmscd /api/v1.0/smscode/:mobile")
	//通过传入参数的URL下的Query，获取url的带参
	beego.Info(r.URL.Query())
	//map[id:[3aafb938-e1f2-4279-a372-1878cc6df9cd] text:[2901]]
	test := r.URL.Query()["text"][0]
	id := r.URL.Query()["id"][0]
	mobile := ps.ByName("mobile")

	//通过正则进行手机号的判断
	//创建正则条件
	mobileReg := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
	//通过条件判断字符串是否匹配规则,返回true或false
	b := mobileReg.MatchString(mobile)
	//如果手机号不匹配，直接返回错误，不调用服务
	if b == false {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//创建并初始化服务
	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := GETSMSCD.NewExampleService("go.micro.srv.GetSmscd", server.Client())
	rsp, err := exampleClient.GetSmscd(context.TODO(), &GETSMSCD.Request{
		Imagestr: test,
		Uuid:     id,
		Mobile:   mobile,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.Errmsg,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//注册
func PostRet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("注册 PostRet /api/v1.0/users")

	// decode the incoming request as json
	//1.接收请求数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//2.校验请求数据
	if request["mobile"].(string) == "" || request["password"].(string) == "" || request["sms_code"].(string) == "" {
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//3.服务创建
	server := grpc.NewService()
	server.Init()

	// call the backend service
	//4.调用注册服务，将数据发送给注册服务
	exampleClient := POSTRET.NewExampleService("go.micro.srv.PostRet", server.Client())
	rsp, err := exampleClient.PostRet(context.TODO(), &POSTRET.Request{
		Mobile: request["mobile"].(string),
		Password:request["password"].(string),
		SmsCode:request["sms_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.设置cookie
	//读取cookie  统一cookie  userlogin
	cookie,err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		//创建cookie对象
		cookie := http.Cookie{Name:"userlogin",Value:rsp.SessionId,Path:"/",MaxAge:3600}
		//对浏览器的cookie进行设置
		http.SetCookie(w,&cookie)
	}

	//6.准备回传数据，发送给前端
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//获取session信息
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取session信息 GetSession /api/v1.0/session")
	//1.获取数据（cookie）
	cookie,err := r.Cookie("userlogin")
	//2.校验数据
	//获取cookie失败，说明用户未登录
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	server := grpc.NewService()
	server.Init()


	// call the backend service
	exampleClient := GETSESSION.NewExampleService("go.micro.srv.GetSession", server.Client())
	rsp, err := exampleClient.GetSession(context.TODO(), &GETSESSION.Request{
		SessionId:cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]string)
	data["name"] = rsp.UserName

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")


	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//登录
func PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("登录 PostLogin /api/v1.0/sessions")

	//1.接收前端发送的数据
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//2.校验数据
	if request["mobile"].(string) == ""|| request["password"].(string) == "" {
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//7.设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//8.返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//3.创建grpc服务
	server := grpc.NewService()
	server.Init()


	//4.调用服务
	// call the backend service
	exampleClient := POSTLOGIN.NewExampleService("go.micro.srv.PostLogin", server.Client())
	rsp, err := exampleClient.PostLogin(context.TODO(), &POSTLOGIN.Request{
		Mobile: request["mobile"].(string),
		Password: request["password"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.设置cookie
	cookie,err := r.Cookie("userlogin")
	if err != nil || cookie.Value == ""{
		cookie := http.Cookie{Name:"userlogin",Value:rsp.SessionId,Path:"/",MaxAge:600}
		http.SetCookie(w,&cookie)
	}

	//6.接收服务端的返回值
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
	}

	//7.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//8.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//退出登录
func DeleteSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("退出登录 DeleteSession /api/v1.0/session")

	//1.获取SessionId
	cookie,err := r.Cookie("userlogin")
	//2.校验数据
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//7.设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//8.返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//3.创建服务
	server := grpc.NewService()
	server.Init()

	//4.调用服务
	// call the backend service
	exampleClient := DELETESESSION.NewExampleService("go.micro.srv.DeleteSession", server.Client())
	rsp, err := exampleClient.DeleteSession(context.TODO(), &DELETESESSION.Request{
		SessionId:cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.删除SessionId（cookie）
	cookie,err = r.Cookie("userlogin")
	if err == nil || cookie.Value != "" {
		cookie := http.Cookie{Name:"userlogin",Value:"",Path:"/",MaxAge:-1}
		http.SetCookie(w,&cookie)
	}

	//6.设置返回值
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
	}

	//7.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//8.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//获取用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取用户信息 GetUserInfo /api/v1.0/user")

	//1.获取sessionId
	cookie,err := r.Cookie("userlogin")
	//2.校验数据
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//3.创建服务
	server := grpc.NewService()
	server.Init()


	//4.调用服务
	// call the backend service
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", server.Client())
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		SessionId: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	//5.设置返回信息

	data := make(map[string]interface{})
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)


	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data": data,
	}
	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")

	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//上传头像
func PostAvatar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("上传头像 PostAvatar /api/v1.0/user/avatar")

	//1.接收数据
	f,h,err := r.FormFile("avatar")
	//2.校验数据
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	beego.Info("文件大小",h.Size)
	beego.Info("文件名",h.Filename)
	//创建一个文件大小的切片
	fileBuffer := make([]byte,h.Size)

	//将file的数据读到fileBuffer
	_,err = f.Read(fileBuffer)
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}


	//1.获取cookie
	cookie,err := r.Cookie("userlogin")
	//2.校验数据
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//3.创建服务
	server := grpc.NewService()
	server.Init()

	//4.调用服务，传递数据
	// call the backend service
	exampleClient := POSTAVATAR.NewExampleService("go.micro.srv.PostAvatar", server.Client())
	rsp, err := exampleClient.PostAvatar(context.TODO(), &POSTAVATAR.Request{
		Avatar:fileBuffer,
		SessionId:cookie.Value,
		FileSize:h.Size,
		FileExt:h.Filename,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.接收数据
	data := make(map[string]string)
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}


	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//更新用户名
func PutUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info(" 更新用户名 Putuserinfo /api/v1.0/user/name")


	//1.接收前端发送的数据
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//2.获取sessionId
	cookie,err := r.Cookie("userlogin")
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//3.创建服务
	server := grpc.NewService()
	server.Init()


	//4.调用服务，传递数据
	// call the backend service
	exampleClient := PUTUSERINFO.NewExampleService("go.micro.srv.PutUserInfo", server.Client())
	rsp, err := exampleClient.PutUserInfo(context.TODO(), &PUTUSERINFO.Request{
		SessionId: cookie.Value,
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.接收数据
	data := make(map[string]string)
	data["name"] = rsp.Name

	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}


	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//用户信息检查
func GetUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("用户信息检查 GetUserAuth /api/v1.0/user/auth")

	//1.获取sessionId
	cookie,err := r.Cookie("userlogin")
	//2.校验数据
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//3.创建服务
	server := grpc.NewService()
	server.Init()


	//4.调用服务
	// call the backend service
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", server.Client())
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		SessionId: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	//5.设置返回信息

	data := make(map[string]interface{})
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)


	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data": data,
	}
	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")

	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//实名认证
func PostUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("实名认证 PostUserAuth /api/v1.0/user/auth")

	//1.接收数据
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//获取sessionId
	cookie,err := r.Cookie("userlogin")
	//2.校验数据
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//3.创建服务
	server := grpc.NewService()
	server.Init()

	//4.调用服务，传递数据
	// call the backend service
	exampleClient := POSTUSERAUTH.NewExampleService("go.micro.srv.PostUserAuth", server.Client())
	rsp, err := exampleClient.PostUserAuth(context.TODO(), &POSTUSERAUTH.Request{
		SessionId: cookie.Value,
		IdCard:request["id_card"].(string),
		RealName:request["real_name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.设置返回数据
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
	}

	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//获取当前用户已发布房源信息
func GetUserHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取当前用户已发布房源信息 GetUserHouses /api/v1.0/user/houses")

	//1.获取sessionId
	cookie,err := r.Cookie("userlogin")
	//2.校验数据
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//3.创建服务
	server := grpc.NewService()
	server.Init()

	//4.调用服务，传递数据
	// call the backend service
	exampleClient := GETUSERHOUSES.NewExampleService("go.micro.srv.GetUserHouses", server.Client())
	rsp, err := exampleClient.GetUserHouses(context.TODO(), &GETUSERHOUSES.Request{
		SessionId: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.处理服务端返回的数据
	//将服务端返回的二进制流解码到切片中
	house_list := []models.House{}
	json.Unmarshal(rsp.Mix,&house_list)

	var houses []interface{}
	//遍历返回的完整房屋信息
	for _, value := range house_list {
		//获取到有用的放到切片中
		houses = append(houses,value.To_house_info())
	}

	//创建一个data的map
	data := make(map[string]interface{})
	data["houses"] = houses

	//6.设置返回数据
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}

	//7.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//8.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//发布房源
func PostHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("发布房源 PostHouses api/v1.0/houses")

	//1.将前端发送的数据解析成二进制流
	//body就是一个json的二进制流
	body,_ := ioutil.ReadAll(r.Body)

	//2.获取sessionId
	cookie,err := r.Cookie("userlogin")
	//3.校验数据
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//4.创建服务
	server := grpc.NewService()
	server.Init()


	//5.调用服务，传递数据
	// call the backend service
	exampleClient := POSTHOUSES.NewExampleService("go.micro.srv.PostHouses", server.Client())
	rsp, err := exampleClient.PostHouses(context.TODO(), &POSTHOUSES.Request{
		SessionId: cookie.Value,
		Body:body,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	//6.设置返回数据

	data := make(map[string]string)
	data["house_id"] = rsp.HouseId

	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}


	//7.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//8.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//上传房源图片信息
func PostHousesImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("上传房源图片信息 PostHousesImage /api/v1.0/houses/:id/images")

	//1.获取数据
	//2.校验数据
	//1.1获取房屋id
	id_str := ps.ByName("id")
	id,_ := strconv.Atoi(id_str)
	if id_str == "" || id <= 0{
		beego.Info("id错了")
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//1.2获取cookie
	cookie,err := r.Cookie("userlogin")
	if err != nil || cookie.Value == ""{
		beego.Info("cookie错了")
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//1.3获取图片
	f,h,err := r.FormFile("house_image")
	if err != nil{
		beego.Info("图片错了")
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//二进制的空间用来存储文件
	fileBuffer := make([]byte,h.Size)
	//将文件读取到filebuffer里
	_,err = f.Read(fileBuffer)
	if err != nil{
		beego.Info("读取错了")
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}


	//3.创建服务
	server := grpc.NewService()
	server.Init()

	//4.调用服务，传递数据
	// call the backend service
	exampleClient := POSTHOUSESIMAGE.NewExampleService("go.micro.srv.PostHousesImage", server.Client())
	rsp, err := exampleClient.PostHousesImage(context.TODO(), &POSTHOUSESIMAGE.Request{
		SessionId: cookie.Value,
		Image:fileBuffer,
		Id:id_str,
		FileSize:h.Size,
		FileName:h.Filename,

	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.设置返回数据
	data := make(map[string]string)
	data["url"] = utils.AddDomain2Url(rsp.Url)

	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}

	//设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//获取房源详细信息
func GetHouseInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")

	//1.获取数据
	//2.校验数据
	//1.1获取房屋id
	id_str := ps.ByName("id")
	id,_ := strconv.Atoi(id_str)
	if id_str == "" || id <= 0{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//1.2获取cookie
	cookie,err := r.Cookie("userlogin")
	if err != nil || cookie.Value == ""{
		// we want to augment the response
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}

		//设置回传数据的数据格式
		w.Header().Set("Content-Type", "application/json")
		//返回数据给前端
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//3.创建服务
	server := grpc.NewService()
	server.Init()

	//4.调用服务，传递数据
	// call the backend service
	exampleClient := GETHOUSEINFO.NewExampleService("go.micro.srv.GetHouseInfo", server.Client())
	rsp, err := exampleClient.GetHouseInfo(context.TODO(), &GETHOUSEINFO.Request{
		SessionId: cookie.Value,
		Id:id_str,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.设置返回数据
	data := make(map[string]interface{})
	house := models.House{}
	json.Unmarshal(rsp.HouseData,&house)
	data["house"] = house.To_one_house_desc()
	data["user_id"] = int(rsp.UserId)

	beego.Info(data["house"])

	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}

	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//获取首页轮播图
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取首页轮播图 GetIndex /api/v1.0/house/index")

	server :=grpc.NewService()
	server.Init()

	exampleClient := GETINDEX.NewExampleService("go.micro.srv.GetIndex", server.Client())
	rsp, err := exampleClient.GetIndex(context.TODO(),&GETINDEX.Request{})
	if err != nil {
		beego.Info(err)
		http.Error(w, err.Error(), 502)
		return
	}
	data := []interface{}{}
	json.Unmarshal(rsp.Max,&data)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  utils.RECODE_OK,
		"errmsg": utils.RecodeText(utils.RECODE_OK),
		"data":data,
	}

	//回传数据的时候要设置数据格式
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//搜索房源
func GetHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("搜索房源 GetHouses /api/v1.0/houses")

	//1.获取数据
	//地区编号
	aid := r.URL.Query()["aid"][0]
	//开始时间
	sd := r.URL.Query()["sd"][0]
	//结束时间
	ed := r.URL.Query()["ed"][0]
	//条件
	sk := r.URL.Query()["sk"][0]
	//页数
	p := r.URL.Query()["p"][0]

	//2.创建服务
	server :=grpc.NewService()
	server.Init()

	//3.调用服务，传递数据
	// call the backend service
	exampleClient := GETHOUSES.NewExampleService("go.micro.srv.GetHouses", server.Client())
	rsp, err := exampleClient.GetHouses(context.TODO(), &GETHOUSES.Request{
		Aid:aid,
		Sd:sd,
		Ed:ed,
		Sk:sk,
		P:p,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//4.设置返回数据
	houses := []interface{}{}
	json.Unmarshal(rsp.Houses,&houses)

	data := make(map[string]interface{})
	data["current_page"] = rsp.CurrentPage
	data["houses"] = houses
	data["total_page"] = rsp.TotalPage

	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data":data,
	}

	//5.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")

	//6.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//发布订单
func PostOrders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("PostOrders  发布订单 /api/v1.0/orders")

	//1.接收数据
	//把post传来的数据转化为二进制
	body,_ := ioutil.ReadAll(r.Body)
	userLogin,err := r.Cookie("userlogin")
	//2.校验数据
	if err != nil||userLogin.Value==""{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}
	//3.创建服务
	service := grpc.NewService()
	service.Init()

	//4.调用服务，传递数据
	// call the backend service
	exampleClient := POSTORDERS.NewExampleService("go.micro.srv.PostOrders", service.Client())
	rsp, err := exampleClient.PostOrders(context.TODO(), &POSTORDERS.Request{
		SessionId:userLogin.Value,
		Body:body,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.设置返回数据
	data := make(map[string]string)
	data["order_id"] = rsp.OrderId
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}

	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

//get 查看房东/租客订单信息请求
func GetUserOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("/api/v1.0/user/orders   GetUserOrder 获取订单 ")

	//1.获取数据（role和sessionId）
	role := r.URL.Query()["role"][0]
	userlogin,err:=r.Cookie("userlogin")

	//2.校验数据
	if role == ""{
		resp := map[string]interface{}{
			"errno": utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}



	if err != nil{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	//3.创建服务
	server :=grpc.NewService()
	server.Init()
	//4.调用服务，传递数据
	// call the backend service
	exampleClient := GETUSERORDER.NewExampleService("go.micro.srv.GetUserOrder", server.Client())
	rsp, err := exampleClient.GetUserOrder(context.TODO(), &GETUSERORDER.Request{
		SessionId:userlogin.Value,
		Role:role,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.设置返回数据
	data := make(map[string]interface{})
	orderList := []interface{}{}

	json.Unmarshal(rsp.Orders,&orderList)
	data["orders"] = orderList

	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data": data,
	}

	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//put房东同意/拒绝订单
func PutOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("/api/v1.0/orders/:id/status  PutOrders 房东同意/拒绝订单 ")

	//1.接收请求携带的数据
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//2.校验数据
	if request["action"].(string) == "" || ps.ByName("id") == "" {
		resp := map[string]interface{}{
			"errno": utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_OK),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 502)
			beego.Info(err)
			return
		}
		return
	}

	//2.获取sessionId
	userlogin,err:=r.Cookie("userlogin")
	if err != nil{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 502)
			beego.Info(err)
			return
		}
		return
	}
	//3.创建服务
	server:=grpc.NewService()
	server.Init()

	//4.调用服务，传递数据
	// call the backend service
	exampleClient := PUTORDERS.NewExampleService("go.micro.srv.PutOrders", server.Client())
	rsp, err := exampleClient.PutOrders(context.TODO(), &PUTORDERS.Request{
		SessionId:userlogin.Value,
		Action:request["action"].(string),
		OrderId:ps.ByName("id"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.接收数据
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
	}

	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//PUT 用户评价订单信请求
func PutComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("PutComment  用户评价 /api/v1.0/orders/:id/comment")

	//1.接收请求携带的数据
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//2.校验数据
	if request["comment"].(string) == "" || ps.ByName("id") == "" {
		resp := map[string]interface{}{
			"errno": utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_OK),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 502)
			beego.Info(err)
			return
		}
		return
	}

	//2.获取sessionId
	userlogin,err:=r.Cookie("userlogin")
	if err != nil{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 502)
			beego.Info(err)
			return
		}
		return
	}
	//3.创建服务
	server:=grpc.NewService()
	server.Init()

	//4.调用服务，传递数据
	// call the backend service
	exampleClient := PUTCOMMENT.NewExampleService("go.micro.srv.PutComment", server.Client())
	rsp, err := exampleClient.PutComment(context.TODO(), &PUTCOMMENT.Request{
		SessionId:userlogin.Value,
		Comment:request["comment"].(string),
		OrderId:ps.ByName("id"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//5.接收数据
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
	}

	//6.设置回传数据的数据格式
	w.Header().Set("Content-Type", "application/json")
	//7.返回数据给前端
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

