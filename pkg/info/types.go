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
	common.User
}

type askTag struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	SubTags *[]askTag `json:"subTags"`
}
