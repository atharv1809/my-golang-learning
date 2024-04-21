package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in Golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Kiwi"
	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Lenght of fruit list is:", len(fruitList))

	var vegList = [3]string{"Tomato", "Potato", "Corriander"}
	fmt.Println("Veg List is:", vegList)
}
