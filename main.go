package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Reverse("I am reversed"))
}

func Reverse(s string) string {
	var b strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}
