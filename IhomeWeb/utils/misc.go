package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/weilaihui/fdfs_client"
	"fmt"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求

func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

	return domain_url
}

//加密函数
func Md5String(s string) string {
	//创建一个md5对象
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

func UploadByBuffer(fileBuffer []byte,fileExt string)(fileId string,err error)  {

	fd_client,err1 := fdfs_client.NewFdfsClient("./conf/client.conf")
	if err1 != nil {
		fmt.Println("创建fdfs句柄失败",err1)
		return "", err1
	}

	fd_rsp,err2 := fd_client.UploadByBuffer(fileBuffer,fileExt)
	if err2 != nil {
		fmt.Println("fdfs上传图片失败",err2)
		return "", err2
	}
	fmt.Println(fd_rsp.GroupName)
	fmt.Println(fd_rsp.RemoteFileId)
	fileId = fd_rsp.RemoteFileId

	return fileId,nil
}







