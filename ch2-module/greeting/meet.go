package greeting

import (
	"fmt"
)

// Public function
func SayMeeting() {
	fmt.Println("Nice to meet you, Go Modules.")
}

// Private function
func sayBye() {
	fmt.Println("Goodbye, See you again.")
}
