package main

import "os"
import "strings"
import "fmt"

func main() {
	os.Setenv("userLoggedIn", "false")
	fmt.Println("userLoggedIn:", os.Getenv("userLoggedIn"))
	fmt.Println("user:", os.Getenv("user"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
