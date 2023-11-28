package main

import (
	"fmt"
)

// スライス

// append make len cap

// 可変長引数
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	// var sl []int
	// fmt.Println(sl)

	// var sl2 []int = []int{100, 200} 
	// fmt.Println(sl2)

	// sl3 := []string{"A", "B"}
	// fmt.Println(sl3)

	// sl4 := make([]int, 5)
	// fmt.Println(sl4)

	// sl2[0] = 1000
	// fmt.Println(sl2)

	// sl5 := []int{1, 2, 3, 4, 5}
	// fmt.Println(sl5)

	// fmt.Println(sl5[0])

	// fmt.Println(sl5[2:4])

	// fmt.Println(sl5[:2])

	// fmt.Println(sl5[2:])

	// append make len cap
	// sl := []int{100, 200}
	// fmt.Println(sl)

	// sl = append(sl, 300)
	// fmt.Println(sl)

	// sl2 := make([]int, 5)
	// fmt.Println(sl2)

	// fmt.Println(len(sl2))

	// fmt.Println(cap(sl2))

	// sl3 := make([]int, 5, 10)

	// fmt.Println(len(sl3))

	// fmt.Println(cap(sl3))

	// copy
	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000
	// fmt.Println(sl)

	// sl := []int{1, 2, 3, 4, 5}
	// sl2 := make([]int, 5, 10)
	// fmt.Println(sl2)
	// n := copy(sl2, sl)

	// fmt.Println(n, sl2)

	// sl := []string{"A", "B", "C"}
	// fmt.Println(sl)

	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// 可変長引数
	// fmt.Println(Sum(1, 2, 3))

	// マップ
	// var m = map[string]int{"A": 100, "B": 200}
	// fmt.Println(m)
	
	// m2 := map[string]int{"A": 100, "B": 200}
	// fmt.Println(m2)

	// m3 := map[int]string{
	// 	1: "A",
	// 	2: "B",
	// }
	// fmt.Println(m3)

	// m4 := make(map[int]string)
	// fmt.Println(m4)

	// m4[1] = "JAPAN"
	// m4[2] = "USA"
	// fmt.Println(m4)

	// fmt.Println(m["A"])
	// fmt.Println(m4[2])

	// s, ok := m4[3]
	// if !ok {
	// 	fmt.Println("error")
	// }
	// fmt.Println(s, ok)

	// m4[2] = "US"
	// fmt.Println(m4)

	// m4[3] = "CHINA"
	// fmt.Println(m4)

	m := map[string]int {
		"Apple": 100,
		"Banana": 200,
	}

	for _, v := range m {
		fmt.Println(v)
	}
}