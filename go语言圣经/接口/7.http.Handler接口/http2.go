package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

/*
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
*/

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	/*
			/
		/list
		/list
		/list
		/listdalfhlf
		/listdalfhlf
		/price
		/price
		/price
	*/
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		/*
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such page: %s\n", req.URL)
		*/
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound) // 404
	}
}

/*
req.URL.Pat
req.URL.Query().Get("item")
w.WriteHeader(http.StatusNotFound) // 404
fmt.Fprintf(w, "no such item: %q\n", item)
%q	单引号围绕的字符字面值，由Go语法安全地转义
//像 Go 源代码中那样带有双引号的输出，使用 %q。
    fmt.Printf("%q\n", "\"string\"")


msg := fmt.Sprintf("no such page: %s\n", req.URL)
http.Error(w, msg, http.StatusNotFound) // 404
*/
