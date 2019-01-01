package GitHubParser

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/maka-nai/trengo/Definition"
	"github.com/maka-nai/trengo/RangeType"
)

func ParseGitHub(rt RangeType.RangeType, lang string) GitHubResponse {
	if lang != "" {
		lang = "/" + lang
	}
	// Request the HTML page.
	res, err := http.Get(Definition.GithubURL + Definition.Trending + lang + "?since=" + rt.QueryString())
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response = GitHubResponse{}
	response.rangeType = rt

	trends := doc.Find("body > div.application-main > div.explore-pjax-container.container-lg.p-responsive.clearfix > div > div.col-md-9.float-md-left > div.explore-content > ol")

	trends.Each(func(i int, s *goquery.Selection) {
		s.Find("div.d-inline-block.col-9.mb-1 > h3 > a").Each(func(i int, s *goquery.Selection) {
			text := strings.Split(strings.Replace(s.Text(), " ", "", -1), "/")
			response.developers[i] = strings.TrimSpace(text[0])
			response.names[i] = strings.TrimSpace(text[1])
			response.URLs[i] = strings.TrimSpace(s.AttrOr("href", "default"))
		})
		s.Find("div.py-1").Each(func(i int, s *goquery.Selection) {
			response.descriptions[i] = strings.TrimSpace(s.Text())
		})
	})

	divf6 := trends.Find("div.f6.text-gray.mt-2")
	divf6.Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(j int, ns *goquery.Selection) {
			text := strings.TrimSpace(ns.Text())
			attr := strings.TrimSpace(ns.Children().AttrOr("class", "default"))
			if attr == "repo-language-color ml-0" && text != "" {
				response.languages[i] = text
			} else if attr == "octicon octicon-star" && text != "" {
				if strings.TrimSpace(ns.Children().AttrOr("aria-label", "default")) == "star" {
					v, err := strconv.Atoi(strings.Replace(text, ",", "", -1))
					if err == nil {
						response.stars[i] = v
					}
				}
			} else if attr == "octicon octicon-repo-forked" && text != "" {
				v, err := strconv.Atoi(strings.Replace(text, ",", "", -1))
				if err == nil {
					response.forks[i] = v
				}
			}
		})
	})

	divf6.Each(func(i int, s *goquery.Selection) {
		ns := s.Find("span.d-inline-block.float-sm-right")
		text := strings.TrimSpace(ns.Text())
		if text == "" {
			response.rangeStar[i] = 0
		} else {
			splited := strings.Split(text, " ")
			star, err := strconv.Atoi(splited[0])
			if err == nil {
				response.rangeStar[i] = star
			}
		}
	})

	response.length = divf6.Length()

	return response
}

func PrintGitHub(response GitHubResponse, isJSONFormat bool) {
	if isJSONFormat {
		fmt.Println("[")
		for i := 0; i < response.length; i++ {
			fmt.Println("{")
			fmt.Printf("  developer: %s,\n", "\""+response.developers[i]+"\"")
			fmt.Printf("  name: %s,\n", "\""+response.names[i]+"\"")
			fmt.Printf("  url: %s,\n", "\""+Definition.GithubURL+response.URLs[i]+"\"")
			fmt.Printf("  description: %s,\n", "\""+response.descriptions[i]+"\"")
			fmt.Printf("  language: %s,\n", "\""+response.languages[i]+"\"")
			fmt.Printf("  sumStars: %d,\n", response.stars[i])
			fmt.Printf("  forks: %d,\n", response.forks[i])
			fmt.Printf("  ranged: {\n")
			fmt.Printf("    type: %s,\n", "\""+response.rangeType.QueryString()+"\"")
			fmt.Printf("    stars: %d\n", response.rangeStar[i])
			fmt.Printf("  }\n")
			fmt.Print("}")
			if i < response.length-1 {
				fmt.Print(",\n")
			} else {
				fmt.Println()
			}
		}
		fmt.Println("]")

	} else {
		fmt.Printf("========== GitHub Trending ==========\n\n")
		for i := 0; i < response.length; i++ {
			fmt.Println("★-----★-----★-----★-----★-----★-----★-----★-----★-----★-----★")
			fmt.Printf("  [rank] %d\n", i+1)
			fmt.Printf("  [developer] %s\n", response.developers[i])
			fmt.Printf("  [name] %s\n", response.names[i])
			fmt.Printf("  [url] %s\n", Definition.GithubURL+response.URLs[i])
			fmt.Printf("  [description] %s\n", response.descriptions[i])
			fmt.Printf("  [language] %s\n", response.languages[i])
			fmt.Printf("  [sumStars] %d\n", response.stars[i])
			fmt.Printf("  [forks] %d\n", response.forks[i])
			// TODO: -d -m などで複数のrangeを投げると、最初のリクエストしか取得できていない
			fmt.Printf("  [trend] %d stars %s\n", response.rangeStar[i], response.rangeType.QueryString())
			fmt.Println()
		}
	}
}
