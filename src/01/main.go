package main

import "fmt"

//结构体是值类型
type person struct {
	name   string
	age    int
	gender bool
	hobby  []string
}

func f1(x person) {
	x.gender = false
}

func f2(y *person) {
	(*y).gender = false
}

func main() {
	var p person
	p.name = "老王"
	p.gender = true
	f1(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)

	var p2 = new(person)
	p2.name = "Aniy"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)
	fmt.Println(p2)
}
