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
