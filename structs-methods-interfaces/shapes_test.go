package main

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{3.5, 7.2}
		checkPerimeter(t, rectangle, 21.4)
	})

	t.Run("perimeter of circle", func(t *testing.T) {
		circle := Circle{10}
		checkPerimeter(t, circle, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{3.5, 7.2}
		checkArea(t, rectangle, 25.2)
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
