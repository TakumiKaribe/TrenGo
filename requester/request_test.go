package requester

import (
	"reflect"
	"testing"
)

func TestMakeURL(t *testing.T) {
	expected := "https://github.com/trending/go?since=daily"
	url := makeURL([]string{githubURL, trending, "go"}, []string{"since=" + Daily.QueryString()})
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

func TestQueryString(t *testing.T) {
	if "daily" != Daily.QueryString() {
		t.Errorf("queryString not daily, got= %s", Daily.QueryString())
	}
	if "weekly" != Weekly.QueryString() {
		t.Errorf("queryString not weekly, got= %s", Weekly.QueryString())
	}
	if "monthly" != Monthly.QueryString() {
		t.Errorf("queryString not monthly, got= %s", Monthly.QueryString())
	}
}

func TestDisplayString(t *testing.T) {
	if "today" != Daily.DisplayString() {
		t.Errorf("queryString not today, got= %s", Daily.DisplayString())
	}
	if "this week" != Weekly.DisplayString() {
		t.Errorf("queryString not this week, got= %s", Weekly.DisplayString())
	}
	if "this month" != Monthly.DisplayString() {
		t.Errorf("queryString not this month, got= %s", Monthly.DisplayString())
	}
}

// TODO: 通信のスタブを用意してparseのテストを行う
