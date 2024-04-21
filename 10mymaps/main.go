package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	for key, value := range languages {
		fmt.Println(key, value)
	}
}
