package model

// Response -
type Response struct {
	Lang         string
	Developers   [50]string
	Names        [50]string
	URLs         [50]string
	Descriptions [50]string
	Languages    [50]string
	Stars        [50]int
	Forks        [50]int
	BuiltBy      [50][]struct {
		Name string
		URL  string
	}
	RangeStar [50]int
	Length    int
}
