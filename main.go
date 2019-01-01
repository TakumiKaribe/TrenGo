package main

import (
	"flag"
)

func main() {
	var (
		lang string
		d    bool
		w    bool
		m    bool
		j    bool
		g    bool
	)
	flag.StringVar(&lang, "l", "", "language name")
	flag.BoolVar(&d, "d", false, "daily search")
	flag.BoolVar(&w, "w", false, "weekly search")
	flag.BoolVar(&m, "m", false, "monthly search")
	flag.BoolVar(&g, "g", true, "GitHub search")
	flag.BoolVar(&j, "j", false, "json format")

	flag.Parse()

	if lang != "" {
		lang = "/" + lang
	}

	rangeType = parseRangeType(d, w, m)

	if g {
		parseGitHub()
		printGitHub(j)
	}
}
