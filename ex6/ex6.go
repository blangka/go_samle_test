package main

import "fmt"

func main() {
	산술연산자()
	비트연산자()
	시프트연산자()
	비교연산자()
}

func 비교연산자() {
	var x int8 = 127 // ❶ 8비트 부호가 있는 정수 최댓값

	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1) // ❷ 비교 연산 수행
	fmt.Printf("x\t= %4d, %08b\n", x, x)           // ❸ x값 확인
	fmt.Printf("x + 1\t= %4d, %08b\n", x+1, x+1)   // ➍ x + 1값 확인
	fmt.Printf("x + 2\t= %4d, %08b\n", x+2, x+2)   // ➎ x + 2값 확인
	fmt.Printf("x + 3\t= %4d, %08b\n", x+3, x+3)   // ➏ x + 3값 확인

	var y int8 = -128                              // 8비트 부호있는 정수 최솟값
	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1) // ➐ 비교 연산 수행
	fmt.Printf("y\t= %4d, %08b\n", y, y)           // ➑ y값 확인
	fmt.Printf("y - 1\t= %4d, %08b\n", y-1, y-1)   // ➒ y - 1값 확인
}

func 시프트연산자() {
	var x int8 = 4  // ❶ 8비트 정수
	var y int8 = 64 // ❷ 8비트 정수

	fmt.Printf("x:%08b x<<2: %08b x<<2: %d\n", x, x<<2, x<<2) // ❸ 왼쪽 시프트
	fmt.Printf("y:%08b y<<2: %08b y<<2: %d\n", y, y<<2, y<<2) // ➍ 왼쪽 시프트
}

func 비트연산자() {
	var x1 int8 = 34   // ❶ 8비트 정수, 00100010
	var x2 int16 = 34  // ❷ 16비트 정수, 00000000 00100010
	var x3 uint8 = 34  // ❸ 8비트 부호가 없는 정수, 00100010
	var x4 uint16 = 34 // ➍ 16비트 부호가 없는 정수,00000000 00100010

	fmt.Printf("^%d = %5d,\t %08b\n", x1, ^x1, uint8(^x1)) // ➎
	fmt.Printf("^%d = %5d,\t %016b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("^%d = %5d,\t %08b\n", x3, ^x3, ^x3)
	fmt.Printf("^%d = %5d,\t %016b\n", x4, ^x4, ^x4)
}

func 산술연산자() {
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t)
}
