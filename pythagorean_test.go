package pythagorean

import "testing"

func TestHypotenuse(t *testing.T) {
	cases := []struct {
		side1, side2, want float64
	}{
		{4, 3, 5},
		{8, 6, 10},
	}
	for _, c := range cases {
		got := Hypotenuse(c.side1, c.side2)
		if got != c.want {
			t.Errorf("Hypotenuse(%f, %f) == %f, want %f", c.side1, c.side2, got, c.want)
		}
	}
}

func TestShorterSide(t *testing.T) {
	cases := []struct {
		hypotenuse, side1, want float64
	}{
		{5, 4, 3},
		{5, 3, 4},
		{10, 6, 8},
		{10, 8, 6},
	}
	for _, c := range cases {
		got := ShorterSide(c.hypotenuse, c.side1)
		if got != c.want {
			t.Errorf("ShorterSide(%f, %f) == %f, want %f", c.hypotenuse, c.side1, got, c.want)
		}
	}
}
