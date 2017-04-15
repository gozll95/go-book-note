package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhelloName)
	log.Fatal(http.ListenAndServe("localhost:7000", nil))
}

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	fmt.Fprintf(w, "method: %q, URL: %q, Proto:%q", r.Method, r.URL.Path, r.Proto)
// 	for k, v := range r.Form {
// 		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
// 	}

// }
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	form_data := r.FormValue("url_long")

	fmt.Fprintf(w, "form_data is %q", form_data)
}
