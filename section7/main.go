package main

import (
	"fmt"
	"time"
)

// switch 型スイッチ
// func anything(a interface{}) {
// 	// fmt.Println(a)
// 	switch v := a.(type) {
// 	case string:
// 		fmt.Println(v + "!?")
// 	case int:
// 		fmt.Println(v + 10000)
// 	}
// }

// defer
// func TestDefer() {
// 	defer fmt.Println("END")
// 	fmt.Println("START")
// }

// func RunDefer() {
// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
// 	defer fmt.Println("3")
// }

// go 並行処理
func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	// 条件分岐
	// a := 1
	// if a == 2 {
	// 	fmt.Println("two")
	// } else if a == 1 {
	// 	fmt.Println("one")
	// } else {
	// 	fmt.Println("I dont know")
	// }

	// if b := 100; b == 100 {
	// 	fmt.Println("one hundred")
	// }

	// x := 0
	// if x:= 2; true {
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// エラーハンドリング
	// var s string = "A"
	
	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("i = %T\n", i)

	// for
	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1,2,3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1,2,3}

	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// switch

	// n := 1
	// switch n {
	// case 1,2:
	// 	fmt.Println("1 or 2")
	// case 3,4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("end")
	// }

	// switch n:=2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I dont know")
	// }

	// n:=6
	// switch {
	// case n > 0 && n < 4:
	// 	fmt.Println("0 < n < 4")
	// case n > 3 && n < 7:
	// 	fmt.Println("3 < n < 7")
	// }

	// switch 型スイッチ

	// anything("aaa")
	// anything(1)

	// var x interface{} = 3
	// i, isInt := x.(int)
	// fmt.Println(i + 2, isInt)

	// f, isFloat64 := x.(float64)
	// fmt.Println(f, isFloat64)

	// if x == nil {
	// 	fmt.Println("None")
	// } else if i, isInt := x.(int); isInt {
	// 	fmt.Println(i,"x is Int")
	// } else if s, isString := x.(string); isString {
	// 	fmt.Println(s, isString)
	// } else {
	// 	fmt.Println("I dont know")
	// }

	// switch x.(type) {
	// case int:
	// 	fmt.Println("int")
	// case string:
	// 	fmt.Println("string")
	// default:
	// 	fmt.Println("I dont know")
	// }

	// ラベル付きfor
// 	Loop:
// 	for {
// 		for {
// 			for {
// 				fmt.Println("START")
// 				break Loop
// 			}
// 			fmt.Println("処理をしない")
// 		}
// 		fmt.Println("処理をしない")
// 	}
// 	fmt.Println("END")
// }

	// for i:=0; i < 3; i++ {
	// 	for j := 1; j < 3; j++ {
	// 		fmt.Println(i, j, i*j)
	// 	}
	// }

	// defer
	// TestDefer()

	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	// RunDefer()

	// file, err := os.Create("test.txt") 
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// file.Write([]byte("Hello"))

	// panic recover
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }()
	// panic("runtime error") 
	// fmt.Println("START")

	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}