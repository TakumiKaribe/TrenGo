package main

import (
	"flag"

	"github.com/maka-nai/trengo/GitHubParser"
	"github.com/maka-nai/trengo/RangeType"
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
	flag.BoolVar(&d, "d", true, "daily search")
	flag.BoolVar(&w, "w", false, "weekly search")
	flag.BoolVar(&m, "m", false, "monthly search")
	flag.BoolVar(&g, "g", true, "GitHub search")
	flag.BoolVar(&j, "j", false, "json format")

	flag.Parse()

	var githubResponses []GitHubParser.GitHubResponse
	if g {
		if d {
			githubResponses = append(githubResponses, GitHubParser.ParseGitHub(RangeType.Daily, lang))
		}
		if w {
			githubResponses = append(githubResponses, GitHubParser.ParseGitHub(RangeType.Weekly, lang))
		}
		if m {
			githubResponses = append(githubResponses, GitHubParser.ParseGitHub(RangeType.Monthly, lang))
		}

		for _, r := range githubResponses {
			GitHubParser.PrintGitHub(r, j)
		}
	}
}
