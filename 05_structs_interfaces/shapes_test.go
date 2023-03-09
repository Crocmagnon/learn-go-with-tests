package _5_structs_interfaces

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangles", shape: Rectangle{12, 6}, want: 72},
		{name: "circles", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangles", shape: Triangle{12, 6}, want: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangles", shape: Rectangle{10, 15}, want: 50},
		{name: "circles", shape: Circle{10}, want: 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()

			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}
