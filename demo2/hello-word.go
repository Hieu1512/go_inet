package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// fmt la package chuyen xuat du lieu ra man hinh
	fmt.Println("Nhap ten cua ban vao day: ")
	// ham nhap xuat tu ban phim
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	//yeu cau ng dung nhap tu ban phim
	name, _ := reader.ReadString('\n')

	fmt.Println("Xin chao " + name)
}
