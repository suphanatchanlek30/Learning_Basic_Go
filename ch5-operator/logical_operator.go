// ch5-operator/logical_operator.go

package main

import "fmt"

func LogicalOperator() {
	var x bool

	// AND และ &&
	x = (5 > 3) && (2 < 4)
	fmt.Println("(5 > 3) && (2 < 4) =", x)

	// OR หรือ ||
	x = (5 > 3) || (2 > 4)
	fmt.Println("(5 > 3) || (2 > 4) =", x)

	// NOT ไม่ !
	x = !(5 > 3)
	fmt.Println("!(5 > 3) =", x)
}
