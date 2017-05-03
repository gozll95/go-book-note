package main

import (
	"flag"
	"fmt"
	"os"
	//"qiniupkg.com/api.v7/conf"
	//"qiniupkg.com/api.v7/kodo"
	//"qiniupkg.com/api.v7/kodocli"
)

/*
 zhu -f up/coverUp/io/getInfo/mv/del/list/ListItems/BatchCopy
*/

func main() {
	sdkFun := flag.String("f", "up", "-f up/coverUp/io/getInfo/mv/del/list/ListItems/BatchCopy")

	flag.Parse()

	sdk_fun := *sdkFun
	switch sdk_fun {
	case "up":
		AK := flag.String("AK", "ak", "-AK xxxx")
		SK := flag.String("SK", "sk", "-SK xxxx")
		bucket := flag.String("-bucket", "", "-bucket xxxx")

		flag.Parse()
		fmt.Println(*AK, *SK, *bucket)
		//sdkUp(*AK, *SK, *bucket)
	//case  "coverUp":
	//	sdkCoverUp()
	//case "io":
	//	sdkIo()
	//case="getInfo":
	//	sdkGetInfo()
	//case "mv":
	//	sdkMv()
	//case "del":
	//	sdkDel()
	//case "List":
	//	sdkList()
	//case "ListItems":
	//	sdkListItems()
	//case "BatchCopy":
	//	sdkBatchCopy()
	default:
		fmt.Println("error! -f usage error!")
		os.Exit(1)
	}
}

/*
func sdkUp(AK, SK, bucket string) {
	//构造返回值字段
	type PutRet struct {
		Hash string `json:"hash"`
		Key  string `json:"key"`
	}

	//初始化AK,SK
	conf.ACCESS_KEY = AK
	conf.SECRET_KEY = SK

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

	//打印返回信息
	fmt.Println(ret)

	//打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}

}

*/
