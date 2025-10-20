// channel คือช่องทางสำหรับการสื่อสารระหว่าง goroutine ในภาษา Go
// ตัวอย่างนี้แสดงการสร้าง channel เพื่อส่งข้อมูลระหว่าง goroutine
// โดย goroutine หนึ่งจะส่งข้อความเข้า channel และอีก goroutine จะรับข้อความจาก channel
// เมื่อรันโปรแกรมนี้ จะเห็นข้อความที่ถูกส่งจาก goroutine หนึ่งไปยังอีก goroutine หนึ่งผ่าน channel

package main

import (
	"fmt"
	"time"
)

func main() {
	// สร้่าง channel สำหรับส่งข้อมูลประเภท string
	ch := make(chan string)

	// เรียกใช้ goroutine เพื่อส่งข้อมูลเข้า channel
	go func() {
		time.Sleep(1 * time.Second) // จำลองการทำงานที่ใช้เวลา
		// ส่งข้อมูลเข้า channel
		ch <- "Hello from goroutine!"
	}()

	// รับข้อมูลจาก channel
	msg := <-ch
	fmt.Println(msg) // แสดงผล: Hello from goroutine!

}
