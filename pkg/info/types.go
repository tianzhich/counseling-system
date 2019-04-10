package info

import (
	"counseling-system/pkg/common"
)

type filter struct {
	Topic  []common.DictInfo `json:"topic"`
	Method []common.DictInfo `json:"method"`
	City   []common.DictInfo `json:"city"`
}

type preInfo struct {
	ID         int    `json:"id"`
	CID        *int   `json:"cID"`
	UserName   string `json:"userName"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	CreateTime string `json:"createTime"`
}

type preCounselorInfo struct {
	ID int
}
