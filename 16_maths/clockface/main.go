package main

import (
	"fmt"
	clockface "learn-go-with-tests/16_maths"
	"os"
	"time"
)

func main() {
	if err := clockface.WriteSVG(os.Stdout, time.Now()); err != nil {
		msg := "There has been an error generating the clock:"
		if _, perr := fmt.Fprintln(os.Stderr, msg); perr != nil {
			panic(perr)
		}
		if _, perr := fmt.Fprintln(os.Stderr, err); perr != nil {
			panic(perr)
		}
	}
}
