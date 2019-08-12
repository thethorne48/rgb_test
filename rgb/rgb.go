package main

import (
	"fmt"
	"os"
	"time"

	ws "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	pin        = 18
	count      = 150
	brightness = 100
)

// main test program
func main() {
	opt := ws.DefaultOptions
	opt.Channels[0].LedCount = count
	opt.Channels[0].Brightness = brightness

	led, err := ws.MakeWS2811(&opt)
	if err != nil {
		// desc := ws.StatusDesc(err.)
		panic(err)
	}
	err = led.Init()
	defer led.Fini()
	if err != nil {
		// desc := ws.StatusDesc(err.)
		panic(err)
	}

	fmt.Println("Press Ctr-C to quit.")
	if err != nil {
		fmt.Println(err)
	} else {
		for ok := true; ok; {
			fmt.Println("Creating blue color wipe")
			err = colorWipe(led, uint32(0x000020))
			if err != nil {
				fmt.Println("Error during wipe " + err.Error())
				os.Exit(-1)
			}

			fmt.Println("Creating red color wipe")
			err = colorWipe(led, uint32(0x002000))
			if err != nil {
				fmt.Println("Error during wipe " + err.Error())
				os.Exit(-1)
			}

			fmt.Println("Creating green color wipe")
			err = colorWipe(led, uint32(0x200000))
			if err != nil {
				fmt.Println("Error during wipe " + err.Error())
				os.Exit(-1)
			}
		}
	}
}

func colorWipe(instance *ws.WS2811, color uint32) error {
	colors := instance.Leds(0)
	for i := 0; i < count; i++ {
		colors[i] = color
		err := instance.SetLedsSync(0, colors)
		err = instance.Render()
		if err != nil {
			instance.Wait()
			return err
		}

		time.Sleep(5 * time.Millisecond)
	}

	return nil
}
