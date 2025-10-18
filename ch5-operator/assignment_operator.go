// ch5-operator/assignment_operator.go

package main

import "fmt"

func AssignmentOperator() {
	var x int

	// การกำหนดค่า
	x = 10
	fmt.Println("x =", x)

	// การเพิ่มค่า
	x += 5
	fmt.Println("x += 5 =", x)

	// การลดค่า
	x -= 3
	fmt.Println("x -= 3 =", x)

	// การคูณค่า
	x *= 2
	fmt.Println("x *= 2 =", x)

	// การหารค่า
	x /= 4
	fmt.Println("x /= 4 =", x)

	// การหารเอาเศษ
	x %= 3
	fmt.Println("x %= 3 =", x)

	// AND บิตแล้วกำหนดค่า
	x &= 2
	fmt.Println("x &= 2 =", x)

	// OR บิตแล้วกำหนดค่า
	x |= 1
	fmt.Println("x |= 1 =", x)

	// XOR บิตแล้วกำหนดค่า
	x ^= 3
	fmt.Println("x ^= 3 =", x)

	// เลื่อนบิตซ้ายแล้วกำหนดค่า
	x <<= 1
	fmt.Println("x <<= 1 =", x)

	// เลื่อนบิตขวาแล้วกำหนดค่า
	x >>= 1
	fmt.Println("x >>= 1 =", x)
}
