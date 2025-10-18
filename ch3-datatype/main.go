// 1. ประกาศ package
package main

// 2. นำเข้าแพ็กเกจ fmt
import "fmt"

// 3. สร้างฟังก์ชัน main
func main() {
	// Data Types in Go
	//  มี 2 ประเภทหลัก ได้แก่
	// 1. ชนิดข้อมูลพื้นฐาน (Basic Data Types)
	//    - int: จำนวนเต็ม เช่น 1, 2, 3
	//    - float64: จำนวนทศนิยม เช่น 1.5, 2.75
	//    - string: ข้อความ เช่น "Hello", "GoLang"
	//    - bool: ค่าความจริง เช่น true, false
	// 2. ชนิดข้อมูลประกอบ (Composite Data Types)
	//    - array: อาร์เรย์ เช่น [3]int{1, 2, 3}
	//    - slice: สไลซ์ เช่น []string{"a", "b", "c"}
	//    - map: แผนที่ เช่น map[string]int{"one": 1, "two": 2}
	//    - struct: โครงสร้าง เช่น type Person struct { Name string; Age int }
	//    - interface: อินเทอร์ฟเซส เช่น type Shape interface { Area() float64 }

	// ประกาศตัวแปรชนิดข้อมูลจำนวนเต็ม
	var a int = 10
	var b int = 20
	// ประกาศตัวแปรชนิดข้อมูลทศนิยม
	var c float64 = 30.5
	var d float64 = 40.5
	// ประกาศตัวแปรชนิดข้อมูลสตริง
	var e string = "Hello"
	var f string = "World"
	// ประกาศตัวแปรชนิดข้อมูลบูลีน
	var g bool = true
	var h bool = false
	// แสดงผลลัพธ์ของตัวแปรทั้งหมด
	fmt.Println("Integer a:", a)
	fmt.Println("Integer b:", b)
	fmt.Println("Float c:", c)
	fmt.Println("Float d:", d)
	fmt.Println("String e:", e)
	fmt.Println("String f:", f)
	fmt.Println("Boolean g:", g)
	fmt.Println("Boolean h:", h)

	// ประกาศแบบสั้น
	x := 50
	y := 60.5
	z := "GoLang"
	w := true
	i, j := 100, 200

	// แสดงผลลัพธ์ของตัวแปรแบบสั้น
	fmt.Println("Short Integer x:", x)
	fmt.Println("Short Float y:", y)
	fmt.Println("Short String z:", z)
	fmt.Println("Short Boolean w:", w)
	fmt.Println("Multiple Integers i and j:", i, j)

	// Escape Sequences (ตัวอักษรพิเศษ)
	// \n - New Line (ขึ้นบรรทัดใหม่)
	// \t - Tab (ช่องว่าง)
	// \\ - Backslash (เครื่องหมายแบ็คสแลช)
	// \" - Double Quote (เครื่องหมายคำพูดคู่)
	// \' - Single Quote (เครื่องหมายคำพูดเดี่ยว)
	// \r - Carriage Return (กลับไปต้นบรรทัด)
	// \b - Backspace (ลบตัวอักษรข้างหน้า)
	// \f - Form Feed (หน้ากระดาษใหม่)

}
