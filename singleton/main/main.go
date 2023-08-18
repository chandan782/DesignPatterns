package main

import (
	"designpatterns/singleton"
	"fmt"
)

func main() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()

	fmt.Println(s1 == s2) // Output: true
}
