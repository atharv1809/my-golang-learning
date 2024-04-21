package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {           // here range returns index value
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {  // if index is not needed you can replace it with _
	// 	fmt.Printf("Index is %v and Value is %v\n", index, day)
	// }

	rogueValue := 1

	for rogueValue < 10 {
		// if rogueValue==8 {
		// 	break   // break illustration same as C++
		//  continue  // continue illustration same as C++
		//  goto label name // goto illustration same as C++
		// }
		fmt.Println("Value is:", rogueValue)
		rogueValue++ // ++rogueValue is not exist in Golang
	}
	// labelName : // for goto
	// statement of code
}
