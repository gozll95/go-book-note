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


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
mux.Handle("/list", http.HandleFunc(db.list))


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

最后，一个重要的提示：就像我们在1.7节中提到的，web服务器在一个新的协程中调用每一个handler，所以当handler获取其它协程或者这个handler本身的其它请求也可以访问的变量时一定要使用预防措施比如锁机制。我们后面的两章中讲到并发相关的知识。


锁机制！！！！
