package main

import (
	"fmt"
	"strconv"
)

// Khai báo biến có tên là name với kiểu là string, string là một chuỗi các ký tự
var name string = "viet hieu"

// Khai báo 2 biến tên là dung và sai có kiểu dữ liệu là bool, bool chỉ chứa 2 giá trị true hoặc false
var dung bool = true
var sai bool = false

// Khai báo 2 biến tên là j,k có kiểu dữ liệu là int, int là các giá trị số nguyên
var j, k int = 11, 10

// Khai báo ngắn gọn biến age và khởi tạo giá trị luôn cho nó, lúc này do 22 là giá trị số nguyên nên biến age sẽ có kiểu int, còn "Xin chào, tôi là" là một chuỗi nên biến text sẽ có kiểu string.
var age = 22
var text = "Xin chao toi la"
//khai báo biến fullnameage lấy giá trị rồi gộp 3 biến name,text, age lại thành 1 biến string
var fullnameage string = text + " " + name + ",toi nam nay " + strconv.Itoa(age) + "tuoi"
//Hoặc fmt.Println(text + " " + name + ",tôi năm nay " + strconv.Itoa(age) +" tuổi")
func main() {
	//Lưu ý: nếu biến nào khai báo nhưng không dùng đến chương trình sẽ báo lỗi. Ở đây tất cả các biến đã được hiển thị

	fmt.Println("Xin chao")
	fmt.Println(name)
	fmt.Println(dung)
	fmt.Println(sai)
	fmt.Println(j, k)
	fmt.Println(text)
	fmt.Println(age)
	fmt.Println(fullnameage)

}
