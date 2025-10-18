package greeting

import (
	"fmt"

	// import internal package
	personal "github.com/suphanatchanlek30/samplego/greeting/internal"
)

// Public function
func SayGreeting() {
	// Call function จาก internal package ได้
	personal.PrintPersonal()

	fmt.Println("Hello, Welcome to Go Modules.")
	// Call private function ภายใน package เดียวกันได้
	sayBye()
}
