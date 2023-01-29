package main

import "fmt"

func hello() string {
	return "hello"
}
func clHello() func(string) string {
	v1 := "Hello"
	var hello func(string) string
	hello = func(world string) string {
		v1 = v1 + world
		return v1
	}

	return hello
}
func CLHello() func(string) string {
	hello := "Hello"
	return func(world string) string {
		return hello + world
	}
}
func main() {
	// result1 := hello()
	// result1()
	// result := CLHello() //result variable is holding a function now
	r1, r2 := CLHello(), CLHello()
	msg1 := r1(" World")
	msg2 := r2(" Golang")
	fmt.Println(msg1, msg2)
	var j interface{}

	j = "Hello"
	j = 1
	j = 1.1
	j = true
	fmt.Println(j)
}
