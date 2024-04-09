package main

import (
	"fmt"
	"learnMap/demo"
)

func main() {
	i := demo.DeferFunc()
	i2 := demo.DeferFuncNameBack()
	fmt.Println(i)
	fmt.Println(i2)
}
