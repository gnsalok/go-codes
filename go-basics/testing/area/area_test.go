package area

import "testing"

// func TestArea(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}
// 	got := Area(rectangle)
// 	want := 100.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestArea(t *testing.T) {
// 	t.Run("circle", func(t *testing.T) {
// 		circle

// 	})
// }

// Instead

// func TestArea(t *testing.T) {

// 	// Validate Shape type (interface)
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	// attaching rectangle type to interface, it will call the Area function for rectangle
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	// attaching circle type to interface , it will call the Area function for circle
// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})

// }

// -- Refactoring

func TestArea(t *testing.T) {

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
