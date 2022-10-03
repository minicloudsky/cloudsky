package main

import "fmt"

// HandlerFunc 函数 实现了Handler接口
type HandlerFunc func(k, v interface{})

func (hf HandlerFunc) Do(k, v interface{}) {
	hf(k, v)
}


type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}


func doMyStuff(k, v interface{}) {
	fmt.Println(k, v, "my stuff")
}


func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26
	Each(persons, HandlerFunc(doMyStuff))
}
