package main

import "fmt"

func main() {
	SamplePointer1()

	fmt.Println("===================================")

	age := 25

	// แสดงอายุเดิม
	fmt.Println("อายุเดิม:", age) // 25

	// เรียกใช้ฟังก์ชันที่เปลี่ยนแปลงค่าอายุ
	changeAge(&age) // ส่ง pointer ของ age ไปให้ฟังก์ชัน

	fmt.Println("===================================")

	// กรณีใช้ pointer
	SamplePointer2()

}
