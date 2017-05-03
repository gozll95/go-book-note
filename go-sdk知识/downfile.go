package main

import (
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	//指定私有空间的域名
	domain = "oop87cviw.bkt.clouddn.com"
	key    = "FpkCYKaZHVImxmL9eedJkMOEIFVX"

	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"
)

func downloadUrl(domain, key string) string {
	//调用MakeBaseUrl()方法将domain,key处理成http://domain/key的形式
	baseUrl := kodo.MakeBaseUrl(domain, key)
	policy := kodo.GetPolicy{}
	c := kodo.New(0, nil)
	return c.MakePrivateUrl(baseUrl, &policy)
}

func main() {
	conf.ACCESS_KEY = ACCESS_KEY
	conf.SECRET_KEY = SECRET_KEY

	println(downloadUrl(domain, key))
}

//http://oop87cviw.bkt.clouddn.com/FpkCYKaZHVImxmL9eedJkMOEIFVX?e=1492680761&token=qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz:1ljzVsNXnIQKT5KI3cLOCkLqWas

/*
func MakeBaseUrl(domain, key string) (baseUrl string) {

	return "http://" + domain + "/" + url.Escape(key)
}
*/
//ak,sk,domain,key,url
