package main

import (
	"fmt"
	"mocxy/generation"
)

func main() {
	err := generation.Parse()
	if err != nil {
		fmt.Println("Parse Error :  ", err)
		return
	}
}
