package requester

type RangeType int

const (
	Daily RangeType = iota
	Weekly
	Monthly
)

func (rt RangeType) QueryString() string {
	switch rt {
	case Daily:
		return "daily"
	case Weekly:
		return "weekly"
	case Monthly:
		return "monthly"
	default:
		return "unknown"
	}
}

func (rt RangeType) DisplayString() string {
	switch rt {
	case Daily:
		return "today"
	case Weekly:
		return "this week"
	case Monthly:
		return "this month"
	default:
		return "unknown"
	}
}
