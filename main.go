package main

import (
	"fmt"
	v1 "github.com/tinyweet/go-lib-toy"
	v1sub "github.com/tinyweet/go-lib-toy/subdir"
	v2 "github.com/tinyweet/go-lib-toy/v2"
	v2sub "github.com/tinyweet/go-lib-toy/v2/subdir"
)

func main() {
	print(v1.Print, "v1")
	print(v1sub.Print, "v1 sub dir")
	print(v2.Print, "v2")
	print(v2sub.Print, "v2 sub dir")
}

func print(handler func(string), str string) {
	handler(str)
	fmt.Println()
}
