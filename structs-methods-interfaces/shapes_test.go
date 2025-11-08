package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{3.5, 7.2}, 21.4},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Height: 3.5, Width: 7.2}, want: 25.2},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Square{Side: 5.2}, want: 27.040000000000003},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
