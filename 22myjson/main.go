package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // this all used to specify the format of json // see the output without this and with this
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` // moit empty means if null then ignore
}

func main() {
	fmt.Println("Welcome to JSON data in Golang.")
	encodedJson()
}
func encodedJson() {
	myCourses := []course{
		{"ReactJs", 299, "learncodeonline.in", "abc123", []string{"web-dev", "fun"}},
		{"MERN", 899, "learncodeonline.in", "def123", []string{"full-stack", "fun"}},
		{"Angular", 299, "learncodeonline.in", "abc123", nil},
	}
	finalJson1, err := json.Marshal(myCourses) // this converts oly to json
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson1)

	finalJson2, err := json.MarshalIndent(myCourses, "", "\t") // this converts using \t
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson2)

	finalJson3, err := json.MarshalIndent(myCourses, "anyprefix", "\t") // this puts prefix
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson3)
}
