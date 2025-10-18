// โปรแกรม Hello World ในภาษา Go
package main

// นำเข้าแพ็กเกจ fmt สำหรับการพิมพ์ข้อความ
import (
	"fmt"
)

// ประกาศฟังก์ชันหลัก
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(20 + 10)

	// Println() จะพิมพ์บรรทัดว่าง แล้วขึ้นบรรทัดใหม่
	fmt.Println()

	// Print() จะพิมพ์ข้อความต่อเนื่องกัน
	fmt.Print("Hello, ")
	fmt.Print("World!")

	// Printf() ใช้สำหรับการจัดรูปแบบข้อความ
	fmt.Printf("Hello, %s!\n", "World") // ขึ้นบรรทัดใหม่

}
