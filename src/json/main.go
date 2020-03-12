package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与JSON

// 把Go语言中的结构体变量 -->  json格式的字符串【序列化】
// 把json格式的字符串  -->  go【反序列化】

type person struct {
	Name string `json:"name",db:"name",ini:"name"` // 反引号
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "小王",
		Age:  9000,
	}

	// 序列化
	b, err := json.Marshal(p1)

	if err != nil {
		fmt.Printf("Marshal Failed,err:%v.", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	// 反序列化
	var p2 person
	str := `{"name":"小李","age":18}`
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能真的修改数据
	fmt.Printf("%#v", p2)
}
