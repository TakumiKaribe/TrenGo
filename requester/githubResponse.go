package requester

import (
	"fmt"
)

type GitHubResponse struct {
	lang         string
	rangeType    RangeType
	developers   [50]string
	names        [50]string
	URLs         [50]string
	descriptions [50]string
	languages    [50]string
	stars        [50]int
	forks        [50]int
	builtBy      [50][]struct {
		name string
		url  string
	}
	rangeStar [50]int
	length    int
}

func (r *GitHubResponse) CLIPrint(n int) {
	fmt.Printf("========== GitHub Trending ==========\n\n")
	count := func() int {
		if n < r.length {
			return n
		} else {
			return r.length
		}
	}
	for i := 0; i < count(); i++ {
		fmt.Println("★-----★-----★-----★-----★-----★-----★-----★-----★-----★-----★")
		fmt.Printf("  [rank] %d\n", i+1)
		fmt.Printf("  [developer] %s\n", r.developers[i])
		fmt.Printf("  [name] %s\n", r.names[i])
		fmt.Printf("  [url] %s\n", githubURL+r.URLs[i])
		fmt.Printf("  [description] %s\n", r.descriptions[i])
		fmt.Printf("  [language] %s\n", r.languages[i])
		fmt.Printf("  [sumStars] %d\n", r.stars[i])
		fmt.Printf("  [forks] %d\n", r.forks[i])
		fmt.Printf("  [builtBy]\n")
		for _, b := range r.builtBy[i] {
			fmt.Printf("    [name] %s\n", b.name)
			fmt.Printf("    [url] %s\n", b.url)
		}
		fmt.Printf("  [trend] %d stars %s\n", r.rangeStar[i], r.rangeType.QueryString())
		fmt.Println()
	}
}
