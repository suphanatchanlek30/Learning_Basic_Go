// go-learning/ch7-controlflow/condition.go
package main

import "fmt"

// ตัวอย่างการใช้ if
func IfCondition() {
	number := 15

	// if statement
	if number%2 == 0 {
		fmt.Println(number, "is even")
	}
}

// ตัวอย่างการใช้ if-else-if
func IfElseIfCondition() {
	number := 15

	// if-else-if statement
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else if number%3 == 0 {
		fmt.Println(number, "is divisible by 3")
	} else {
		fmt.Println(number, "is odd")
	}
}

// ตัวอย่างการใช้ switch
func SwitchCondition() {
	day := 3
	var dayName string

	switch day {
	case 1:
		dayName = "Monday"
	case 2:
		dayName = "Tuesday"
	case 3:
		dayName = "Wednesday"
	case 4:
		dayName = "Thursday"
	case 5:
		dayName = "Friday"
	case 6:
		dayName = "Saturday"
	case 7:
		dayName = "Sunday"
	default:
		dayName = "Invalid day"
	}

	fmt.Println("Day", day, "is", dayName)
}

// ตัวอย่างการใช้ switch กับหลายกรณี
func SwitchWithMultipleCases() {
	dayOfWeek := "wednesday"

	switch dayOfWeek {
	case "saturday", "sunday":
		fmt.Println(dayOfWeek, "is a weekend")
	default:
		fmt.Println(dayOfWeek, "is a weekday")
	}
}
