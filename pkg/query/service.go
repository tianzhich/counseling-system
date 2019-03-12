package query

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"fmt"
	"strings"
)

func queryCounselor(p pagination, option *filterOption, orderBy string) (pagination, []counselor) {
	var queryCountStr = "select count(*) from counselor"
	var queryStr = "select u_id, name, gender, description, work_years, good_rate, motto, audio_price, video_price, ftf_price, city, topic, topic_other, create_time from counselor"

	var firstRecordIndex = (p.PageNum - 1) * p.PageSize

	// append with query params
	if option != nil {
		var base string
		if option.Method != nil || option.City != nil || option.Topic != nil {
			queryCountStr += " where "
			queryStr += " where "
		}
		// method
		if option.Method != nil {
			methodID := *(option.Method)
			methodCode := (*(common.GetDictInfoByID(methodID))).ID
			if methodCode == 1 {
				base = "ftf_price>0 and "
			} else if methodCode == 2 {
				base = "video_price>0 and "
			} else if methodCode == 3 {
				base = "audio_price>0 and "
			} else {
				base = ""
			}
			queryCountStr += base
			queryStr += base
		}
		// city
		if option.City != nil {
			city := *(option.City)
			if city > 0 {
				base = fmt.Sprintf("city='%v' and ", city)
			} else {
				base = ""
			}
			queryCountStr += base
			queryStr += base
		}
		// topic
		if option.Topic != nil {
			topic := *(option.Topic)
			if topic > 0 {
				base = fmt.Sprintf("topic='%v'", topic)
			} else {
				base = ""
			}
			queryCountStr += base
			queryStr += base
		}
		queryCountStr = strings.TrimSuffix(queryCountStr, " and ")
		queryStr = strings.TrimSuffix(queryStr, " and ")
	}

	// append using pagination and order
	queryStr += fmt.Sprintf(" %v LIMIT %v,%v", orderBy, firstRecordIndex, p.PageSize)

	// db op
	count := utils.QueryDBRow(queryCountStr)
	var pp = pagination{PageNum: p.PageNum, PageSize: p.PageSize, Total: count}

	// empty
	if count == 0 || count <= firstRecordIndex {
		var emptyCounslor = []counselor{}
		return pp, emptyCounslor
	}

	var counselorList []counselor
	rows := utils.QueryDB(queryStr)
	for rows.Next() {
		var c counselor
		var cityID *int
		var topicID int
		rows.Scan(&c.UID, &c.Name, &c.Gender, &c.Description, &c.WorkYears, &c.GoodRate, &c.Motto, &c.AudioPrice, &c.VideoPrice, &c.FtfPrice, &cityID, &topicID, &c.TopicOther, &c.ApplyTime)

		// city和topic单独处理
		if cityID == nil {
			c.City = nil
		} else {
			c.City = common.GetDictInfoByID(*cityID)
		}
		c.Topic = *(common.GetDictInfoByID(topicID))

		counselorList = append(counselorList, c)
	}
	rows.Close()

	return pp, counselorList
}
