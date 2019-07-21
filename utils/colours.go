package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// RGBColor RBG Color Type
type rgbColour struct {
	Red   int
	Green int
	Blue  int
}

// GetHex Converts a decimal number to hex representations
func getHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

func getRandomColourInRgb() rgbColour {
	rand.Seed(time.Now().UnixNano())

	red, green, blue := rand.Intn(255), rand.Intn(255), rand.Intn(255)
	return rgbColour{red, green, blue}
}

func RandomColor() int {
	colour := getRandomColourInRgb()

	hex := fmt.Sprintf("0x%s%s%s", getHex(colour.Red), getHex(colour.Green), getHex(colour.Blue))
	hexInt, _ := strconv.ParseInt(hex, 0, 64)

	return int(hexInt)
}
