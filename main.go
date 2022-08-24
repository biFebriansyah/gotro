package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	// variable()
	// conditions()
	// perulangan()
	// arrays()
	// Slices()
	maps()

}

func variable() {
	// var firstName string
	// lastName := "Shiddiq"

	var angka8 int = 129
	var unAngka8 uint8 = 12
	var isTrue bool = true

	fmt.Println(angka8)
	fmt.Println(unAngka8)
	fmt.Println(isTrue)
}

func conditions() {
	username := "ebiebi"
	password := "abcd1234"

	if username != "ebiebi" {
		fmt.Println("username salah")

	} else if password != "abcd1234" {
		fmt.Println("password salah")

	} else {
		fmt.Println("anda login")
	}
}

func perulangan() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var angka = 0
	for {
		if angka == 10 {
			break
		}

		fmt.Println(angka)

		angka++
	}
}

func arrays() {
	var buahan = [3]string{"mangga", "rambutan", "data"}
	var numbers = [...]int{1, 2, 3}

	for i := 0; i < len(buahan); i++ {
		fmt.Println(buahan[i])
	}

	for _, value := range numbers {
		fmt.Println(value)
	}

	fmt.Println(buahan)
	fmt.Println(numbers)
}

func Slices() {
	var buahan = []string{"mangga", "rambutan", "durian", "jambu"}
	// var numbers = []int{1, 2, 3}

	var buahan1 = buahan[0:2]
	var buahan2 = buahan[1:4]

	buahan2[0] = "ikan"

	fmt.Println(buahan)
	fmt.Println(buahan1)
	fmt.Println(buahan2)
}

func maps() {
	var bulan = map[string]int{
		"januari":  1,
		"februari": 2,
	}

	var person = [...]map[string]string{
		{"name": "nada", "gender": "female"},
		{"name": "nada", "gender": "female"},
		{"name": "nada", "gender": "female"},
	}

	var _, isExist = bulan["januaris"]
	delete(bulan, "januari")

	fmt.Println(isExist)
	fmt.Println(bulan)
	fmt.Println(person)
}

func calculate(a int, b int) int {
	return a + b
}

func isLogin(username string) (string, bool) {
	return "string", true
}
