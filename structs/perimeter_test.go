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
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}

	// }

	// t.Run("area of rectangle", func(t *testing.T) {
	// 	rectangle := RecTangle{10.0, 8.0}
	// 	// got := rectangle.Area()
	// 	// want := 80.0
	// 	// if got != want {
	// 	// 	t.Errorf("got %v want %v", got, want)
	// 	// }
	// 	checkArea(t, rectangle, 80.0)
	// })
	// t.Run("area of circle", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	// got := circle.Area()
	// 	// want := 314.1592653589793
	// 	// if got != want {
	// 	// 	t.Errorf("got %v want %v ", got, want)
	// 	// }
	// 	checkArea(t, circle, 314.1592653589793)
	// })

	// 创建一个匿名的结构体数组
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "RecTangle", shape: RecTangle{10.0, 8.0}, want: 80.0},
		{name: "Triangle", shape: Triangle{10.0, 6.0}, want: 30.1},
	}
	for _, v := range areaTest {
		t.Run(v.name, func(t *testing.T) {
			got := v.shape.Area()
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
