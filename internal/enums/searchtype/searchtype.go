package searchtype

type SearchType int

const (
	Movie SearchType = iota
	Series
	Episode
)

var searchTypeToString = map[SearchType]string{
	Movie:   "movie",
	Series:  "series",
	Episode: "episode",
}

func (st SearchType) String() string {
	if s, ok := searchTypeToString[st]; ok {
		return s
	}
	return "unknown"
}
