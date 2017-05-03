package main

import (
	"fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	// 设置需要操作的空间
	bucket = "zhulilei"
	// 设置需要操作的文件的key
	key = "FiTTyM4R-GARSxNVO-xzeLS7hz6j"

	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"
)

func main() {
	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"
	// new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)
	// 调用Delete方法删除文件
	res := p.Delete(nil, key)
	// 打印返回值以及出错信息
	if res == nil {
		fmt.Println("Delete success")
	} else {
		fmt.Println(res)
	}
}
