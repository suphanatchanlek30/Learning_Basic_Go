package main

import (
	"fmt"
	// import package ที่ติดตั้งจาก go.mod
	"github.com/google/uuid"

	// import package ที่สร้างขึ้นเอง เอาชื่อ package ตามที่ตั้งไว้ใน go.mod ไม่สนใจชื่อไฟล์ เพื่อให้เรียก call function ได้
	"github.com/suphanatchanlek30/samplego/greeting"

	// ตั้งไว้สำหรับ import package ที่ไม่ได้ใช้งานจริงๆ แต่ต้องการให้ go.mod มี package นี้อยู่ เช่น กรณีที่ต้องการใช้ package นี้ในอนาคต
	_ "math"
)

func main() {
	fmt.Println("Hello, Go Modules!")

	// Generate a new UUID
	var id string = uuid.New().String()
	fmt.Println("Generated UUID:", id)

	// Call the function from the greeting package หลังจาก import package เรียบร้อยแล้ว เราสามารถเรียกใช้ได้เลย
	greeting.SayGreeting()
	greeting.SayMeeting()

	// ประกาศตัวแปรแบบไม่ใช้ (blank identifier) เพื่อให้โค้ดไม่ error
	_ = id
}
