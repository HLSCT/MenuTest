package main

import "fmt"

func main() {

	obj := Constructor()
	obj.AddAtTail(1)
	obj.AddAtTail(2)
	obj.AddAtTail(3)
	obj.AddAtTail(4)
	obj.PrintList()
	obj.AddAtHead(5)
	obj.PrintList()
	obj.DeleteAtIndex(2)
	obj.PrintList()
}

func menu() {

	var cmd string
	for {
		fmt.Print("输入：")
		fmt.Scanln(&cmd)
		if cmd == "help" {
			fmt.Println("help cmd")
		} else if cmd == "quit" {
			fmt.Println("quit cmd")
		} else {
			fmt.Println("Wrong cmd")
		}
	}
}
