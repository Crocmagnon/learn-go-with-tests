package main

import (
	clockface "learn-go-with-tests/16_maths"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
