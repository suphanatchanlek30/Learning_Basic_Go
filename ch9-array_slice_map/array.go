// go-learning/ch9-array_slice_map/array.go

package main

import "fmt"

func SampleArray1() {
	// ประกาศอาเรย์ที่มีขนาด 3 ช่องเก็บข้อมูลประเภท string
	var fruits [3]string

	// กำหนดค่าให้กับแต่ละช่องในอาเรย์
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"

	// แสดงผลค่าทั้งหมดในอาเรย์
	fmt.Println("Fruits:", fruits)

	// แสดงผลค่าที่ช่องที่ 1 ของอาเรย์
	fmt.Println("Fruit at index 1:", fruits[1])

	// แสดงความยาวของอาเรย์ (จํานวนช่อง) โดยใช้คําสั่ง len
	fmt.Println("Length of fruits array:", len(fruits))

	// แสดงผลค่าทั้งหมดในอาเรย์ โดยใช้ลูป for ปกติ
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("fruits[%d] = %s\n", i, fruits[i])
	}

	// แสดงผลค่าทั้งหมดในอาเรย์ โดยใช้ลูป for range
	for i, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, fruit)
	}
}

// เรียนรู้การใช้งานอาเรย์ในภาษา Go การสร้าง array และกำหนดค่าสมาชิกต่างๆ
func SampleArray2() {
	// ประกาศอาเรย์ที่มีขนาด 5 ช่องเก็บข้อมูลประเภท string พร้อมกำหนดค่าเริ่มต้น
	var carbrands = [5]string{"One", "Two", "Three", "Four", "Five"}

	// ประกาศและกำหนดค่าให้กับอาเรย์ในบรรทัดเดียว
	numbers := [5]int{10, 20, 30, 40, 50}

	// กรณี Array ขนาด 10 ช่องแต่กำหนดค่าเพียง 5 ค่า
	var scores = [10]int{100, 90, 80, 70, 60}

	// แสดงผลค่าทั้งหมดในอาเรย์ scores
	fmt.Println("Scores:", scores)

	// แสดงผลค่าทั้งหมดในอาเรย์
	fmt.Println("Numbers:", numbers)
	fmt.Println("Car Brands:", carbrands)

	// นอกจากนั้นยังสามารถกำหนดสมาชิกใน array โดยระบุ index ได้ เช่น
	// index 0 = "1", index 2 = "2", index 4 = "4"
	var colors = [5]int{0: 1, 2: 2, 4: 4}

	// แสดงผลค่าทั้งหมดในอาเรย์ colors
	fmt.Println("Colors:", colors) // Output: Colors: [1 0 2 0 4]

	// สามารถสร้าง array ที่มีขนาดไม่แน่นอนได้โดยใช้ ... แทนขนาดของ array สามารถเพิ่มได้เรื่อยๆ
	var pets = [...]string{"Dog", "Cat", "Bird"}

	// แสดงผลค่าทั้งหมดในอาเรย์ pets
	fmt.Println("Pets:", pets)

	// เพิ่มสมาชิกใน array โดยใช้ slice
	sliceOfPets := pets[:]
	sliceOfPets = append(sliceOfPets, "Fish") // เพิ่ม "Fish" เข้าไปใน slice
	fmt.Println("Pets:", pets)

}
