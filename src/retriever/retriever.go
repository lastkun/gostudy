package retriever

import "fmt"

// 这里实现了两个接口 并不需要表示它实现了 只需要关注方法的实现
type Retriever struct {
	Data string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Contents = %s", r.Data)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Data = form["key"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Data
}
