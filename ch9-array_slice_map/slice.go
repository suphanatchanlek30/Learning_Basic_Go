package main

import "fmt"

func SampleSlice1() {
	// ประกาศตัวแปรแบบ Slice
	var fruits []string

	// กำหนดค่าให้กับ Slice
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana", "Cherry")

	// แสดงผลค่าใน Slice
	fmt.Println("Fruits:", fruits)

	fmt.Println("------------------------------")

	// การเข้าถึงค่าภายใน Slice
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Second fruit:", fruits[1])

	fmt.Println("------------------------------")

	// การหาความยาวของ Slice
	fmt.Println("Number of fruits:", len(fruits))

	fmt.Println("------------------------------")

	// สร้าง Slice โดยใช้ literal syntax
	vegetables := []string{"Carrot", "Broccoli", "Spinach"}

	// แสดงผลค่าใน Slice
	fmt.Println("Vegetables:", vegetables)

	// print index 0-1 of vegetables slice
	fmt.Println("First vegetable:", vegetables[0:2])

	fmt.Println("------------------------------")

	// ตัด Slice
	citrus := []string{"Orange", "Lemon", "Lime", "Grapefruit"}
	slicedCitrus := citrus[1:3] // ตัดตั้งแต่ index 1 ถึง 2
	// ขนาดของ slicedCitrus จะเป็น 2 (Lemon, Lime)
	fmt.Println("Number of sliced citrus fruits:", len(slicedCitrus))
	fmt.Println("Citrus fruits:", slicedCitrus)

	fmt.Println("------------------------------")
	// การใช้ make เพื่อสร้าง Slice
	numbers := make([]int, 5) // สร้าง Slice ของ int ที่มีความยาว 5
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Numbers:", numbers)

	// หา capacity ของ Slice
	fmt.Println("Capacity of numbers slice:", cap(numbers))

}
