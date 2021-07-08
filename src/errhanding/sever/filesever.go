package main

import (
	"lib/errhanding/sever/filehandler"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, req *http.Request) error

//返回handleFunc所需要的函数
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		e := handler(writer, req)
		if e != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(e):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/file/", errWrapper(filehandler.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
