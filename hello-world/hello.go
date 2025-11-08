package main

import "fmt"

const HelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return HelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
