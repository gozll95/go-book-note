package main

import (
	"fmt"
	"os"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"

	bucket = "zhulilei"
	key    = "FiTTyM4R-GARSxNVO-xzeLS7hz6j"

	callbackurl  = "http://ooovu1db4.bkt.clouddn.com/callback"
	callbackbody = `{"key":$(key), "hash":$(etag),"filesize":$(fsize)}`
)

type PutRet struct {
	Hash     string `json:"hash"`
	Key      string `json:"key"`
	Filesize int    `json:"filesize"`
}

func main() {
	conf.ACCESS_KEY = ACCESS_KEY
	conf.SECRET_KEY = SECRET_KEY

	c := kodo.New(0, nil)

	policy := &kodo.PutPolicy{
		Scope:        bucket + ":" + key,
		Expires:      3600,
		CallbackUrl:  callbackurl,
		CallbackBody: callbackbody,
	}

	token := c.MakeUptoken(policy)

	fmt.Fprint(os.Stdout, " token is \n", token)

	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	// 设置上传文件的路径
	filepath := "/Users/flower/Desktop/erha2.jpg"
	// 调用PutFile方式上传，这里的key需要和上传指定的key一致
	res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
	// 打印返回的信息
	fmt.Println(ret)
	// 打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}

}

//客户端post callbackbody-->七牛的服务器-->根据policy构造出post body-》回调的服务器
