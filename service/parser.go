package service

import (
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"trengo/constants"
	"trengo/model"
	"trengo/requester/condition"
)

func Parse(body io.ReadCloser, rt condition.RangeType) model.Response {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	var response = model.Response{}

	trends := doc.Find("body > div.application-main > main > div.explore-pjax-container.container-lg.p-responsive.clearfix > div > div.col-md-9.float-md-left > div.explore-content > ol")

	trends.Each(func(i int, s *goquery.Selection) {
		s.Find("div.d-inline-block.col-9.mb-1 > h3 > a").Each(func(i int, s *goquery.Selection) {
			text := strings.Split(strings.Replace(s.Text(), " ", "", -1), "/")
			response.Developers[i] = strings.TrimSpace(text[0])
			response.Names[i] = strings.TrimSpace(text[1])
			response.URLs[i] = strings.TrimSpace(s.AttrOr("href", "default"))
		})
		s.Find("div.py-1").Each(func(i int, s *goquery.Selection) {
			response.Descriptions[i] = strings.TrimSpace(s.Text())
		})
	})

	divf6 := trends.Find("div.f6.text-gray.mt-2")
	divf6.Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(j int, ns *goquery.Selection) {
			text := strings.TrimSpace(ns.Text())
			attr := strings.TrimSpace(ns.Children().AttrOr("class", "default"))
			if attr == "repo-language-color ml-0" && text != "" {
				response.Languages[i] = text
			} else if attr == "octicon octicon-star" && text != "" {
				if strings.TrimSpace(ns.Children().AttrOr("aria-label", "default")) == "star" {
					v, err := strconv.Atoi(strings.Replace(text, ",", "", -1))
					if err == nil {
						response.Stars[i] = v
					}
				}
			} else if attr == "octicon octicon-repo-forked" && text != "" {
				v, err := strconv.Atoi(strings.Replace(text, ",", "", -1))
				if err == nil {
					response.Forks[i] = v
				}
			} else {
				ns.Children().Each(func(k int, nss *goquery.Selection) {
					builtByName := strings.TrimSpace(nss.AttrOr("href", "default"))
					builtBy := struct {
						Name string
						URL  string
					}{
						builtByName[1:],
						constants.GitHubURL + builtByName,
					}
					response.BuiltBy[i] = append(response.BuiltBy[i], builtBy)
				})
			}
		})
	})

	// parse today stars
	divf6.Each(func(i int, s *goquery.Selection) {
		ns := s.Find("span.d-inline-block.float-sm-right")
		text := strings.TrimSpace(ns.Text())
		if text == "" {
			response.RangeStar[i] = 0
		} else {
			splited := strings.Split(text, " ")
			star, err := strconv.Atoi(strings.Replace(splited[0], ",", "", -1))
			if err == nil {
				response.RangeStar[i] = star
			}
		}
	})

	response.Length = divf6.Length()

	return response
}
