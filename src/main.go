package main

import writeFile "lib/errhanding/defer"

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://baidu.com"

//调用这个方法就是对接口的调用
func Post(poster Poster) {
	poster.Post("http://baidu.com", map[string]string{
		"name": "baidu",
	})
}

// 接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
	//可以定义其他方法
}

func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string{
		"key": "value",
	})
	return rp.Get(url)
}

func download(r Retriever) string {
	return r.Get("http://baidu.com")
}

func main() {
	writeFile.WriteFile("a.txt")
}
