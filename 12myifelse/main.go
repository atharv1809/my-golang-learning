package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to If Else")

	loginCount := 23
	var ans string

	if loginCount < 10 {
		ans = "Regular User"
	} else if loginCount >= 10 && loginCount <= 20 {
		ans = "Semi Premium User"
	} else {
		ans = "Premium User"
	}
	fmt.Println("User Catogory is:", ans)
}
