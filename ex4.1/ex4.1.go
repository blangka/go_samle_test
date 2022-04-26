//ch4/ex4.1/ex4.1.go
package main

import "fmt"

func main() {
	var a int = 10 // 선언
	var msg string = "Hello Variable"

	a = 20               // 변경
	msg = "Good Morning" // msg 변수 값 변경 - ❹
	fmt.Println(msg, a)  // msg 와 a 값 출력 - ❺
}
