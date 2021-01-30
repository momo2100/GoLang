package main

import "fmt"

func main() {
	fmt.Println("Hello, สวัสดีชาว โก")
	var num int
	num = 7
	fmt.Println("My num", num)

	var message string
	message = "Expecto Patronum!"

	fmt.Println(message)

	// Decare variable
	age := 41
	fmt.Println("I am ", age, " year olds.")
	fmt.Printf("I am %d year olds.", age)
	fmt.Println()

	var bDate = 9
	var bMonth = "Sep"
	var bYear = 1980
	//var bd = fmt.Sprintf("%v-%v / %v ", bDate, bMonth, bYear)
	//fmt.Println("My birthday is", bd)

	var bd string
	if bDate < 10 {
		bd = fmt.Sprintf("0%v-%v / %v ", bDate, bMonth, bYear)
	} else {
		bd = fmt.Sprintf("%v-%v / %v ", bDate, bMonth, bYear)
	}
	fmt.Println("My birthday is", bd)

	// loop
	for i := 0; i < num; i++ {
		fmt.Println("ไปกินข้าวเหอะ !!!")
	}

	i := 3
	for i < num {
		fmt.Println("กินไรดี??")
		i++
	}

	// Switch
	switch bMonth {
	case "Jan":
		fmt.Println("มกราคม")
	case "Feb":
		fmt.Println("กุมภาพันธ์")
	case "Sep":
		fmt.Println("กันยายน")
	default:
		fmt.Println("เดือนไรงะ ??")
	}

	// array
	var foods [3]string
	foods[0] = "Tom Yam"
	foods[1] = "Pad ka Praw"
	foods[2] = "KFC"

	fmt.Printf("ไปกิน %v ไหม ??", foods[2])
	fmt.Println()

	//slice
	var opt []string
	opt = append(opt, "ม่ายย มันอ้วน")
	opt = append(opt, "เลี้ยงก็ปายยย")
	fmt.Printf("..... %v", opt[1])
	fmt.Println()

	for index, fname := range foods {
		fmt.Println(index, fname)
	}

}
