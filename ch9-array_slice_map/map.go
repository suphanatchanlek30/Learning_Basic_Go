// map คือ เก็บ key กับ value คู่กัน

package main

import "fmt"

func MapExample() {
	// การประกาศ map
	// var m map[key]value

	// การสร้าง map
	// จะต้องใช้ฟังก์ชัน make ในการสร้าง map
	var m map[string]int
	// สร้าง map แบบย่อ
	m = make(map[string]int)
	// m := make(map[string]int)

	// การเพิ่มข้อมูลใน map
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println("Map:", m)
	// การเข้าถึงข้อมูลใน map
	fmt.Println("Value for key 'two':", m["two"])

	println("-------------------------")
	// เข้าถึงสมาชิกใน map โดยใช้ key
	fmt.Println("Value for key 'one':", m["one"])

	// ตรวจสอบว่ามี key อยู่ใน map หรือไม่
	if val, ok := m["four"]; ok {
		fmt.Println("Value for key 'four':", val)
	} else {
		fmt.Println("Key 'four' not found in map")
	}

	println("-------------------------")
	// การลบข้อมูลใน map
	delete(m, "two")
	fmt.Println("Map after deleting key 'two':", m)

	println("-------------------------")
	// การวนลูปผ่าน map
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	println("-------------------------")
	// การสร้าง Map แบบลัด โดยใช้ Literal
	// key ต้องเป็นชนิดข้อมูลเดียวกันทั้งหมด
	// value ต้องเป็นชนิดข้อมูลเดียวกันทั้งหมด
	// ตัวอย่างการสร้าง map ที่มี key เป็น string และ value เป็น string
	n := map[string]string{
		"name":    "John",
		"country": "USA",
		"city":    "New York",
	}
	fmt.Println("Map n:", n)
}
