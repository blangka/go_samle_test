package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	basic("basic : ")
	너비지정("너비지정 : ")
	소수정지정("소수점지정 : ")
	입력("입력 : ")
	표준입력("표준입력 : ")
}

func 표준입력(name string) {
	stdin := bufio.NewReader(os.Stdin) // ❷ 표준 입력을 읽는 객체

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil { // 에러 발생 시
		fmt.Println(name, err) // 에러 출력
		stdin.ReadString('\n') // ❸ 표준 입력 스트림 지우기
	} else {
		fmt.Println(name, n, a, b)
	}
	n, err = fmt.Scanln(&a, &b) // ❹ 다시 입력받기
	if err != nil {
		fmt.Println(name, err)
	} else {
		fmt.Println(name, n, a, b)
	}
}

func 입력(name string) {
	var a int // ❶ 값을 저장할 변수
	var b int

	n, err := fmt.Scan(&a, &b) // ❷ 입력 두 개 받기
	if err != nil {            // ❸ 에러가 발생하면 에러 코드 출력
		fmt.Println(name, n, err)
	} else { // ➍ 정상 입력되면 입력값 출력
		fmt.Println(name, n, a, b)
	}
}

func 소수정지정(name string) {
	var a = 324.13455
	var b = 324.13455
	var c = 3.14

	fmt.Printf("%s %08.2f\n", name, a) // ❶ 최소너비:8 소수점이하:2자리 0을 채움.
	fmt.Printf("%s %08.2g\n", name, b) // ❷ 최소너비:8 총숫자: 2자리 0을 채움
	fmt.Printf("%s %8.5g\n", name, b)  // ❸ 최소너비:8 총숫자: 5자리
	fmt.Printf("%s %f\n", name, c)     // ❹ 소수점이하 6자리까지 출력
}

func 너비지정(name string) {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%s %5d, %5d\n", name, a, b)    // ❶ 최소 너비보다 짧은 값 너비 지정
	fmt.Printf("%s %05d, %05d\n", name, a, b)  // ❷ 최소 너비보다 짧은 값 0 채우기
	fmt.Printf("%s %-5d, %-05d\n", name, a, b) // ❸ 최소 너비보다 짧은 값 왼쪽 정렬

	fmt.Printf("%s %5d, %5d\n", name, c, c)    // ➍ 최소 너비보다 긴 값 너비 지정
	fmt.Printf("%s %05d, %05d\n", name, c, c)  // ➎ 최소 너비보다 긴 값 0 채우기
	fmt.Printf("%s %-5d, %-05d\n", name, c, c) // ➏ 최소 너비보다 긴 값 왼쪽 정렬
}

func basic(name string) {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print(name, "a:", a, "b:", b)                  // ❶
	fmt.Println(name, "a:", a, "b:", b, "f:", f)       // ❷
	fmt.Printf("%s a: %d b: %d f:%f\n", name, a, b, f) // ❸
}
