package http

import (
	"net/http"
	"github.com/iamqiutian/uploadFile/g"
)

func init() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/choose", ChooseFile)
	http.HandleFunc("/upload", UploadFile)
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(g.Config.Upload.Path))))
}

