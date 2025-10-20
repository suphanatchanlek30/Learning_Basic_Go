// โปรแกรมตัวอย่างการใช้ goroutines ในการชงกาแฟ ชงชา และเสิร์ฟเครื่องดื่มพร้อมกัน
// goruoutines ช่วยให้เราสามารถทำงานหลายๆ อย่างพร้อมกันได้ในเวลาเดียวกัน
// โดยไม่ต้องรอให้งานหนึ่งเสร็จก่อนถึงจะเริ่มงานถัดไป
// ในตัวอย่างนี้ เราจะใช้ goroutines เพื่อชงกาแฟ ชงชา และเสิร์ฟเครื่องดื่มพร้อมกัน

package main

import (
	"fmt"
	"time"
)

func brewCoffee() {
	fmt.Println("Brewing coffee...")
	time.Sleep(2 * time.Second)
	fmt.Println("Coffee is ready!")
}

func makeTea() {
	fmt.Println("Making tea...")
	time.Sleep(3 * time.Second)
	fmt.Println("Tea is ready!")
}

func serveDrinks() {
	fmt.Println("Serving drinks...")
	time.Sleep(1 * time.Second)
	fmt.Println("Drinks served!")
}

func main() {

	// เริ่มต้น goroutine สำหรับชงกาแฟ
	go brewCoffee()

	// เริ่มต้น goroutine สำหรับชงชา
	go makeTea()

	// เริ่มต้น goroutine สำหรับเสิร์ฟเครื่องดื่ม
	go serveDrinks()

	// รอให้ goroutines ทำงานเสร็จ
	time.Sleep(3 * time.Second) // รอ 3 วินาทีเพื่อให้ทุกงานเสร็จ
	fmt.Println("All tasks completed.")

	// หมายเหตุ: ในโปรแกรมจริง ควรใช้ sync.WaitGroup หรือช่องทางอื่นๆ
	// เพื่อรอให้ goroutines เสร็จสมบูรณ์ แทนการใช้ time.Sleep

	// เรียนรู้เพิ่มเติมเกี่ยวกับ goroutines ใน Go ได้ที่ https://go.dev/tour/concurrency

}
