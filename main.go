package main

import (
	"fmt"
)

/*
   1. main() 与 init() 两个函数在定义时不能有任何的参数和返回值，且Go程序会自动调用；
*/

func init() {
	fmt.Println("对于一个go文件的 init() 调用顺序是从上到下的---1")
}
func init() {
	fmt.Println("对于一个go文件的 init() 调用顺序是从上到下的---2")
}

func main() {
	fmt.Printf("hello world ")
	fmt.Println("java")
	a := 0.1
	b := 0.2
	fmt.Println(a + b)
}
