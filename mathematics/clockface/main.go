package main

import (
	"os"
	clockface "tdd/mathematics"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
