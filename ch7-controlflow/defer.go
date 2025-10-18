// go-learning/ch7-controlflow/defer.go
package main

import "fmt"

// ตัวอย่างการใช้ defer จะถูกเรียกใช้เมื่อฟังก์ชันครอบคลุมเสร็จสิ้น
func Defer() {
	defer fmt.Println("Deferred: This will run last")
	fmt.Println("This will run first")
}
