package main

import "fmt"

func main() {

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
