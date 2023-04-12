package main

import (
	"fmt"
	"sort"
)

func main() {
	// Array
	var numList [4]int
	numList[0] = 4
	numList[1] = 2
	numList[3] = 69
	fmt.Println("numList:", numList, "with length", len(numList))

	var initList = [5]string{"Fire", "Earth", "Water"}
	fmt.Println("initList", initList, "with length", len(initList))

	// Slices
	var lyricList = []string{"Apple", "Bottom", "Jeans"}
	fmt.Printf("Type of lyricList %T\n", lyricList)
	lyricList = append(lyricList[1:2], "Boots", "Fur")
	fmt.Println(lyricList)

	sliceInt := make([]int, 3)
	sliceInt[0] = 17
	sliceInt[1] = 38
	sliceInt = append(sliceInt, 9, 10, 21)
	fmt.Println(sliceInt, sort.IntsAreSorted(sliceInt))
	sort.Ints(sliceInt)
	fmt.Println(sliceInt, sort.IntsAreSorted(sliceInt))

	var languages = []string{"Javascript", "C++", "Python", "Golang"}
	fmt.Println(append(languages[:2], languages[3:]...))

	// Maps
	languageShorts := make(map[string] string)
	languageShorts["JS"] = "Javascript"
	languageShorts["CPP"] = "C++"
	languageShorts["PY"] = "Python"
	languageShorts["GO"] = "Golang"
	languageShorts["RB"] = "Ruby"
	fmt.Println("Languages shortened: ", languageShorts)
	delete(languageShorts, "RB")
	fmt.Println("Languages shortened: ", languageShorts)
	for key, val := range languageShorts {
		fmt.Println(key, ": ", val)
	}

	// Struct
	kinantan := KTP{"Kinantan", "Sangkuriang S7", 22, false}
	fmt.Println(kinantan)
	fmt.Printf("%+v %v %v\n", kinantan, kinantan.Name, kinantan.Address)
}

type KTP struct {
	Name string
	Address string
	Age int
	IsMarried bool
}