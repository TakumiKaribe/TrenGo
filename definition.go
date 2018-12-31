package main

type RangeType int

const (
	daily RangeType = iota
	weekly
	monthly
)

func (rt RangeType) String() string {
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
