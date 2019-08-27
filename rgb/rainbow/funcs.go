package main

// RGB - a set of arrays to hold pre-calculated RGB values
// type RGB struct {
// 	red   []uint8
// 	green []uint8
// 	blue  []uint8
// }

// // SinAppend special function to build my SIN map slices
// func SinAppend(slice []uint8, value float64, maxBrightness float64) []uint8 {
// 	if value >= 0 {
// 		return append(slice, uint8(value*maxBrightness))
// 	}
// 	return append(slice, 0)
// }

// // Color creates a uint32 value of the RGB color
// func Color(red uint8, green uint8, blue uint8) uint32 {
// 	const white uint8 = 0
// 	return (uint32(white) << 24) | (uint32(red) << 16) | (uint32(green) << 8) | uint32(blue)
// }

// // RainbowCosColor takes an angle and creates a color from it.
// // To use, simply map your light strip to 360 degrees of colors.
// func RainbowCosColor(floats RGB, seed int) uint32 {
// 	maxIndex := len(floats.red)
// 	// splitIndex := len(floats.red) / 3
// 	red := floats.red[seed%maxIndex]
// 	green := floats.green[seed%maxIndex]
// 	blue := floats.blue[seed%maxIndex]
// 	return Color(red, green, blue)
// }
