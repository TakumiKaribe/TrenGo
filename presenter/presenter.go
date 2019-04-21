package presenter

import (
	"fmt"
	"trengo/constants"
	"trengo/model"
	"trengo/requester/condition"
)

// Print -
func Print(r model.Response, rt condition.RangeType, n int) {
	fmt.Printf("========== GitHub Trending ==========\n\n")
	count := func() int {
		if n < r.Length {
			return n
		} else {
			return r.Length
		}
	}
	for i := 0; i < count(); i++ {
		fmt.Println("★-----★-----★-----★-----★-----★-----★-----★-----★-----★-----★")
		fmt.Printf("  [rank] %d\n", i+1)
		fmt.Printf("  [developer] %s\n", r.Developers[i])
		fmt.Printf("  [name] %s\n", r.Names[i])
		fmt.Printf("  [url] %s\n", constants.GitHubURL+r.URLs[i])
		fmt.Printf("  [description] %s\n", r.Descriptions[i])
		fmt.Printf("  [language] %s\n", r.Languages[i])
		fmt.Printf("  [sumStars] %d\n", r.Stars[i])
		fmt.Printf("  [forks] %d\n", r.Forks[i])
		fmt.Printf("  [builtBy]\n")
		for _, b := range r.BuiltBy[i] {
			fmt.Printf("    [name] %s\n", b.Name)
			fmt.Printf("    [url] %s\n", b.URL)
		}
		fmt.Printf("  [trend] %d stars %s\n", r.RangeStar[i], rt.QueryString())
		fmt.Println()
	}
}
