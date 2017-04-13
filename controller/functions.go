package controller

import (
	"math"
	"strconv"
)

func Coord(x, y float64) (float64, float64) {
	x += 90
	y += 90

	y = clamp(90-math.Sqrt((180*x)-math.Pow(x, 2)), 90+math.Sqrt((180*x)-math.Pow(x, 2)), y)
	x = clamp(90-math.Sqrt((180*y)-math.Pow(y, 2)), 90+math.Sqrt((180*y)-math.Pow(y, 2)), x)

	return x, y
}

func clamp(min, max, number float64) float64 {

	return math.Min(max, math.Max(min, number))

}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
