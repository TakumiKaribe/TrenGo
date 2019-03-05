package requester

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const githubURL string = "https://github.com"
const trending string = "trending"

func ParseGitHub(rt RangeType, lang string) GitHubResponse {
	url := makeURL([]string{githubURL, trending, lang}, []string{"since=" + rt.QueryString()})

	// Request the HTML page.
	res, err := http.Get(url)
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

	return parse(doc, rt)
}

func parse(doc *goquery.Document, rt RangeType) GitHubResponse {
	var response = GitHubResponse{}
	response.rangeType = rt

	trends := doc.Find("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.clearfix > div > div.col-md-9.float-md-left > div.explore-content > ol")

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
			} else {
				ns.Children().Each(func(k int, nss *goquery.Selection) {
					builtByName := strings.TrimSpace(nss.AttrOr("href", "default"))
					builtBy := BuiltBy{name: builtByName[1:], url: githubURL + builtByName}
					response.builtBy[i][k] = builtBy
				})
			}
		})
	})

	// parse today stars
	divf6.Each(func(i int, s *goquery.Selection) {
		ns := s.Find("span.d-inline-block.float-sm-right")
		text := strings.TrimSpace(ns.Text())
		if text == "" {
			response.rangeStar[i] = 0
		} else {
			splited := strings.Split(text, " ")
			star, err := strconv.Atoi(strings.Replace(splited[0], ",", "", -1))
			if err == nil {
				response.rangeStar[i] = star
			}
		}
	})

	response.length = divf6.Length()

	return response
}

func makeURL(path []string, parameter []string) string {
	pathString := strings.Join(filter(path, func(str string) bool { return str != "" }), "/")
	parameterString := strings.Join(filter(parameter, func(str string) bool { return str != "" }), "&")
	return strings.Join([]string{pathString, parameterString}, "?")
}

func filter(target []string, isIncluded func(str string) bool) []string {
	var ret []string
	for _, e := range target {
		if isIncluded(e) {
			ret = append(ret, e)
		}
	}

	return ret
}
