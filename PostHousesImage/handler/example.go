package handler

import (
	"context"

		example "renting/PostHousesImage/proto/example"
	"github.com/astaxie/beego"
	"renting/IhomeWeb/utils"
	"path"
	"renting/IhomeWeb/models"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostHousesImage(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("上传房源图片信息 PostHousesImage /api/v1.0/houses/:id/images")

	//1.初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//2.获取文件后缀名
	beego.Info("后缀名",path.Ext(req.FileName))
	fileExt := path.Ext(req.FileName)

	//4.将图片存储到fdfs,获得编号
	fileId,err := utils.UploadByBuffer(req.Image,fileExt[1:])
	if err !=nil{
		beego.Info("fdfs存储图片错误" ,err)
		rsp.Error = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//5.根据房子的id更新房屋图片
	house_id,_ := strconv.Atoi(req.Id)
	var house  = models.House{Id:house_id}

	o := orm.NewOrm()
	err = o.Read(&house)
	if err !=nil{
		rsp.Error  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return  nil
	}

	//判断index_image_url 是否为空
	if house.Index_image_url ==""{
		//空就把这张图片设置为主图片
		house.Index_image_url = fileId
	}

	//将该图片添加到HouseImage当中
	houseimage := models.HouseImage{House:&house,Url:fileId}
	_,err  =o.Insert(&houseimage)
	if  err !=nil{
		rsp.Error  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}

	//将该图片添加到 house的全部图片当中
	house.Images = append(house.Images,&houseimage)
	//对house表进行更新
	_ , err =o.Update(&house)
	if err !=nil{
		rsp.Error  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Error)
		return nil
	}


	//6.返回url
	rsp.Url = fileId

	return nil
}


