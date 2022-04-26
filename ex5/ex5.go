package main

import "fmt"

func main() {

	basic("basic : ")
}

func basic(name string) {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print(name, "a:", a, "b:", b)                  // ❶
	fmt.Println(name, "a:", a, "b:", b, "f:", f)       // ❷
	fmt.Printf("%s a: %d b: %d f:%f\n", name, a, b, f) // ❸
}
