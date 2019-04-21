package main

import (
	"flag"
	"log"

	"trengo/presenter"
	"trengo/requester"
	"trengo/requester/condition"
)

var (
	lang string
	w    bool
	m    bool
	g    bool
	n    int
)

func main() {
	flag.StringVar(&lang, "l", "", "language name")
	flag.BoolVar(&w, "w", false, "weekly search")
	flag.BoolVar(&m, "m", false, "monthly search")
	flag.IntVar(&n, "n", 10, "num print")

	flag.Parse()
	if w && m {
		log.Fatalf("too many parameters.")
	}

	rangeType := condition.Daily
	if w {
		rangeType = condition.Weekly
	}
	if m {
		rangeType = condition.Monthly
	}

	response := requester.Request(rangeType, lang)
	presenter.Print(response, rangeType, n)
}
