package main

import (
	"fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	//设置上传到的空间
	bucket = "zhulilei"
	//设置上传文件的key
	key = "FiTTyM4R-GARSxNVO-xzeLS7hz6j"

	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"
)

//构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func main() {
	//初始化AK,SK
	conf.ACCESS_KEY = ACCESS_KEY
	conf.SECRET_KEY = SECRET_KEY

	//创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope:   bucket + ":" + key,
		Expires: 3600,
	}

	//生成一个上传token
	token := c.MakeUptoken(policy)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet

	// 设置上传文件的路径
	filepath := "/Users/flower/Desktop/二哈1.jpg"

	res := uploader.PutFile(nil, &ret, token, key, filepath, nil)

	fmt.Println(ret)

	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}

}

//覆盖上传，多了个原文件的key
// 采用uploader.PutFile
