package main

import "testing"

// func TestPerimeter(t *testing.T) {
// 	rec := Rectangle{10.0, 10.0}
// 	got := Perimeter(rec)
// 	want := 40.0
// 	if got != want {
// 		t.Errorf("got %g want %g", got, want)
// 	}
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectange", shape: Rectangle{12, 6}, hasArea: 72},
		{name: "Circle", shape: Circle{10}, hasArea: 315.1592653589793},
		{name: "Trianglr", shape: Triangle{12, 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.name, got, tt.hasArea)
			}
		})

	}
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }
	// t.Run("rectangles", func(t *testing.T) {
	// 	rec := Rectangle{12.0, 6.0}
	// 	checkArea(t, rec, 72.0)

	// })

	// t.Run("circles", func(t *testing.T) {
	// 	cir := Circle{10}
	// 	checkArea(t, cir, 314.1592653589793)

	// })

}
