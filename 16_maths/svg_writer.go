package clockface

import (
	"bufio"
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
	b := bufio.NewWriter(w)
	_, _ = b.WriteString(svgStart)
	_, _ = b.WriteString(bezel)
	_, _ = b.WriteString(makeHourHand(t))
	_, _ = b.WriteString(makeMinuteHand(t))
	_, _ = b.WriteString(makeSecondHand(t))
	_, _ = b.WriteString(svgEnd)
	if err := b.Flush(); err != nil {
		return err
	}
	return nil
}

func makeSecondHand(t time.Time) string {
	p := makeHandPoint(secondHandPoint(t), secondHandLength)
	s := fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
	return s
}

func makeMinuteHand(t time.Time) string {
	p := makeHandPoint(minuteHandPoint(t), minuteHandLength)
	s := fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#888;stroke-width:3px;"/>`, p.X, p.Y)
	return s
}

func makeHourHand(t time.Time) string {
	p := makeHandPoint(hourHandPoint(t), hourHandLength)
	s := fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:5px;"/>`, p.X, p.Y)
	return s
}

func makeHandPoint(p Point, length float64) Point {
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
