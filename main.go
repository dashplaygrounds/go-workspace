package main

import (
	"fmt"
	// "reflect"
)

func main() {
	// == SWITCH STATEMENT ==

	// for i:=0;i<5; i++ {
	// 	fmt.Println(i)
	// }

	// var i, j = 10, 50

	// switch {
	// case i+j == 60:
	// 		fmt.Println("equal to 60")
	// case i+j <= 60:
	// 		fmt.Println("less than or equal to 60")
	// 		fallthrough
	// default:
	// 		fmt.Println("greater than 60")
	// }

	// day := "wednesday"
	// switch day {
	// case "monday":
	// 		fmt.Println("monday")
	// case "tuesday":
	// 		fmt.Println("tuesday")
	// case "wednesday":
	// 		fmt.Println("wednesday")
	// 		fallthrough
	// case "thursday":
	// 		fmt.Println("thursday")
	// 		fallthrough
	// case "friday":
	// 		fmt.Println("friday")
	// case "saturday", "sunday":
	// 		fmt.Println("weekend")
	// default:
	// 		fmt.Println("default")
	// }

	// day := "sunday"
	// switch day {
	// case "monday":
	// 		fmt.Println("monday")
	// case "tuesday":
	// 		fmt.Println("tuesday")
	// case "wednesday":
	// 		fmt.Println("wednesday")
	// case "thursday":
	// 		fmt.Println("thursday")
	// case "friday":
	// 		fmt.Println("friday")
	// case "saturday", "sunday":
	// 		fmt.Println("weekend")
	// default:
	// 		fmt.Println("default")
	// }

	// var a, b = 100, 5
	// switch {
	// case a/b == 10:
	// 		fmt.Println("10")
	// case a/b == 20:
	// 		fmt.Println("20")
	// case a/b == 10:
	// 		fmt.Println("30")
	// default:
	// 		fmt.Println("default")
	// }
	
	

	// var i int = 800
	// switch i {
	// case 10:
	// 	fmt.Println("i is 10")
	// case 100, 200:
	// 	fmt.Println("i is 100 or 200")
	// default:
	// 	fmt.Println("i is neither 10, 100, or 200")
	// }
	// == CONTROL FLOW ==

	// == BITWISE OPERATORS ==

	// var x, y int = 100,90
	// fmt.Println(!(((x+y) >> 2 ) == 47))
	// = false

	// var x, y int = 100,90
	// fmt.Println((x+y) >> 2)
	// = 47

	// var x, y int = 100,90
	// fmt.Println(x & y)
	// fmt.Println(x | y)
	// >> == RIGHT SHIFT (>>) ==	
	// all bits are shifted to position to the right and
	// the first vacated position is filled with 0
	// >> == LEFT SHIFT (<<) ==
	// all bits are shifted to n position to the left and
	// the last vacated position is filled with 0
	// >> == BITWISE XOR (^) ==
	// the result if 1 if two bits are opposite

	// >> == BITWISE OR (|) ==
	// convert into byte and perform operation on every bit

	// >> == BITWISE AND (&) ==
	// var x, y int = 12, 25
	// z := x & y
	// fmt.Println(z)
	// = 8
	// Because they allow greater precision and require fewer resources, bitwise operators can make some code faster and more efficient. Examples of uses of bitwise operations include encryption, compression, graphics, communications over ports/sockets, embedded systems programming and finite state machines.

	// == OTHERS ==
	// var x, y string = "foo", "bar"
	// x += y
	// fmt.Println(x)

	// var a, b bool = true, false
	// fmt.Println(a && b)
	// fmt.Println(a || b)

	// var a, b bool = false, false
	// fmt.Println(a && b)
	// fmt.Println(a || b)

	// var a, b bool = false, true
	// fmt.Println(!a)
	// fmt.Println(b)

	// var a bool = false
	// result := 10 > 50
	// fmt.Println(!(a && result))

	// var a bool = true
	// result := 10 > 50
	// fmt.Println(!(a || result))

	// == AND ==
	// var x int = 10
	// fmt.Println((x < 100) && (x < 200))
	// fmt.Println((x < 300) && (x < 0))

	// == OR ==
	// var x int = 10
	// fmt.Println((x < 0) && (x < 200))
	// fmt.Println((x < 0) && (x > 200))

	// == NOT ==
	// var x, y int = 10, 20
	// fmt.Println(!(x > y))
	// fmt.Println(!(true))
	// fmt.Println(!(false))

	// a := "Hello"
	// b := 32
	// a = "Hi"
	// fmt.Printf("%v %v", a, b)

	// var (
	// s string = "Hey"
	// t int = 123
	// u bool = false)
	// fmt.Println(s,t,u)

	// title := "Sir"
	// fmt.Println(title)
	// fmt.Println("Type of variable:",reflect.ValueOf(title).Kind())

	// var name string = "Mark"
	// var age int = 27
	// var salary float64 = 50000.0
	// var is_valid bool = true
	// fmt.Printf("Details: %v %v %.2f %v", name, age, salary, is_valid)
}