package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

const (
	pin        = 18
	count      = 150
	brightness = 100
	maxAngle   = 255
)

func main() {
	defer ws2811.Fini()
	err := ws2811.Init(pin, count, brightness)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Press Ctr-C to quit.")
		fmt.Println("Creating rainbow circle")
		index := 0
		for {
			err = rainbowCycle(index % maxAngle)
			if err != nil {
				fmt.Println("Error during wipe " + err.Error())
				os.Exit(-1)
			}
			index++
		}
	}
}

func rainbowCycle(angle int) error {
	for i := 0; i < count; i++ {
		ws2811.SetLed(i, RainbowColor((((angle+i)%maxAngle)*maxAngle)/150))
	}
	err := ws2811.Render()
	if err != nil {
		ws2811.Clear()
		return err
	}
	time.Sleep(50 * time.Millisecond)
	return nil
}

func colorWipe(color uint32) error {
	for i := 0; i < count; i++ {
		ws2811.SetLed(i, color)
		err := ws2811.Render()
		if err != nil {
			ws2811.Clear()
			return err
		}

		time.Sleep(5 * time.Millisecond)
	}

	return nil
}
