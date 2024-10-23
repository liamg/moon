package main

import (
	"fmt"

	"github.com/liamg/moon"
)

func main() {
	phase := moon.GetPhase()
	fmt.Printf("The moon phase is currently %s - %s\n", phase, phase.Emoji())
}
