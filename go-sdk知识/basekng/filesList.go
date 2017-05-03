package main

import (
	"fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"
	bucket     = "zhulilei"
)

func main() {
	conf.ACCESS_KEY = ACCESS_KEY
	conf.SECRET_KEY = SECRET_KEY

	// new一个Bucket对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)
	// 调用List方法，第二个参数是前缀,第三个参数是delimiter,第四个参数是marker，第五个参数是列举条数
	// 可以参考 https://github.com/qiniu/api.v7/blob/f956f458351353a3a75a3a519fed4e3069f14df0/kodo/bucket.go#L131
	ListItem, _, _, err := p.List(nil, "", "", "", 100)
	if err == nil {
		fmt.Println("List success")
	} else {
		fmt.Println("List failed:", err)
	}
	// 循环遍历每个操作的返回结果
	for _, item := range ListItem {
		fmt.Println(item.Key, item.Fsize)
	}
}
