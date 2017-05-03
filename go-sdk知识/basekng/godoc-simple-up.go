package main

import (
	"fmt"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
)

var (
	//设置上传到的空间
	ACCESS_KEY = "qDaofcO6GE4O3o0a5FbK9COZR6P-0u-rbbiWj0jz"
	SECRET_KEY = "CdOykzwCI5rRtHM4RjLbpz8Jfxqnr-5Adt0kEqTJ"
	mybucket   = "zhulileiBucket"
)

func main() {
	kodo.SetMac(ACCESS_KEY, SECRET_KEY)

	zone := 0                // 您空间(Bucket)所在的区域
	c := kodo.New(zone, nil) // 用默认配置创建 Client

	bucket := c.Bucket(mybucket)
	ctx := context.Background()

	localFile := "/Users/flower/workspace/learngo/src/bee.json"
	err := bucket.PutFile(ctx, nil, "foo/bar.json", localFile, nil)
	if err != nil {
		fmt.Println("io.Put failed:", res)
		return
	}
}
