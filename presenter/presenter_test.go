package presenter

import (
	"testing"
	"trengo/requester/condition"
)

func TestRangeTypeString(t *testing.T) {
	if "today" != rangeTypeString(condition.Daily) {
		t.Errorf("queryString not today, got= %s", rangeTypeString(condition.Daily))
	}
	if "this week" != rangeTypeString(condition.Weekly) {
		t.Errorf("queryString not this week, got= %s", rangeTypeString(condition.Weekly))
	}
	if "this month" != rangeTypeString(condition.Monthly) {
		t.Errorf("queryString not this month, got= %s", rangeTypeString(condition.Monthly))
	}
}
