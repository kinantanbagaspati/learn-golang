package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello() {
	fmt.Println("Selamat Pagi, Siang, Sore, Malam")
}
func star(a int, b int) int {
	return a * b + a + b
}
func starAdv(params ...int) int {
	mul := 1
	sum := 0
	for i := range params{
		mul *= params[i]
		sum += params[i]
	}
	return mul + sum
}

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	myDefer()

	number := 69
	if number > 420 {
		fmt.Println("Besar")
	} else if number > 42 {
		fmt.Println("Sedang")
	} else {
		fmt.Println("Kecil")
	}
	if number := 10; number < 10 {
		fmt.Println("Kecil")
	}else{
		fmt.Println("Besar")
	}

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Roll result: ", diceNum)
	switch diceNum {
	case 1:
		fmt.Println("Satu")
		fallthrough
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
		fallthrough
	case 4:
		fmt.Println("Empat")
	case 5:
		fmt.Println("Lima")
		fallthrough
	default:
		fmt.Println("Enam")
	}

	mataAngin := []string{"Utara", "Timur", "Selatan", "Barat"}
	for idx, val := range mataAngin{
		fmt.Printf("%v : %v\n", idx, val)
	}
	whileIdx := 8
	for whileIdx >= 0 {
		if whileIdx == 2 {
			goto labelEx
		} else if whileIdx == 4 {
			whileIdx--
			continue
		}
		fmt.Printf("%v ", whileIdx)
		whileIdx--
	}

labelEx:
	fmt.Println("Kelar")
	
	hello()
	fmt.Println(star(6, 9))
	fmt.Println(starAdv(1, 4, 3, 7))

	kinantan := KTP{"Kinantan", "Sangkuriang S7", 22, false}
	kinantan.SetStatus(true)
	kinantan.GetStatus()
}

type KTP struct {
	Name string
	Address string
	Age int
	IsMarried bool
}

func (u KTP) SetStatus(newStatus bool) {
	u.IsMarried = newStatus
}
func (u KTP) GetStatus() {
	fmt.Println("Is user married: ", u.IsMarried)
}
func myDefer() {
	for i:=0; i<3; i++ {
		defer fmt.Println(i)
	}
}