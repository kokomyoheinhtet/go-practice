package main

import (
	"fmt"
	"strings"
)

func main() {
	namea := "koko"
	nameb := "kOkO"
	fmt.Println(namea)
	fmt.Println(nameb)

	fmt.Println(strings.EqualFold(namea, nameb))
}
