package main

import (
	"io"
	"net/http"
	"qbox.us/cc/config"
	"github.com/qiniu/log.v1"
	"qiniupkg.com/api.v7/kodo"
	kodoconf "qiniupkg.com/api.v7/conf"
	"sync"
	"qbox.us/runner/taskpool"
	"os"
	"path"
)

type Config struct {
	RsfHost     string    `json:"rsf_host"`
	UpHost      string    `json:"up_host"`
	IoHost      string    `json:"io_host"`
	RsHost      string    `json:"rs_host"`
	Domain      string    `json:"domain"`
	AK          string    `json:"ak"`
	SK          string    `json:"sk"`
	Bucket      string    `json:"bucket"`
	Worker      int       `json:"worker"`
	Dir         string    `json:"dir"`
}

func main() {
	var conf Config
	wg := sync.WaitGroup{}
	config.Init("f", "qbox", "listdownload.conf")
	if err := config.Load(&conf); err != nil {
		log.Fatal("config.Load failed:", err)
	}
	log.Println("conf",conf)

	kodoconf.ACCESS_KEY = conf.AK
	kodoconf.SECRET_KEY = conf.SK

	config := kodo.Config{RSFHost:conf.RsfHost, RSHost:conf.RsHost,
		UpHosts:[]string{conf.UpHost}, IoHost:conf.IoHost,
		AccessKey:conf.AK, SecretKey:conf.SK}
	bucket := conf.Bucket

	c := kodo.New(0, &config)
	p := c.Bucket(bucket)
	var listitem []kodo.ListItem
	var maker string
	var err error

	listitem ,_,maker,err = p.List(nil,"","","",1000)
	log.Info(maker)
	log.Info(err)
	log.Info(listitem)
	for maker != ""{
		log.Info("add")
		lt := listitem
		run(lt,conf.Worker,conf.Domain,conf.Dir,&wg)
		listitem ,_,maker,err = p.List(nil,"","",maker,1000)
	}
	log.Info("add")
	lt := listitem
	run(lt,conf.Worker,conf.Domain,conf.Dir,&wg)
	wg.Wait()
}

func run(listitem []kodo.ListItem,wokernumber int,domain string,dir string,wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	task := taskpool.New(wokernumber, 1)
	for _, item := range listitem {
		wg.Add(1)
		log.Info("add")
		it := item
		task.Run(func() {
			req, err := http.Get(domain + "/" + it.Key)
			log.Info(domain +"/"+ it.Key)
			if err!=nil {
				log.Error(err)
			}
			log.Info(req.StatusCode)
			os.MkdirAll(path.Dir(dir+it.Key),0775)
			file, err := os.Create(dir+it.Key)
			if err != nil {
				log.Error(err)
			}
			io.Copy(file, req.Body)
			log.Info("done",dir+it.Key)
			wg.Done()
		})
	}
	log.Info("ok")
}