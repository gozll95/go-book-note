package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
func main() {
	url := "http://www.baidu.com"
	fmt.Println("URL:>", url)

	//buf1 := bytes.New
	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
*/
func main() {
	url := "http://www.baidu.com"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", string(result))
}
