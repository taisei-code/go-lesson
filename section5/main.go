package main

import (
	"fmt"
)

// 算術演算子

func main() {
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "DE")

	fmt.Println(5 - 1)
	
	fmt.Println(5 * 4)

	fmt.Println(60 / 3)

	fmt.Println(9 % 4)

	n := 0
	n += 4 // n = n + 4
	fmt.Println(n)

	n++
	fmt.Println(n) // n = n + 1

	n--
	fmt.Println(n) // n = n - 1

	// 比較演算子
	fmt.Println(1 == 1)

	fmt.Println(1 == 2)

	fmt.Println(4 <= 8)

	fmt.Println(1 >= 8)

	fmt.Println(1 < 8)

	fmt.Println(3 > 1)

	fmt.Println(true == false)

	fmt.Println(true != false)

	// 論理演算子
	fmt.Println(true && false == true)

	fmt.Println(true && true == true)
}