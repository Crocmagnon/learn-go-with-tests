package clockface_test

import (
	"bytes"
	"encoding/xml"
	"errors"
	clockface "learn-go-with-tests/16_maths"
	"testing"
	"time"
)

func TestWriteSVG(t *testing.T) {
	t.Run("bad writer", func(t *testing.T) {
		writer := badWriter{}
		err := clockface.WriteSVG(writer, time.Now())
		if err == nil {
			t.Error("didn't get an error, expected one")
		}
	})

	t.Run("second hand", func(t *testing.T) {
		cases := []struct {
			time time.Time
			line Line
		}{
			{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
			{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
			{simpleTime(6, 21, 48), Line{150, 150, 64.405, 122.188}},
		}
		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				b := bytes.Buffer{}
				if err := clockface.WriteSVG(&b, c.time); err != nil {
					t.Fatalf("got an error: %q, didn't expect one", err)
				}

				svg := SVG{}
				if err := xml.Unmarshal(b.Bytes(), &svg); err != nil {
					t.Fatalf("error during xml parsing: %q", err)
				}

				if !containsLine(c.line, svg.Line) {
					t.Errorf("Expected to find the second hand line %+v in the SVG lines %+v", c.line, svg.Line)
				}
			})
		}
	})

	t.Run("minute hand", func(t *testing.T) {
		cases := []struct {
			time time.Time
			line Line
		}{
			{simpleTime(0, 0, 0), Line{150, 150, 150, 80}},
			{simpleTime(6, 21, 48), Line{150, 150, 202.99, 195.739}},
		}
		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				b := bytes.Buffer{}
				if err := clockface.WriteSVG(&b, c.time); err != nil {
					t.Fatalf("got an error: %q, didn't expect one", err)
				}

				svg := SVG{}
				if err := xml.Unmarshal(b.Bytes(), &svg); err != nil {
					t.Fatalf("error during xml parsing: %q", err)
				}

				if !containsLine(c.line, svg.Line) {
					t.Errorf("Expected to find the minute hand line %+v in the SVG lines %+v", c.line, svg.Line)
				}
			})
		}
	})

	t.Run("hour hand", func(t *testing.T) {
		cases := []struct {
			time time.Time
			line Line
		}{
			{simpleTime(6, 0, 0), Line{150, 150, 150, 200}},
			{simpleTime(6, 21, 48), Line{150, 150, 140.545, 199.098}},
		}
		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				b := bytes.Buffer{}
				if err := clockface.WriteSVG(&b, c.time); err != nil {
					t.Fatalf("got an error: %q, didn't expect one", err)
				}

				svg := SVG{}
				if err := xml.Unmarshal(b.Bytes(), &svg); err != nil {
					t.Fatalf("error during xml parsing: %q", err)
				}

				if !containsLine(c.line, svg.Line) {
					t.Errorf("Expected to find the minute hand line %+v in the SVG lines %+v", c.line, svg.Line)
				}
			})
		}
	})
}

type badWriter struct {
}

func (b badWriter) Write(_ []byte) (n int, err error) {
	return 0, errors.New("unexpected error during write")
}

func containsLine(want Line, lines []Line) bool {
	for _, line := range lines {
		if line == want {
			return true
		}
	}
	return false
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 24, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}
