package main

import "testing"

type AreaTest struct {
	Item Shape
	Want float64
}

func assert(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func assertArea(t testing.TB, at AreaTest) {
	got := at.Item.Area()
	assert(t, got, at.Want)
}

func TestPerimeter(t *testing.T) {
	rect := Rect{10.0, 10.0}
	assert(t, rect.Perimeter(), 40.0)
}

func TestArea(t *testing.T) {

	areaTests := []AreaTest{
		{Rect{10.0, 10.0}, 100.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, at := range areaTests {
		assertArea(t, at)
	}

	// t.Run("rect", func(t *testing.T) {
	// 	rect := Rect{10.0, 10.0}
	// 	assertArea(t, rect, 100.0)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	assertArea(t, circle, 314.1592653589793)
	// })
}
