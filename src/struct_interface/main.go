package main

import "fmt"

// 同一个结构体实现多个接口
// 接口嵌套
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat同时实现了move()接口和eat()接口
func (c *cat) move() {
	fmt.Println("走猫步~")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s~", food)
}

func main() {

}
