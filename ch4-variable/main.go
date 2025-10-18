package main

import "fmt"

// การประกาศตัวแปรใน Go
// แบบ Global Variable
var globalVar string = "Hello, Go!"

func main() {
	// แบบ Local Variable
	localVar := "Welcome to Go programming."

	fmt.Println(globalVar)
	fmt.Println(localVar)

	// การประกาศตัวแปรแบบหลายตัวพร้อมกัน โดยไม่การกำหนดค่าเริ่มต้น
	var (
		num   float32 = 3.14
		count int     = 10
		name  string  = "Gopher"
	)

	// print ค่า
	fmt.Println(num)
	fmt.Println(count)
	fmt.Println(name)

	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)

	var p, q, r = "A", 3.5, true
	fmt.Println(p, q, r)

	// ตัวเลขเยอะๆ ใส่ _ เพื่อให้อ่านง่ายขึ้น
	var w = 1_000_000
	fmt.Println(w)

	// ประกาศค่าคงที่ไม่เปลี่ยนแปลง
	// const ตัวแปรชื่อ ประเภทข้อมูล = ค่า
	const pi float32 = 3.14159
	fmt.Println(pi)

}
