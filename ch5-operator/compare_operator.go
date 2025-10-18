// ch5-operator/compare_operator.go

package main

func CompareOperator() {
	var x bool

	// เท่ากับ
	x = (5 == 3)
	println("5 == 3 =", x)

	// ไม่เท่ากับ
	x = (5 != 3)
	println("5 != 3 =", x)

	// มากกว่า
	x = (5 > 3)
	println("5 > 3 =", x)

	// น้อยกว่า
	x = (5 < 3)
	println("5 < 3 =", x)

	// มากกว่าหรือเท่ากับ
	x = (5 >= 3)
	println("5 >= 3 =", x)

	// น้อยกว่าหรือเท่ากับ
	x = (5 <= 3)
	println("5 <= 3 =", x)
}
