package main

import (
	"fmt"
	"github.com/barbarum/go-test/src/main/greeting"
	// "github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, go application developer!")
	fmt.Println(greeting.Reverse("!oG"))
	// fmt.Println(cmp.Diff("test", "test1"))
}
