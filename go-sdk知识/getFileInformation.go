package main

import (
	"fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	//设置需要操作的空间
	bucket = "zhulilei2"
	key    = "FpkCYKaZHVImxmL9eedJkMOEIFVX"

	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"
)

func main() {
	conf.ACCESS_KEY = ACCESS_KEY
	conf.SECRET_KEY = SECRET_KEY

	//new一个bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	fmt.Println("p.Conn.RSHost is", p.Conn.RSHost)
	//调用Stat方法获取文件的信息
	entry, err := p.Stat(nil, key)
	//打印列取的信息

	fmt.Println(entry)

	if err != nil {
		fmt.Println(err)
	}
}

//client->bucket->stat-->entry
