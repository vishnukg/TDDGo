// Package main is entry point
package main

import "fmt"

// Hello returns Hello name
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Vish"))
}
