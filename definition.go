package main

const githubURL string = "https://github.com"
const trending string = "/trending"

type RangeType int

const (
	daily RangeType = iota
	weekly
	monthly
)

func (rt RangeType) QueryString() string {
	switch rt {
	case daily:
		return "daily"
	case weekly:
		return "weekly"
	case monthly:
		return "monthly"
	default:
		return "unknown"
	}
}

func (rt RangeType) DisplayString() string {
	switch rt {
	case daily:
		return "today"
	case weekly:
		return "this week"
	case monthly:
		return "this month"
	default:
		return "unknown"
	}
}

// TODO: 排他的なオプションにする
// e.g.) main.go -d -w -m  => error
// e.g.) main.go -w  => weekly
func parseRangeType(d bool, w bool, m bool) RangeType {
	if d {
		return daily
	}

	if w {
		return weekly
	}

	if m {
		return monthly
	}

	return daily
}
