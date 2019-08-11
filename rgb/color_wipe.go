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
)

func main() {
	defer ws2811.Fini()
	err := ws2811.Init(pin, count, brightness)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Press Ctr-C to quit.")
		fmt.Println("Creating blue color wipe")
		err = colorWipe(uint32(0x000020))
		if err != nil {
			fmt.Println("Error during wipe " + err.Error())
			os.Exit(-1)
		}

		fmt.Println("Creating red color wipe")
		err = colorWipe(uint32(0x002000))
		if err != nil {
			fmt.Println("Error during wipe " + err.Error())
			os.Exit(-1)
		}

		fmt.Println("Creating green color wipe")
		err = colorWipe(uint32(0x200000))
		if err != nil {
			fmt.Println("Error during wipe " + err.Error())
			os.Exit(-1)
		}

		fmt.Println("Creating blank color wipe")
		err = colorWipe(uint32(0x000000))
		if err != nil {
			fmt.Println("Error during wipe " + err.Error())
			os.Exit(-1)
		}
	}
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
