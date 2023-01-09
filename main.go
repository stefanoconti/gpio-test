package main

import (
	"os"
	"strconv"

	"github.com/brian-armstrong/gpio"
)

func main() {
	pinNumber, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		panic(err)
	}
	value, err := strconv.ParseUint(os.Args[2], 10, 32)
	if err != nil {
		panic(err)
	}
	pin := gpio.NewOutput(uint(pinNumber), false)
	if value == 1 {
		pin.High()
	} else {
		pin.Low()
	}
	defer pin.Cleanup()
}
