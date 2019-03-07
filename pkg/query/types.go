package query

// pagination 表示分页属性
type pagination struct {
	pageNum  int
	pageSize int
	total    int
}

// filterOption 表示条件查询
type filterOption struct {
	topic  int
	method int
	city   int
}

// counselor list
type counselor struct {
	UID         int    `json:"uid"`
	Name        int    `json:"name"`
	Gender      int    `json:"gender"`
	Description string `json:"description"`
	WorkYears   int    `json:"workYears"`
	GoodRate    int    `json:"goodRate"`
	Motto       string `json:"motto"`
	AudioPrice  int    `json:"audioPrice"`
	VideoPrice  int    `json:"videoPrice"`
	FtfPrice    int    `json:"ftfPrice"`
	CityID      int    `json:"cityId"`
	CityName    string `json:"cityName"`
}
