package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(10)
	y := big.NewInt(10)
	fmt.Println(x.Sub(x, y))
}
