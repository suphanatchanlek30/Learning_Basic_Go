// go-learning/ch8-function/main.go

package main

import "fmt"

func main() {
	// Call the function from gofunction.go
	GoFunction1()
	GoFunction2(5, 10)
	result := GoFunction3(5, 10)
	fmt.Println("Returned Sum:", result)
	sum, diff := GoFunction4(5, 10)
	fmt.Println("Returned Sum:", sum)
	fmt.Println("Returned Difference:", diff)
}
