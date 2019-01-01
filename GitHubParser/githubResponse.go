package GitHubParser

import "github.com/maka-nai/trengo/RangeType"

type GitHubResponse struct {
	lang         string
	rangeType    RangeType.RangeType
	developers   [50]string
	names        [50]string
	URLs         [50]string
	descriptions [50]string
	languages    [50]string
	stars        [50]int
	forks        [50]int
	rangeStar    [50]int
	length       int
}
