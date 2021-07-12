package filehandler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter, req *http.Request) error {
	if strings.Index(req.URL.Path, prefix) != 0 {
		return userError("path must start with" + prefix)
	}

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
