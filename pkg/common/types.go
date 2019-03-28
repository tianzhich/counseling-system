package common

// Response declare the struct of API response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// DictInfo includes info about cities, counseling methods, counseling topics
type DictInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Notification (notification table)
type Notification struct {
	ID     int    `json:"id"`
	UID    int    `json:"uID"`
	Type   string `json:"type"`
	IsRead int    `json:"isRead"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Time   string `json:"time"`
}
