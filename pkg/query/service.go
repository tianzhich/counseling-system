package query

import (
	"counseling-system/pkg/utils"
	"fmt"
	"strings"
)

func queryCounselor(p pagination, option filterOption) {
	var queryCountStr = "select count(*) from counselor where "
	var queryStr = "select u_id, name, gender, description, work_years, good_rate, motto, audio_price, video_price, ftf_price, city from counselor where "

	// append queryStr
	var base string
	if option.method == 1 {
		base = "ftf_price>0 and "
	} else if option.method == 2 {
		base = "video_price>0 and "
	} else if option.method == 3 {
		base = "audio_price>0 and "
	} else {
		base = ""
	}
	queryCountStr += base
	queryStr += base
	if option.city > 0 {
		base = fmt.Sprintf("city='%v' and ", option.city)
	} else {
		base = ""
	}
	queryCountStr += base
	queryStr += base
	if option.topic > 0 {
		base = fmt.Sprintf("topic='%v' and ", option.topic)
	} else {
		base = ""
	}
	queryCountStr += base
	queryStr += base

	queryCountStr = strings.TrimSuffix(queryCountStr, " and ")
	queryStr = strings.TrimSuffix(queryStr, " and ")

	// db op
	count := utils.QueryDBRow(queryCountStr)
	if count != 0 {

	}
}
