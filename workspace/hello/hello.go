package main

import (
	"fmt"

	"other/pkg"

	"golang.org/x/example/hello/reverse"
)

func main() {
	pkg.GetOtherFunc()
	fmt.Println(reverse.String("Hello"))
}
