package operation

// RecordForm xxx(咨询记录)
type RecordForm struct {
	CID          int    `json:"cID"`
	Method       string `json:"method"`
	Times        int    `json:"times"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Gender       int    `json:"gender"`
	Phone        string `json:"phone"`
	ContactPhone string `json:"contactPhone"`
	ContactName  string `json:"contactName"`
	ContactRel   string `json:"contactRel"`
	Desc         string `json:"desc"`
	Status       int    `json:"status"`
}
