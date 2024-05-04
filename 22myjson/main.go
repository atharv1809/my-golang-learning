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
	// encodedJson()
	decodeJson()
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

func decodeJson() {
	jsonData := []byte(`
		{
			"coursename": "ReactJs",
			"Price": 299,
			"website": "learncodeonline.in",
			"tags": ["web-dev", "fun"]
		}
	`)
	var lcoCourse course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON was valid.")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v", lcoCourse)
	} else {
		fmt.Println("JSON was not valid.")
	}
	fmt.Println("")
	// some cases where you just want to ad data to key value
	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v", myData)

	// for types and to know more go to 10:53 at https://www.youtube.com/watch?v=a96veXdifys&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=32
}
