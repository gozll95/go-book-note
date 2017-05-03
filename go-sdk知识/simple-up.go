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
	bucket     = "zhulilei2"
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
		Scope: bucket,
		//设置Token过期时间
		Expires: 3600,
	}

	//生成一个上传token
	token := c.MakeUptoken(policy)
	fmt.Println("token is", token)

	//构建一个uploader
	zone := 0

	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	//设置文件上传的路径
	filepath := "/Users/flower/Desktop/me.jpeg"

	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)

	/*


			func (p Bucket) PutFileWithoutKey(
		    									ctx Context, ret interface{}, localFile string, extra *PutExtra) (err error)

			   上传一个文件。自动以文件的 hash 作为文件的访问路径（key）。
			   和 PutWithoutKey 不同的只是一个通过提供文件路径来访问文件内容，一个通过 io.Reader 来访问。

			   ctx 是请求的上下文。 ret 是上传成功后返回的数据。返回的是 PutRet 结构。可选，可以传 nil 表示不感兴趣。
			   localFile 是要上传的文件的本地路径。
			   extra 是上传的一些可选项。详细见 PutExtra 结构的描述。
	*/
	//打印返回信息
	fmt.Println(ret)

	//打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}
}

/*

AK,SK
client-policy(bucket+expire)-token

zone-uploader

token为什么
token是一堆字符串，老苟来我这里请求，我给老苟一串字符串，老苟第二次来请求，会带着这串字符串来
这个字符串是存放在cookie里的

zone(bucket所在的区域)/bucket/key
*/

/*
ak
sk

PutRet:Hash/Key

client+policy(bucket/expire)生成token

uploader(zone)->putfile(返回PutRet)
*/

//http://ooovu1db4.bkt.clouddn.com/FiTTyM4R-GARSxNVO-xzeLS7hz6j
