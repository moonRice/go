package main

import "fmt"

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 2.8.1
func nineXnine() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %-2d  ", i, j, i*j)
		}
		fmt.Println()
	}
}

// 4.8.1
func arraySum() {
	var sum = 0
	var arraySum = [...]int{1, 3, 5, 7, 8}
	fmt.Println(arraySum)
	fmt.Printf("%T\n", arraySum)
	for i := 0; i < len(arraySum); i++ {
		sum = sum + arraySum[i]
	}
	fmt.Println(sum)
}

// 4.8.2
func findSum() {
	var sum = 8
	var arraySum = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arraySum); i++ {
		for j := i + 1; j < len(arraySum); j++ {
			if arraySum[i]+arraySum[j] == sum {
				fmt.Printf("(%d,%d)", i, j)
			}
		}
	}
}

// 结构体
type person struct {
	name   string
	age    int
	gender bool
	hobby  []string
}

// main
func main() {
	fmt.Println("Hello World!")

	// 输出10个连续数字
	//forDemo()

	// 9×9口诀表
	//nineXnine()

	// 数组求和
	//arraySum()

	// 输出下标，和为8
	//findSum()

	var p person
	p.name = "老王"
	p.age = 32
	p.gender = true
	p.hobby = []string{
		"篮球",
		"跑步",
	}
	fmt.Println(p)
	fmt.Printf("%T\n", p)
}
