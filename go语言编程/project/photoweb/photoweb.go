package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

var template = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[tmpl] = t
	}
}

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("uploda.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename,
			http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId = r.FormValue("id")
	imagePath = UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	/*
		var listHtml string
		for _, fileInfo := range fileInfoArr {
			imgid := fileInfo.Name
			listHtml += "<li><a href=\"/view?id=" + imgid + "\">imgid</a></li>"
		}
		io.WriteString(w, "<ol>"+listHtml+"</ol>")
	*/
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name)
	}
	locals["images"] = images
	t, err := template.ParseFiles("list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, locals)
}

func safeHandler(fn http.HandlerFunc) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}

}

//静态
