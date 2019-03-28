package info

import (
	"counseling-system/pkg/common"
	"encoding/json"
	"fmt"
	"net/http"
)

// const
const (
	CITY   = 8
	METHOD = 10
	TOPIC  = 4
)

// CounselorFilterHandler get all filters about counselor
func CounselorFilterHandler(w http.ResponseWriter, r *http.Request) {
	var res common.Response
	var cities = common.GetAllDictInfoByTypeCode(CITY)
	var methods = common.GetAllDictInfoByTypeCode(METHOD)
	var topics = common.GetAllDictInfoByTypeCode(TOPIC)

	res.Code = 1
	res.Message = "ok"
	res.Data = filter{Topic: topics, City: cities, Method: methods}

	resJSON, _ := json.Marshal(res)
	fmt.Fprintln(w, string(resJSON))
}
