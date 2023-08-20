package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	r := reverse.String("Hello, OTUS!")

	fmt.Println(r)
}
