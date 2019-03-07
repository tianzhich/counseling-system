package info

// dictInfo struct includes City, CounselingMethod,
type dictInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type filter struct {
	Topic  []dictInfo `json:"topic"`
	Method []dictInfo `json:"method"`
	City   []dictInfo `json:"city"`
}
