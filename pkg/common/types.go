package common

var (
	// ServerErrorMessage xxx
	ServerErrorMessage = "内部服务器错误，请稍后重试！"
)

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
	UID    *int   `json:"uID"`
	Type   string `json:"type"`
	IsRead *int   `json:"isRead"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Time   string `json:"time"`
}

// RecordForm xxx(咨询记录)
type RecordForm struct {
	ID           int    `json:"id"`
	CID          *int   `json:"cID"`
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
	Status       string `json:"status"`
	CreateTime   string `json:"createTime"`
}

// Message xxx(留言)
type Message struct {
	ID           int    `json:"id"`
	Sender       int    `json:"sender"`
	SenderName   string `json:"senderName"`
	Receiver     int    `json:"receiver"`
	ReceiverName string `json:"receiverName"`
	Detail       string `json:"detail"`
	Time         string `json:"time"`
}
