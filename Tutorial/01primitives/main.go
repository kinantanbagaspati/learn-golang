package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const Mod int = 1000000009

func main() {
	fmt.Println("Hello world")

	// Variables
	var username string = "John Doe"
	fmt.Println(username)
	fmt.Printf("Type of var: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of var: %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Type of var: %T \n", smallInt)

	var smallFloat float32 = 255.3140958160873465
	fmt.Println(smallFloat)
	fmt.Printf("Type of var: %T \n", smallFloat)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("Type of var: %T \n", defaultInt)

	var noType = "string"
	noVar := 42069
	fmt.Println(noType, noVar, Mod)

	//I/O
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a float")

	input, _ := reader.ReadString('\n')
	fmt.Println("Done, ", input)
	fmt.Printf("Type of input %T", input)

	inputFloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(inputFloat + inputFloat)
	}

	// Time
	timeNow := time.Now()
	fmt.Println(timeNow.Format("01-02-2006 15:04:05 Monday"))

	timeVar := time.Date(2001, time.May, 12, 12, 12, 12, 0, time.Local)
	fmt.Println(timeVar.Format("01-02-2006 15:04:05 Monday"))

	var toPoint int = 69
	var ptr *int = &toPoint
	fmt.Println(ptr, *ptr, toPoint)
	*ptr = *ptr * 6 + 6
	fmt.Println("End result: ", toPoint)
}