package main

import (
	"fmt"
	"os"
	"os/exec"
)

type student struct {
	id   int64
	name string
}

// student类型的构造函数
func addStudents(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

// 全局变量
var (
	allStudentsINFO map[int64]*student
)

func headINFO() {
	fmt.Println("欢迎使用学生信息管理系统【人数上限：100人】")
	fmt.Println("当前版本号为 V 0.2")
}

func menue() {
	fmt.Println(`
功能菜单：
	show    >  查看所有学生信息
	add     >  新增学生【新指令】
	delete  >  删除学生

	help    >  打印功能菜单
	cls     >  清除屏幕
	new     >  查看更新日志
	bye     >  退出系统
		`)
}

func waitLogo() {
	fmt.Printf("User > ")
}

func choice(userChoice string) {
	switch userChoice {
	case "show":
		showStudentsInfo()
	case "add":
		addStudentsInfo()
	case "delete":
		deleteStudentsInfo()
	case "help":
		menue()
	case "bye":
		fmt.Println("[INFO] Gooooooood Byeeee!")
		// bye()
		os.Exit(1)
	case "":

	case "cls":
		c := exec.Command("cmd", "/c", "cls") //可以根据自己的需要修改参数，自己试试，我也不清楚
		c.Stdout = os.Stdout
		c.Run()
	case "new":
		newVersion()
	default:
		fmt.Println("[WARNING] Wrong Functions!")
	}
}

func showStudentsInfo() {
	fmt.Printf("************************\n")
	for k, v := range allStudentsINFO {
		fmt.Printf("学号：%d  姓名：%s\n", k, v.name)
		fmt.Printf("************************\n")
	}

	fmt.Println("[INFO] Search successfully!")
}

func addStudentsInfo() {
	var (
		id   int64
		name string
	)

	fmt.Printf("[INFO] 请输入学号：\n")
	waitLogo()
	fmt.Scanln(&id)
	fmt.Printf("[INFO] 请输入姓名：\n")
	waitLogo()
	fmt.Scanln(&name)

	// 调用student构造函数，新建学生信息
	newStu := addStudents(id, name)

	// 加入到map中
	allStudentsINFO[id] = newStu

	fmt.Println("[INFO] Add information successfully!")
}

func deleteStudentsInfo() {
	var (
		deleteID int64
	)

	fmt.Printf("[INFO] 请输入你要删除的学生的学号：\n")
	waitLogo()
	fmt.Scanln(&deleteID)

	delete(allStudentsINFO, deleteID)

	fmt.Println("[INFO] Delete successfully!")
}

func newVersion() {
	fmt.Println(`
更新日志：
————2020年3月9日【V 0.2】————
	1、"新增学生" 指令更改为 > add
	2、新增指令
		new     >  查看更新日志
		cls     >  清除屏幕
	3、修复BUG
		空指令不该判 "错误指令"
————2020年3月1日【V 0.1】————
	1、初代系统上线
			`)
}

func main() {
	// 1、打印菜单
	headINFO()
	menue()

	// 初始化map，人数上限100人
	allStudentsINFO = make(map[int64]*student, 100)

	for {
		waitLogo()
		// 2、等待用户选择
		var userChoice string
		fmt.Scanln(&userChoice)
		choice(userChoice)
	}

	// 3、执行对应的函数
}
