package requester

import "fmt"

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
	rangeStar    [50]int
	length       int
}

func (r *GitHubResponse) JSONPrint() {
	fmt.Println("[")
	for i := 0; i < r.length; i++ {
		fmt.Println("{")
		fmt.Printf("  developer: \"%s\",\n", r.developers[i])
		fmt.Printf("  name: \"%s\",\n", r.names[i])
		fmt.Printf("  url: \"%s\",\n", githubURL+r.URLs[i])
		fmt.Printf("  description: \"%s\",\n", r.descriptions[i])
		fmt.Printf("  language: \"%s\",\n", r.languages[i])
		fmt.Printf("  sumStars: %d,\n", r.stars[i])
		fmt.Printf("  forks: %d,\n", r.forks[i])
		fmt.Printf("  ranged: {\n")
		fmt.Printf("    type: \"%s\",\n", r.rangeType.QueryString())
		fmt.Printf("    stars: %d\n", r.rangeStar[i])
		fmt.Printf("  }\n")
		fmt.Print("}")
		if i < r.length-1 {
			fmt.Print(",\n")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("]")
}

func (r *GitHubResponse) CLIPrint() {
	fmt.Printf("========== GitHub Trending ==========\n\n")
	for i := 0; i < r.length; i++ {
		fmt.Println("★-----★-----★-----★-----★-----★-----★-----★-----★-----★-----★")
		fmt.Printf("  [rank] %d\n", i+1)
		fmt.Printf("  [developer] %s\n", r.developers[i])
		fmt.Printf("  [name] %s\n", r.names[i])
		fmt.Printf("  [url] %s\n", githubURL+r.URLs[i])
		fmt.Printf("  [description] %s\n", r.descriptions[i])
		fmt.Printf("  [language] %s\n", r.languages[i])
		fmt.Printf("  [sumStars] %d\n", r.stars[i])
		fmt.Printf("  [forks] %d\n", r.forks[i])
		fmt.Printf("  [trend] %d stars %s\n", r.rangeStar[i], r.rangeType.QueryString())
		fmt.Println()
	}
}
