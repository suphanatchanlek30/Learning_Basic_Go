// go-learning/ch7-controlflow/forloop.go
package main

import (
	"fmt"
)

// ตัวอย่างการใช้ for loop แบบธรรมดา
func Forloop1() {
	// ตัวอย่างการใช้ for loop แบบธรรมดา
	fmt.Println("For loop แบบธรรมดา:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
}

// ตัวอย่างการใช้ for loop แบบ while
func Forloop2() {
	// ตัวอย่างการใช้ for loop แบบ while
	fmt.Println("For loop แบบ while:")
	i := 1
	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}
}

// ตัวอย่างการใช้ for loop แบบไม่มีเงื่อนไข
func Forloop3() {
	// ตัวอย่างการใช้ for loop แบบ infinite
	fmt.Println("For loop แบบ infinite:")
	for {
		fmt.Println("Infinite loop:")
		break // ใช้ break เพื่อออกจากลูป
	}
}

// ตัวอย่างการใช้ for loop แบบ range
func Forloop4() {
	// ตัวอย่างการใช้ for loop แบบ range
	fmt.Println("For loop แบบ range:")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
