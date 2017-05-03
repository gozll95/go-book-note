package main

import (
	"fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	// 指定需要抓取到的空间
	bucket = "zhulilei"
	// 指定需要抓取的文件的url，必须是公网上面可以访问到的
	target_url = "https://image.baidu.com/search/detail?ct=503316480&z=0&ipn=d&word=二哈&hs=2&pn=1&spn=0&di=76635592660&pi=0&rn=1&tn=baiduimagedetail&is=0%2C0&ie=utf-8&oe=utf-8&cl=2&lm=-1&cs=1787296709%2C3178714066&os=475331169%2C2697898129&simid=4162438048%2C577216183&adpicid=0&lpn=0&ln=30&fr=ala&fm=&sme=&cg=&bdtype=0&oriquery=二哈&objurl=http%3A%2F%2Fimgsrc.baidu.com%2Fforum%2Fw%3D580%2Fsign%3D907442957fd98d1076d40c39113eb807%2F8af1a4e736d12f2e07917a094cc2d562843568e6.jpg&fromurl=ippr_z2C%24qAzdH3FAzdH3Fptjkw_z%26e3Bkwt17_z%26e3Bv54AzdH3FrAzdH3Fnnb09c9aac&gsm=0"
	// 指定抓取保存到空间的文件的key指
	key = "test.jpg"
)

func main() {
	conf.ACCESS_KEY = "xxxx"
	conf.SECRET_KEY = "xxxx"
	// new一个Bucket对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)
	// 调用Fetch方法
	err := p.Fetch(nil, key, target_url)
	if err != nil {
		fmt.Println("bucket.Fetch failed:", err)
	} else {
		fmt.Println("fetch success")
	}
}
