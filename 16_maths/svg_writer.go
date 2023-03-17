package clockface

import (
	"fmt"
	"io"
	"time"
)

const secondHandLength = 90
const minuteHandLength = 70
const hourHandLength = 50
const clockCentreX = 150
const clockCentreY = 150

// SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) error {
	_, err := io.WriteString(w, svgStart)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, bezel)
	if err != nil {
		return err
	}
	err = hourHand(w, t)
	if err != nil {
		return err
	}
	err = minuteHand(w, t)
	if err != nil {
		return err
	}
	err = secondHand(w, t)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, svgEnd)
	if err != nil {
		return err
	}
	return nil
}

func secondHand(w io.Writer, t time.Time) error {
	p := makeHand(secondHandPoint(t), secondHandLength)
	_, err := fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
	if err != nil {
		return err
	}
	return nil
}

func minuteHand(w io.Writer, t time.Time) error {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	_, err := fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#888;stroke-width:3px;"/>`, p.X, p.Y)
	if err != nil {
		return err
	}
	return nil
}

func hourHand(w io.Writer, t time.Time) error {
	p := makeHand(hourHandPoint(t), hourHandLength)
	_, err := fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:5px;"/>`, p.X, p.Y)
	if err != nil {
		return err
	}
	return nil
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
