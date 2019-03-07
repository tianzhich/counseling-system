package info

import (
	"counseling-system/pkg/common"
	"encoding/json"
	"fmt"
	"net/http"
)

// CounselorFilterHandler get all filters about counselor
func CounselorFilterHandler(w http.ResponseWriter, r *http.Request) {
	var res common.Response
	var cities = common.GetAllDictInfoByTypeCode(8)
	var methods = common.GetAllDictInfoByTypeCode(10)
	var topics = common.GetAllDictInfoByTypeCode(4)

	res.Code = 1
	res.Message = "ok"
	res.Data = filter{Topic: topics, City: cities, Method: methods}

	resJSON, _ := json.Marshal(res)
	fmt.Fprintln(w, string(resJSON))
}
