package main

import (
	"fmt"
)

func main() {
	j(3, 15000000)
}

func j(m, n int) {
	f := 0
	for i := 1; i <= n; i++ {
		f = (f + m) % i
	}
	fmt.Println(f + 1)
}
