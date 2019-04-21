package requester

import (
	"log"
	"net/http"
	"strings"

	"trengo/constants"
	"trengo/model"
	"trengo/requester/condition"
	"trengo/service"
)

const trending string = "trending"

// Request -
func Request(rt condition.RangeType, lang string) model.Response {
	url := makeURL([]string{constants.GitHubURL, trending, lang}, []string{"since=" + queryTypeString(rt)})

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return service.Parse(res.Body, rt)
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

func queryTypeString(rt condition.RangeType) string {
	switch rt {
	case condition.Daily:
		return "daily"
	case condition.Weekly:
		return "weekly"
	case condition.Monthly:
		return "monthly"
	default:
		return ""
	}
}
