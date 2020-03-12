package main

import "fmt"

// type speakHit interface {
// 	speak()
// 	// 只要实现speak()方法的变量，全部都是speakHit类型
// }

// // 引出接口的实例
// type cat struct {
// }
// type dog struct {
// }
// type person struct {
// }

// func (c cat) speak() {
// 	fmt.Println("miao miao miao~")
// }
// func (d dog) speak() {
// 	fmt.Println("wang wang wang~")
// }
// func (p person) speak() {
// 	fmt.Println("a a a~")
// }
// func da(x speakHit) {
// 	// 接收一个参数，传进来，我就打谁
// 	x.speak() // 挨打会叫
// }

// func main() {
// 	var c1 cat
// 	var d1 dog
// 	var p1 person

// 	da(c1)
// 	da(d1)
// 	da(p1)
// }

// type car interface {
// 	run()
// }

// type falali struct {
// 	brand string
// }

// func (f falali) run() {
// 	fmt.Printf("%s速度700迈~", f.brand)
// }

// type baoshijie struct {
// 	brand string
// }

// func (f baoshijie) run() {
// 	fmt.Printf("%s速度200迈~", f.brand)
// }

// func drive(c car) {
// 	c.run()
// }

// func main() {
// 	var f1 = falali{
// 		brand: "法拉利",
// 	}
// 	var f2 = baoshijie{
// 		brand: "保时捷",
// 	}

// 	drive(f1)
// 	drive(f2)
// }

// 接口的实现
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}
type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动！")
}
func (c chicken) eat(food string) {
	fmt.Printf("吃%s！", food)
}

func (c *cat) move() {
	fmt.Println("走猫步~")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s~", food)
}

func main() {
	var a1 animal
	fmt.Printf("a1=%T.", a1)
	bc := cat{
		name: "蓝猫",
		feet: 4,
	}

	// 少一个参数eat(string)
	a1 = &bc
	a1.eat("小黄鱼")
	fmt.Println(a1)

	var a2 chicken
	kfc := chicken{
		feet: 2,
	}
	a2.eat("白斩鸡")
	a2 = kfc
	fmt.Println(a2)

	fmt.Printf("a1=%T,a2=%T.", a1, a2)
	fmt.Println()
	var cc animal
	c1 := cat{"tom", 4}
	c2 := &cat{"jerry", 4}

	cc = &c1
	cc = c2
	fmt.Println(cc)
}
