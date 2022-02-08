package main

import "fmt"

func main() {
	h := HelloWorld()
	fmt.Println(h)
}

func HelloWorld() string {
	return "Hello, World!"
}
