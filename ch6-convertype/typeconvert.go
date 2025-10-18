// go-learning/ch6-convertype/typeconvert2.go

package main

import (
	"fmt"
	"strconv"
)

// แปลงจำนวนเต็มเป็นจำนวนทศนิยม
func TypeConvert1() {
	var i int = 42
	// แปลง i จำนวนเต็มเป็น f จำนวนทศนิยม
	var f float64 = float64(i)

	fmt.Printf("i = %d, f = %.2f\n", i, f)
}

// แปลงจำนวนทศนิยมเป็นจำนวนเต็ม
func TypeConvert2() {
	var f float64 = 3.14159
	// แปลง f จำนวนทศนิยมเป็น i จำนวนเต็ม
	var i int = int(f)

	fmt.Printf("f = %.2f, i = %d\n", f, i)
}

// แปลงจากจำนวนเต็มเป็นสตริง
func TypeConvert3() {
	var i int = 100
	// แปลง i จำนวนเต็มเป็น s สตริง
	var s string = strconv.Itoa(i)
	// Ttoa = "Integer to ASCII"

	fmt.Printf("i = %d, s = %s\n", i, s)
}

// แปลงจากสตริงเป็นจำนวนเต็ม
func TypeConvert4() {
	var s string = "256.5"
	// แปลง s สตริงเป็น i จำนวนเต็ม
	var i int
	var err error
	i, err = strconv.Atoi(s)
	// Atoi = "ASCII to Integer"

	// ตรวจสอบข้อผิดพลาดในการแปลง นิยมมาก
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	fmt.Printf("s = %s, i = %d\n", s, i)
}

// แปลงจาก bool เป็น จำนวนเต็ม (แสดงผลลัพธ์เป็น 1 หรือ 0)
func TypeConvert5() {
	var b bool = true
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	fmt.Printf("b = %t, i = %d\n", b, i)
}
