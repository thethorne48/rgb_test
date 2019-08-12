package main

// LightDefaults are the 360 values precalculated to enable rainbow displays
var LightDefaults = [...]uint8{
	0, 0, 0, 0, 0, 1, 1, 2,
	2, 3, 4, 5, 6, 7, 8, 9,
	11, 12, 13, 15, 17, 18, 20, 22,
	24, 26, 28, 30, 32, 35, 37, 39,
	42, 44, 47, 49, 52, 55, 58, 60,
	63, 66, 69, 72, 75, 78, 81, 85,
	88, 91, 94, 97, 101, 104, 107, 111,
	114, 117, 121, 124, 127, 131, 134, 137,
	141, 144, 147, 150, 154, 157, 160, 163,
	167, 170, 173, 176, 179, 182, 185, 188,
	191, 194, 197, 200, 202, 205, 208, 210,
	213, 215, 217, 220, 222, 224, 226, 229,
	231, 232, 234, 236, 238, 239, 241, 242,
	244, 245, 246, 248, 249, 250, 251, 251,
	252, 253, 253, 254, 254, 255, 255, 255,
	255, 255, 255, 255, 254, 254, 253, 253,
	252, 251, 251, 250, 249, 248, 246, 245,
	244, 242, 241, 239, 238, 236, 234, 232,
	231, 229, 226, 224, 222, 220, 217, 215,
	213, 210, 208, 205, 202, 200, 197, 194,
	191, 188, 185, 182, 179, 176, 173, 170,
	167, 163, 160, 157, 154, 150, 147, 144,
	141, 137, 134, 131, 127, 124, 121, 117,
	114, 111, 107, 104, 101, 97, 94, 91,
	88, 85, 81, 78, 75, 72, 69, 66,
	63, 60, 58, 55, 52, 49, 47, 44,
	42, 39, 37, 35, 32, 30, 28, 26,
	24, 22, 20, 18, 17, 15, 13, 12,
	11, 9, 8, 7, 6, 5, 4, 3,
	2, 2, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0}

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

// RainbowColor takes an angle and creates a color from it.
// To use, simply map your light strip to 360 degrees of colors.
func RainbowColor(angle int) uint32 {
	return Color(LightDefaults[(angle+len(LightDefaults)/3)%len(LightDefaults)], LightDefaults[angle%len(LightDefaults)], LightDefaults[(angle+(len(LightDefaults)/3)*2)%len(LightDefaults)])
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
