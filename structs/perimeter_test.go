package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangel := RecTangle{10.0, 10.0}
	got := Perimeter(rectangel)
	want := 40.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("area of rectangle", func(t *testing.T) {
		rectangle := RecTangle{10.0, 8.0}
		got := rectangle.Area()
		want := 80.0
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %v want %v ", got, want)
		}
	})
}
