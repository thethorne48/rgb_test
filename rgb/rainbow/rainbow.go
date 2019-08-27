package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

const (
	pin           = 18
	count         = 150
	brightness    = 100
	maxAngle      = 360
	maxBrightness = float64(180)
	piDivision    = float64(2)
)

// RGB - a set of arrays to hold pre-calculated RGB values
type RGB struct {
	red   []uint8
	green []uint8
	blue  []uint8
}

var (
	floats = RGB{
		red:   []uint8{},
		green: []uint8{},
		blue:  []uint8{},
	}
)

func main() {
	defer ws2811.Fini()
	colorValues := initRange(count)
	err := ws2811.Init(pin, count, brightness)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Press Ctr-C to quit.")

		fmt.Println("Creating cosine rainbow")
		for index := 0; index <= 1000; index++ {
			err = rainbowCosCycle(colorValues, index)
			if err != nil {
				fmt.Println("Error during cycle " + err.Error())
				os.Exit(-1)
			}
		}

		fmt.Println("Creating color flash")
		colors := [...]uint32{
			0xFF0000, // green
			0x888800, // yellow
			0x00FF00, // red
			0x00FFFF, // purple
			0x0000FF, // blue
			0xFF00FF, // cyan
			0x000000, // blank
		}
		for i := 0; i < len(colors); i++ {
			err = colorFlash(colors[i])
			if err != nil {
				fmt.Println("Error during flash " + err.Error())
				os.Exit(-1)
			}
		}

	}
}

func initRange(ledCount int) RGB {
	// floats := RGB{
	// 	red:   []uint8{},
	// 	green: []uint8{},
	// 	blue:  []uint8{},
	// }

	segmentSize := (math.Pi * 3) / float64(ledCount)

	for i := math.Pi * -1; i <= math.Pi*2; i += segmentSize {
		red := math.Sin(i/piDivision + math.Pi/2)
		blue := math.Sin(i / piDivision)
		floats.red = SinAppend(floats.red, red, maxBrightness)
		floats.blue = SinAppend(floats.blue, blue, maxBrightness)
		if i <= 0 {
			green := math.Sin(i/piDivision + math.Pi)
			floats.green = SinAppend(floats.green, green, maxBrightness)
		} else {
			green := math.Sin(i/piDivision - math.Pi/2)
			floats.green = SinAppend(floats.green, green, maxBrightness)
		}
	}

	return floats
}

func rainbowCosCycle(floats RGB, seed int) error {
	for i := 0; i < count; i++ {
		fmt.Printf("i: %d - seed: %d - color: %X\n", i, seed, RainbowCosColor(floats, i+seed))
		hexColor := RainbowCosColor(floats, i+seed)
		white := float64((hexColor&0xFF000000)>>24) / maxBrightness
		red := float64((hexColor&0x00FF0000)>>16) / maxBrightness
		green := float64((hexColor&0x0000FF00)>>8) / maxBrightness
		blue := float64(hexColor&0x000000FF) / maxBrightness
		fmt.Printf("white: %f - red: %f - green: %f - blue: %f\n", white, red, green, blue)

		ws2811.SetLed(i, RainbowCosColor(floats, i+seed))
	}
	err := ws2811.Render()
	if err != nil {
		ws2811.Clear()
		return err
	}
	time.Sleep(5 * time.Millisecond)
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

func colorFlash(color uint32) error {
	for i := 0; i < count; i++ {
		ws2811.SetLed(i, color)
	}
	err := ws2811.Render()
	if err != nil {
		ws2811.Clear()
		return err
	}

	time.Sleep(5 * time.Second)

	return nil
}

// func crossFade(newColor uint32) (err error) {
// 	ws2811.
// 	return nil
// }

// SinAppend special function to build my SIN map slices
func SinAppend(slice []uint8, value float64, maxBrightness float64) []uint8 {
	if value >= 0 {
		return append(slice, uint8(value*maxBrightness))
	}
	return append(slice, 0)
}

// Color creates a uint32 value of the RGB color
func Color(red uint8, green uint8, blue uint8) uint32 {
	const white uint8 = 0
	return (uint32(white) << 24) | (uint32(red) << 16) | (uint32(green) << 8) | uint32(blue)
}

// RainbowCosColor takes an angle and creates a color from it.
// To use, simply map your light strip to 360 degrees of colors.
func RainbowCosColor(floats RGB, seed int) uint32 {
	maxIndex := len(floats.red)
	// splitIndex := len(floats.red) / 3
	red := floats.red[seed%maxIndex]
	green := floats.green[seed%maxIndex]
	blue := floats.blue[seed%maxIndex]
	return Color(red, green, blue)
}
