package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of Golang.")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(2025, time.September, 18, 9, 00, 00, 00, time.UTC)
	fmt.Println(createdDate)
}
