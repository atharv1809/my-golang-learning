package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the Slices")
	var fruitList = []string{"Apple", "Banana", "Mango"}

	fruitList = append(fruitList, "Kiwi", "Orange")
	fmt.Println("FruitList is:", fruitList)

	fruitList = append(fruitList[1:3]) // 0 based indexing, 3 is excluded
	fmt.Println("FruitList is:", fruitList)

	scores := make([]int, 4)

	scores[0] = 1
	scores[1] = 2
	scores[2] = 3
	scores[3] = 4

	scores = append(scores, 5, 6, 7)
	fmt.Println("Scores are:", scores)

	sort.Ints(scores)
	fmt.Println("Scores are:", scores)

	var courses = []string{"Golang", "C++", "Java", "Python"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Coures are:", courses)
}
