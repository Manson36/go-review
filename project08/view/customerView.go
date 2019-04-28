package main

import "fmt"

type customerView struct {
	key int
	loop bool
}

//显示主菜单
func (this *customerView)mainMenu() {

	for {
		fmt.Println("------客户信息管理软件------")
		fmt.Println("        1.添加客户")
		fmt.Println("        2.修改客户")
		fmt.Println("        3.删除客户")
		fmt.Println("        4.客户列表")
		fmt.Println("        5.退    出")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&this.key)
		switch this.key {
		case 1:
			fmt.Println("添加客户")
		case 2:
			fmt.Println("修改客户")
		case 3:
			fmt.Println("删除客户")
		case 4:
			fmt.Println("客户列表")
		case 5:
			this.loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户管理系统")

}

func main() {

	customerView:= customerView{
		key:0,
		loop:true,
	}
	customerView.mainMenu()
}