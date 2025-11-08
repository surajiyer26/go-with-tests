package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 3.5, Width: 7.2}, want: 21.4},
		{name: "Circle", shape: Circle{Radius: 10}, want: 62.83185307179586},
		{name: "Square", shape: Square{Side: 5.2}, want: 20.8},
	}

	for _, tt := range perimeterTests {
		t.Run("Perimeter for "+tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 3.5, Width: 7.2}, want: 25.2},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Square", shape: Square{Side: 5.2}, want: 27.040000000000003},
	}

	for _, tt := range areaTests {
		t.Run("Area for "+tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
