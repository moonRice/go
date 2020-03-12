package main

import "fmt"

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "小李"
	m1["age"] = 18
	m1["married"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap", "篮球"}
	fmt.Println(m1)
}
