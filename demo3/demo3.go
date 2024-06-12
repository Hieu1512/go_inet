package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var a = " vd1"
	fmt.Println(a)
	var b, c string = "vd2", "vd 3"
	fmt.Println(b, c)
	var d, e int = 3, 4
	fmt.Println(d, e)

	g := true
	fmt.Println(g)

	var h int
	fmt.Println(h)

	y := 1
	fmt.Println(y)

	const n = 500

	const o = 3e20 / n
	fmt.Println(o)

	m := 12.3
	fmt.Println(int64(m))

	fmt.Println("gia tri cua m là ", math.Sin(m))

	// vog lap for
	for j := 3; j <= 9; j++ {
		fmt.Println("gia tri k", j)
	}

	k := 0
	for k < 9 {
		fmt.Println(k)
		k++
	}

	for l := 0; l < 5; l++ {
		fmt.Println("range", l)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := range "h333333" {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	// If/Else
	if 7%4 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
		if 8%4 == 0 {
			fmt.Println("8 is divisible by 4")
		} else {
			fmt.Println("no")
		}
	}
	if 8 == 2 || 5 == 0 {
		fmt.Println("đung")
	} else {
		fmt.Println("sai")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 9 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// chuyen doi Switch
	i := 3
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//Hàm lấy ngày trong tuần hiện tại
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	//hàm lấy tgian thực
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	var v [5]int
	fmt.Println("emp:", v)
	fmt.Println("get:", v[4])

	fmt.Println("len:", len(v))

	// hàm defer

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println(1)

	i, j := 42, 2701
	// Can dung nhieu Pointers
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	//dùng nhiều
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[4:6]
	fmt.Println(s)

}
