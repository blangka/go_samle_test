//ch4/ex4.1/ex4.1.go
package main

import "fmt"

func main() {

	hello("hello : ")
	basic("basic : ")
	초기값("초기값 : ")
	변수변환("변수변환 : ")
	타입변환("타입변환 : ")
}

func 타입변환(name string) {
	var a int16 = 3456
	var c int8 = int8(a) // int16타입에서 int8타입으로 변환

	fmt.Println(name, a)
	fmt.Println(name, c) // int8타입인 c값 출력
}

func 변수변환(name string) {
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // ❶ float64에서 int로 변환
	d := float64(a * c) // int에서 float6로 변환

	var e int64 = 7
	f := int64(d) * e // float64에서 int64로 변환

	var g int = int(b * 3) // float64에서 int로 변환
	var h int = int(b) * 3 // float64에서 int로 변환 g와 값이 다릅니다.
	fmt.Println(name, g, h, f)
}

func 초기값(name string) {
	var a int = 3 // 가장 기본 형태
	var b int     // 초깃값 생략 - 초깃값은 타입별 기본값으로 대체
	var c = 4     // 타입 생략 - 변수의 타입은 우변 값의 타입이 됩니다
	d := 5        // 선언대입문 := 을 사용해서 var 키워드와 타입 생략

	fmt.Println(name, a, b, c, d)
}

func basic(name string) {
	var minimumWage int = 10 // ❶ 변수 minimumWage 선언 및 초기화
	var workingHour int = 20 // ❷ 변수 workingHour 선언 및 초기화

	// ❸ 변수 income 선언 및 초기화
	var income int = minimumWage * workingHour

	// 변수 minimumWage, workingHour, income 출력
	fmt.Println(name, minimumWage, workingHour, income)
}

func hello(name string) {
	var a int = 10 // 선언
	var msg string = "Hello Variable"

	a = 20                    // 변경
	msg = "Good Morning"      // msg 변수 값 변경
	fmt.Println(name, msg, a) // msg 와 a 값 출력
}
