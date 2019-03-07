package info

import (
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// CounselorFilterHandler get all filters about counselor
func CounselorFilterHandler(w http.ResponseWriter, r *http.Request) {
	var res utils.Response
	var cities = getDictInfo(8)
	var methods = getDictInfo(10)
	var topics = getDictInfo(4)

	res.Code = 1
	res.Message = "ok"
	res.Data = filter{Topic: topics, City: cities, Method: methods}

	resJSON, _ := json.Marshal(res)
	fmt.Fprintln(w, string(resJSON))
}
