package main

import (
	"machine"
	"time"
)

func main() {
	println("Let get blinky!")

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		println("LED off")
		led.Low()
		time.Sleep(time.Millisecond * 1000)

		println("LED on")
		led.High()
		time.Sleep(time.Millisecond * 1000)
	}
}
