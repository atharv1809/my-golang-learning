package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in Golang")

	content := "This is the content going into the file."
	file, err := os.Create("./myFile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:", length)
	readFile("./myFile.txt")
	defer file.Close()
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName) // ioutil.ReadFile(fileName) was previously used but newer version of Golang uses os.Read(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text inside the file is:", string(dataByte))
}
