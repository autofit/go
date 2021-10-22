package main

import (
	"fmt"
	"github.com/autofit/go/autofit"
)

func main() {
	for i := 0; i < 1000000; i++ {
		fmt.Println(autofit.GetId())
	}
}
