package requester

import (
	"reflect"
	"testing"

	"trengo/constants"
	"trengo/requester/condition"
)

func TestMakeURL(t *testing.T) {
	expected := "https://github.com/trending/go?since=daily"
	url := makeURL([]string{constants.GitHubURL, trending, "go"}, []string{"since=" + queryTypeString(condition.Daily)})
	if expected != url {
		t.Errorf("makeURL not %s, got= %s", expected, url)
	}
}

func TestFilter(t *testing.T) {
	expected := []string{"a", "b", "c", "d", "e"}
	target := []string{"a", "target", "b", "target", "c", "target", "d", "target", "e"}
	filtered := filter(target, func(str string) bool { return str != "target" })
	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("makeURL not %q, got= %q", expected, filtered)
	}
}

func TestQueryTypeString(t *testing.T) {
	if "daily" != queryTypeString(condition.Daily) {
		t.Errorf("queryString not daily, got= %s", queryTypeString(condition.Daily))
	}
	if "weekly" != queryTypeString(condition.Weekly) {
		t.Errorf("queryString not weekly, got= %s", queryTypeString(condition.Weekly))
	}
	if "monthly" != queryTypeString(condition.Monthly) {
		t.Errorf("queryString not monthly, got= %s", queryTypeString(condition.Monthly))
	}
}
