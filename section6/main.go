package main

import (
	"fmt"
)

// 関数
// func Plus(x, y int) int {
// 	return x + y
// }

// func Div(x, y int) (int, int) {
// 	q := x / y
// 	r := x % y
// 	return q, r
// }

// func Double(price int) (result int) {
// 	result = price * 2
// 	return
// }

// func Noreturn() {
// 	fmt.Println("No Return")
// 	return
// }

// 関数を返す関数
// func ReturnFunc() func() {
// 	return func() {
// 		fmt.Println("Im a Taisei")
// 	}
// }

// 関数を引数に取る関数
// func CallFunction(f func()) {
// 	f()
// }

// クロージャー
// func Later() func(string) string {
// 	var store string
// 	return func(next string) string {
// 		s := store
// 		store = next
// 		return s
// 	}
// }

// ジェネレーター
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 関数
	// i := Plus(1, 2)
	// fmt.Println(i)

	// i2, i3 := Div(9, 3)
	// fmt.Println(i2, i3)

	// i4 := Double(1000) 
	// fmt.Println(i4)

	// Noreturn()

	// 無名関数
	// f := func(x, y int) int {
	// 	return x + y
	// }
	// i := f(1,2)
	// fmt.Println(i)

	// i2 := func(x, y int) int {
	// 	return x + y
	// }(1,2)
	// fmt.Println(i2)

	// 関数を返す関数
	// f := ReturnFunc()
	// f()

	// 関数を引数に取る関数
	// CallFunction(func() {
	// 	fmt.Println(("Im a Taisei"))
	// })

	// クロージャー
	// f := Later()
	// fmt.Println(f("Hello"))
	
	// ジェネレーター
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}