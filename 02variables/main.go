package main

import "fmt"

const LoginToken string = "kajhdf" // this is public since first letter is capital

func main() {
	// var username string = "Atharv"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type: %T \n", username)

	// var isLoggedIn bool = false
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// var small uint8 = 255 // >=256 gives error
	// fmt.Println(small)
	// fmt.Printf("Variable is of type: %T \n", small)

	// var smallFloat float32 = 255.3463456435 // >=256 gives error
	// fmt.Println(smallFloat)
	// fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases
	// var anotherVariable int
	// fmt.Println(anotherVariable)
	// fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	// var website = "myNameIsAtharv"
	// fmt.Println(website)
	// fmt.Printf("Variable is of type: %T \n", website)

	// other type
	// temp := 900000
	// fmt.Println(temp)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
