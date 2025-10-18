// go-learning/ch8-function/gofunction.go

package main

import (
	"fmt"
)

// Simple function prints a message to the console
func GoFunction1() {
	fmt.Println("This is Go Function 1")
}

// แบบไม่มีการคืนค่า (ไม่มี return)
func GoFunction2(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("Sum:", sum)
}

// แบบมีการคืนค่า (มี return) 1 ตัวแปร
func GoFunction3(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

// แบบมีการคืนค่า (มี return) หลายตัวแปร 2 ตัวแปรขึ้นไป
func GoFunction4(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	diff := n1 - n2
	return sum, diff
}
