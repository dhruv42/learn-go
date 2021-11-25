package main

import "fmt"

const englistHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englistHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
