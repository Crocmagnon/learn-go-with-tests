package _5_structs_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 15}
	got := Perimeter(rectangle)
	want := 50.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangles", Rectangle{12, 6}, 72},
		{"circles", Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}
