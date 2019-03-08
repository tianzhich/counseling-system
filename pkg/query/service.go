package query

import (
	"counseling-system/pkg/utils"
	"fmt"
	"strings"
)

func queryCounselor(p pagination, option filterOption, orderBy string) (pagination, []counselor) {
	var queryCountStr = "select count(*) from counselor where "
	var queryStr = "select u_id, name, gender, description, work_years, good_rate, motto, audio_price, video_price, ftf_price, city, topic, topic_other from counselor where "

	var firstRecordIndex = (p.pageNum - 1) * 10

	// append with query params
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
		base = fmt.Sprintf("topic='%v'", option.topic)
	} else {
		base = ""
	}
	queryCountStr += base
	queryStr += base
	queryCountStr = strings.TrimSuffix(queryCountStr, " and ")
	queryStr = strings.TrimSuffix(queryStr, " and ")

	// append using pagination
	queryStr += fmt.Sprintf(" LIMIT %v,%v ORDER BY %v", firstRecordIndex, p.pageSize, orderBy)

	// db op
	count := utils.QueryDBRow(queryCountStr)
	var pp = pagination{pageNum: p.pageNum, pageSize: p.pageSize, total: count}

	// empty
	if count == 0 || count <= firstRecordIndex {
		var emptyCounslor = []counselor{}
		return pp, emptyCounslor
	}

	var counselorList []counselor
	rows := utils.QueryDB(queryCountStr)
	for rows.Next() {
		var c counselor
		rows.Scan(&c.UID, &c.Name, &c.Gender, &c.Description, &c.WorkYears, &c.GoodRate, &c.Motto, &c.AudioPrice, &c.VideoPrice, &c.FtfPrice, &c.City, &c.Topic, &c.TopicOther)

		counselorList = append(counselorList, c)
	}

	return pp, counselorList
}
