package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	writeFile()
	readFile("./write.txt")
}

func writeFile() {
	content := "Write file test";
	file, _ := os.Create("./write.txt")

	length, _ := io.WriteString(file, content)
	fmt.Println("length is: ", length)
	defer file.Close()
}

func readFile(filename string) {
	data, _ := ioutil.ReadFile(filename)

	fmt.Println("Readings: ", string(data));
}