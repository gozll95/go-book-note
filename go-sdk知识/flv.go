package main

import (
	"fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	//设置上传到的空间
	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"
	bucket     = "zhulilei"

	key = "zhukey"
	//设置转码参数
	fops = "avthumb/mp4/s/640x360/vb/1.25m"

	//设置转码用的队列
	pipeline = "0.default"
)

type PutRet struct {
	Hash         string `json:"hash"`
	Key          string `json:"key"`
	PersistentId string `json:"persistentId"`
}

func main() {
	conf.ACCESS_KEY = ACCESS_KEY
	conf.SECRET_KEY = SECRET_KEY

	c := kodo.New(0, nil)

	policy := &kodo.PutPolicy{
		Scope:              bucket + ":" + key,
		Expires:            3600,
		InsertOnly:         1,
		PersistentOps:      fops,
		PersistentPipeline: pipeline,
	}

	// 生成一个上传token
	token := c.MakeUptoken(policy)
	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	var ret PutRet
	// 设置上传文件的路径
	filepath := "/Users/flower/Desktop/v1.mp4"
	//调用PutFile方式上传，这里的key需要和上传指定的key一致
	res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
	// 打印返回的信息
	fmt.Println(ret)
	// 打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}
}

/*
InsertOnly是什么意思
io.Put failed: no such pipeline
*/
