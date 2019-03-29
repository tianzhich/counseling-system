package query

import (
	"counseling-system/pkg/common"
)

// pagination 表示分页属性
type pagination struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

// filterOption 表示条件查询
type filterOption struct {
	Topic  *int `json:"topic"`
	Method *int `json:"method"`
	City   *int `json:"city"`
}

// counselor list
type counselor struct {
	ID          int              `json:"id"`
	UID         int              `json:"uid"`
	Name        string           `json:"name"`
	Gender      int              `json:"gender"`
	Description string           `json:"description"`
	WorkYears   int              `json:"workYears"`
	GoodRate    *int             `json:"goodRate"`
	Motto       string           `json:"motto"`
	AudioPrice  int              `json:"audioPrice"`
	VideoPrice  int              `json:"videoPrice"`
	FtfPrice    int              `json:"ftfPrice"`
	City        *common.DictInfo `json:"city"`
	Topic       common.DictInfo  `json:"topic"`
	TopicOther  string           `json:"topicOther"`
	ApplyTime   string           `json:"applyTime"`
}

// counselor response data
type counselorRespData struct {
	pagination
	List []counselor `json:"list"`
}

type profileAll struct {
	Notification []common.Notification `json:"notification"`
}
