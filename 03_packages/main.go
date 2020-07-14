package main

import (
	"fmt"
	"math"

	"github.com/rttolipas/go_basics/03_packages/strutil"
)

func main() {
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Floor(2.7))
	fmt.Println(strutil.Reverse("hello"))
}
