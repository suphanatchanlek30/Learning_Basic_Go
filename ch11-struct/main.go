package main

import "fmt"

// กำหนดโครงสร้าง struct ชื่อ Person
type Person struct {
	Name string
	Age  int
}

// สร้าง function sayHello สำหรับ struct Person
func (p Person) sayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// สร้าง function sayHello2 สำหรับ struct Person
func (p Person) sayHello2() string {
	return "Hello my name is " + p.Name
}

// กำหนด function ชื่อ HaveBirthday สำหรับ struct Person
// ใช้ pointer receiver เพื่อแก้ไขค่า Age ของ Person โดยตรง
func (p *Person) HaveBirthday() {
	// เพิ่มอายุขึ้น 1 ปี
	p.Age++
}

func main() {
	// สร้าง instance ของ struct Person
	p := Person{
		Name: "Alice",
		Age:  30,
	}

	// แสดงข้อมูลของ Person
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)

	// เรียกใช้ method sayHello
	p.sayHello()

	// เรียกใช้ method sayHello2 และแสดงผลลัพธ์
	fmt.Println(p.sayHello2())

	// เรียกใช้ method HaveBirthday เพื่อเพิ่มอายุ
	p.HaveBirthday()

	// แสดงข้อมูลของ Person หลังจากมีวันเกิด
	fmt.Println("After birthday:")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
