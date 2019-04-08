package query

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"fmt"
	"strings"
)

// 咨询师列表
func queryCounselors(p pagination, option *filterOption, orderBy string, like string) (pagination, []counselor) {
	var queryCountStr = "select count(*) from counselor"
	var queryStr = "select id, u_id, name, gender, description, work_years, good_rate, motto, audio_price, video_price, ftf_price, city, topic, topic_other, create_time from counselor"

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
				base = fmt.Sprintf("topic='%v' and ", topic)
			} else {
				base = ""
			}
			queryCountStr += base
			queryStr += base
		}
	}

	// fuzzy query
	if like != "" {
		var likeStr string
		if strings.Contains(queryStr, "where") {
			likeStr = "name LIKE '%" + like + "%'"
			queryCountStr += likeStr
			queryStr += likeStr
		} else {
			likeStr = " where name LIKE '%" + like + "%'"
			queryCountStr += likeStr
			queryStr += likeStr
		}
	}
	queryCountStr = strings.TrimSuffix(queryCountStr, " and ")
	queryStr = strings.TrimSuffix(queryStr, " and ")

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
		rows.Scan(&c.ID, &c.UID, &c.Name, &c.Gender, &c.Description, &c.WorkYears, &c.GoodRate, &c.Motto, &c.AudioPrice, &c.VideoPrice, &c.FtfPrice, &cityID, &topicID, &c.TopicOther, &c.ApplyTime)

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

func queryCounselor(id int) *counselor {
	var queryStr = fmt.Sprintf("select id, u_id, name, gender, description, work_years, good_rate, motto, audio_price, video_price, ftf_price, city, topic, topic_other, create_time from counselor where id='%v'", id)

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		var c counselor
		var cityID *int
		var topicID int
		rows.Scan(&c.ID, &c.UID, &c.Name, &c.Gender, &c.Description, &c.WorkYears, &c.GoodRate, &c.Motto, &c.AudioPrice, &c.VideoPrice, &c.FtfPrice, &cityID, &topicID, &c.TopicOther, &c.ApplyTime)

		// city和topic单独处理
		if cityID == nil {
			c.City = nil
		} else {
			c.City = common.GetDictInfoByID(*cityID)
		}
		c.Topic = *(common.GetDictInfoByID(topicID))
		rows.Close()
		return &c
	}

	return nil
}

// 查看未读通知消息
func queryNotifications(uID int) []common.Notification {
	var notis []common.Notification
	var queryStr = fmt.Sprintf("select id, type, title, `desc`, create_time from notification where u_id=%v and is_read=0 ORDER BY create_time DESC", uID)
	rows := utils.QueryDB(queryStr)
	for rows.Next() {
		var noti common.Notification
		rows.Scan(&noti.ID, &noti.Type, &noti.Title, &noti.Desc, &noti.Time)
		notis = append(notis, noti)
	}
	rows.Close()
	return notis
}

// 查询咨询记录
func queryCounselingRecords(userType int, id int) []common.RecordForm {
	var records []common.RecordForm
	var queryStr = "select id, c_id, method, times, name, age, gender, phone, contact_phone, contact_name, contact_rel, `desc`, status, create_time, start_time, location, cancel_reason1, cancel_reason2, rating_score, rating_text, letter from counseling_record where"
	if userType == 1 {
		queryStr = fmt.Sprintf("%v c_id=%v", queryStr, id)
	} else {
		queryStr = fmt.Sprintf("%v u_id=%v", queryStr, id)
	}

	queryStr += " ORDER BY create_time DESC"

	rows := utils.QueryDB(queryStr)
	for rows.Next() {
		var r common.RecordForm
		rows.Scan(&r.ID, &r.CID, &r.Method, &r.Times, &r.Name, &r.Age, &r.Gender, &r.Phone, &r.ContactPhone, &r.ContactName, &r.ContactRel, &r.Desc, &r.Status, &r.CreateTime, &r.StartTime, &r.Location, &r.CancelReason1, &r.CancelReason2, &r.RatingScore, &r.RatingText, &r.Letter)
		r.CounselorName = common.GetCounselorNameByCID(r.CID)
		r.Method = strings.Replace(r.Method, "\\", "", -1)
		records = append(records, r)
	}
	rows.Close()
	return records
}

// 查询未读留言记录
func queryMessage(uid int) []common.Message {
	var messages []common.Message
	var queryStr = fmt.Sprintf("select id, sender, detail, create_time from message where receiver=%v and is_read=0 ORDER BY create_time DESC", uid)

	rows := utils.QueryDB(queryStr)
	for rows.Next() {
		var m common.Message
		var senderID int
		rows.Scan(&m.ID, &senderID, &m.Detail, &m.Time)
		m.SenderName = common.GetNameByUID(senderID)
		m.SenderID = senderID
		messages = append(messages, m)
	}
	rows.Close()

	return messages
}

// 查询咨询记录，按ID查询
func queryCounselingRecordByID(userType int, id int, rID int) *(common.RecordForm) {
	var r common.RecordForm
	var queryStr = "select id, c_id, method, times, name, age, gender, phone, contact_phone, contact_name, contact_rel, `desc`, status, create_time, start_time, location, cancel_reason1, cancel_reason2, rating_score, rating_text, letter from counseling_record where"
	if userType == 1 {
		queryStr = fmt.Sprintf("%v c_id=%v and id=%v", queryStr, id, rID)
	} else {
		queryStr = fmt.Sprintf("%v u_id=%v and id=%v", queryStr, id, rID)
	}

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&r.ID, &r.CID, &r.Method, &r.Times, &r.Name, &r.Age, &r.Gender, &r.Phone, &r.ContactPhone, &r.ContactName, &r.ContactRel, &r.Desc, &r.Status, &r.CreateTime, &r.StartTime, &r.Location, &r.CancelReason1, &r.CancelReason2, &r.RatingScore, &r.RatingText, &r.Letter)
		r.CounselorName = common.GetCounselorNameByCID(r.CID)
		r.Method = strings.Replace(r.Method, "\\", "", -1)

		rows.Close()
		return &r
	}
	return nil
}
