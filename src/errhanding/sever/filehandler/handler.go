package filehandler

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, req *http.Request) error {
	//fmt.Println("全路径",req.URL.Path)
	path := req.URL.Path[len("/file/"):]
	//fmt.Println("截取路径",path)
	file, e := os.Open(path)
	if e != nil {
		return e
	}
	defer file.Close()
	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		return e
	}
	writer.Write(bytes)
	return nil
}
