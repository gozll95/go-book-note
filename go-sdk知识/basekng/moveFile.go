package main

import (
	"fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	//设置需要操作的空间
	bucket = "zhulilei"
	key    = "FiTTyM4R-GARSxNVO-xzeLS7hz6j"

	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"

	movekey = "movekey"
)

func main() {
	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"

	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	res := p.Move(nil, key, movekey)

	// 打印返回值以及出错信息
	if res == nil {
		fmt.Println("Move success")
	} else {
		fmt.Println("Move failed:", res)
	}
}

//bad-token
