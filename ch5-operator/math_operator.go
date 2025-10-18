// ch5-operator/math_operator.go

package main

// ชื่อตัวแรกของฟังก์ชันต้องเป็นตัวพิมพ์ใหญ่ แสดงว่าสามารถเรียกใช้จาก package อื่นได้ เป็น public
func MathOperator() {
	var x int

	// การบวก
	x = 5 + 3
	println("5 + 3 =", x)

	// การลบ
	x = 5 - 3
	println("5 - 3 =", x)

	// การคูณ
	x = 5 * 3
	println("5 * 3 =", x)

	// การหาร
	x = 5 / 3
	println("5 / 3 =", x)

	// การหารเอาเศษ
	x = 5 % 3
	println("5 % 3 =", x)

	// การเพิ่มค่า ไป 1
	x++
	println("x++ =", x)

	// การลดค่า ไป 1
	x--
	println("x-- =", x)

	// การคูณและหาร
	x = 5 * 3 / 2
	println("5 * 3 / 2 =", x)
}
