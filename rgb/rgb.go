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
	brightness = 255
)

var gamma8 = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5,
	5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10,
	10, 10, 11, 11, 11, 12, 12, 13, 13, 13, 14, 14, 15, 15, 16, 16,
	17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 24, 24, 25,
	25, 26, 27, 27, 28, 29, 29, 30, 31, 32, 32, 33, 34, 35, 35, 36,
	37, 38, 39, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 50,
	51, 52, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 66, 67, 68,
	69, 70, 72, 73, 74, 75, 77, 78, 79, 81, 82, 83, 85, 86, 87, 89,
	90, 92, 93, 95, 96, 98, 99, 101, 102, 104, 105, 107, 109, 110, 112, 114,
	115, 117, 119, 120, 122, 124, 126, 127, 129, 131, 133, 135, 137, 138, 140, 142,
	144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 167, 169, 171, 173, 175,
	177, 180, 182, 184, 186, 189, 191, 193, 196, 198, 200, 203, 205, 208, 210, 213,
	215, 218, 220, 223, 225, 228, 231, 233, 236, 239, 241, 244, 247, 249, 252, 255,
}

var defaultOptions = ws.Option{
	Frequency: 800,
	DmaNum:    ws.DefaultDmaNum,
	Channels: []ws.ChannelOption{
		{
			GpioPin:    23,
			LedCount:   count,
			Brightness: ws.DefaultBrightness,
			StripeType: ws.WS2812Strip,
			Invert:     false,
			Gamma:      gamma8,
		},
	},
}

// main test program
func main() {
	ws2812, err := ws.MakeWS2811(&defaultOptions)
	if err != nil {
		// desc := ws.StatusDesc(err.)
		panic(err)
	}
	defer ws2812.Fini()
	err = ws2812.Init()
	fmt.Println("Press Ctr-C to quit.")
	if err != nil {
		fmt.Println(err)
	} else {
		for ok := true; ok; {
			fmt.Println("Creating blue color wipe")
			err = colorWipe(ws2812, uint32(0x000020))
			if err != nil {
				fmt.Println("Error during wipe " + err.Error())
				os.Exit(-1)
			}

			fmt.Println("Creating red color wipe")
			err = colorWipe(ws2812, uint32(0x002000))
			if err != nil {
				fmt.Println("Error during wipe " + err.Error())
				os.Exit(-1)
			}

			fmt.Println("Creating green color wipe")
			err = colorWipe(ws2812, uint32(0x200000))
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
