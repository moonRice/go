package main

import "fmt"

// 结构体实现”继承“
type animal struct {
	name string
}

// 给animal实现一个西东的方法
func (a animal) move() {
	fmt.Printf("%s会动。", a.name)
}

// 狗类
type dog struct {
	feet uint8
	animal
}

// 给狗实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s会汪汪汪~", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "www",
		},
	}

	d1.wang()
	d1.move()
}
