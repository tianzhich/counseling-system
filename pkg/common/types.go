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
	ID      int    `json:"id"`
	UID     *int   `json:"uID"`
	Type    string `json:"type"`
	IsRead  *int   `json:"isRead"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Time    string `json:"time"`
	Payload int    `json:"payload"`
}

// RecordForm xxx(咨询记录)
type RecordForm struct {
	ID            int    `json:"id"`
	CID           int    `json:"cID"`
	CounselorName string `json:"counselorName"`
	Method        string `json:"method"`
	Times         int    `json:"times"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Gender        int    `json:"gender"`
	Phone         string `json:"phone"`
	ContactPhone  string `json:"contactPhone"`
	ContactName   string `json:"contactName"`
	ContactRel    string `json:"contactRel"`
	Desc          string `json:"desc"`
	Status        string `json:"status"`
	CreateTime    string `json:"createTime"`
	StartTime     string `json:"startTime"`
	Location      string `json:"location"`
	CancelReason1 string `json:"cancelReason1"`
	CancelReason2 string `json:"cancelReason2"`
	RatingScore   int    `json:"ratingScore"`
	RatingText    string `json:"ratingText"`
	Letter        string `json:"letter"`
	Price         int    `json:"price"`
}

// Message xxx(留言)
type Message struct {
	ID           int    `json:"id"`
	SenderID     int    `json:"senderId"`
	SenderName   string `json:"senderName"`
	Receiver     int    `json:"receiver"`
	ReceiverName string `json:"receiverName"`
	Detail       string `json:"detail"`
	Time         string `json:"time"`
}

// Counselor info
type Counselor struct {
	ID          int       `json:"id"`
	UID         int       `json:"uid"`
	Name        string    `json:"name"`
	Gender      int       `json:"gender"`
	Description string    `json:"description"`
	WorkYears   int       `json:"workYears"`
	GoodRate    *int      `json:"goodRate"`
	Motto       string    `json:"motto"`
	AudioPrice  int       `json:"audioPrice"`
	VideoPrice  int       `json:"videoPrice"`
	FtfPrice    int       `json:"ftfPrice"`
	City        *DictInfo `json:"city"`
	Topic       DictInfo  `json:"topic"`
	TopicOther  string    `json:"topicOther"`
	ApplyTime   string    `json:"applyTime"`
}

// User info
type User struct {
	ID         int    `json:"id"`
	CID        *int   `json:"cID"`
	UserName   string `json:"userName"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	CreateTime string `json:"createTime"`
}

// CounselorForm 咨询师注册，更新信息
type CounselorForm struct {
	Name        string            `json:"name"`
	Gender      int               `json:"gender"`
	WorkYears   int               `json:"workYears"`
	Description string            `json:"description"`
	Motto       string            `json:"motto"`
	Detail      []CounselorDetail `json:"detail"`
	AudioPrice  int               `json:"audioPrice"`
	VideoPrice  int               `json:"videoPrice"`
	FtfPrice    int               `json:"ftfPrice"`
	City        string            `json:"city"`
	Topic       string            `json:"topic"`
	OtherTopic  string            `json:"otherTopic"`
}

// CounselorDetail xxx
type CounselorDetail struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Article 文章信息
type Article struct {
	ID       *int   `json:"id"`
	Cover    string `json:"cover"`
	Title    string `json:"title"`
	Excerpt  string `json:"excerpt"`
	Content  string `json:"content"`
	IsDraft  int    `json:"isDraft"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
	CID      int    `json:"cID"`
	PostTime string `json:"postTime"`
}
