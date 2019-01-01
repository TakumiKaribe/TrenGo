package main

import (
	"flag"
	"log"

	"github.com/maka-nai/trengo/requester"
)

var (
	lang string
	d    bool
	w    bool
	m    bool
	j    bool
	g    bool
)

func main() {
	flag.StringVar(&lang, "l", "", "language name")
	flag.BoolVar(&d, "d", false, "daily search")
	flag.BoolVar(&w, "w", false, "weekly search")
	flag.BoolVar(&m, "m", false, "monthly search")
	flag.BoolVar(&g, "g", true, "GitHub search")
	flag.BoolVar(&j, "j", false, "json format")

	flag.Parse()
	checkSearchRange()

	var githubResponse requester.GitHubResponse
	if g {
		var rangeType requester.RangeType
		if d {
			rangeType = requester.Daily
		}
		if w {
			rangeType = requester.Weekly
		}
		if m {
			rangeType = requester.Monthly
		}

		githubResponse = requester.ParseGitHub(rangeType, lang)
		githubResponse.Print(j)
	}
}

func checkSearchRange() {
	rangeTypes := []bool{d, w, m}
	cnt := 0
	for _, e := range rangeTypes {
		if e {
			cnt++
		}
	}

	if cnt > 1 {
		log.Fatalf("too many parameters.")
	}
}
