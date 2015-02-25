package pythagorean

import "math"

func Hypotenuse(side1, side2 float64) float64 {
	hypotenuse := math.Sqrt(math.Pow(side1, 2) + math.Pow(side2, 2))
	return hypotenuse
}

func ShorterSide(hypotenuse, side1 float64) float64 {
	side2 := math.Sqrt(math.Pow(hypotenuse, 2) - math.Pow(side1, 2))
	return side2
}
