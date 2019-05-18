package info

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"fmt"
	"strconv"
)

func getLogginUserInfo(uid int) preInfo {
	var info preInfo
	var queryStr = fmt.Sprintf("select id, username, phone, email, create_time from user where id=%v", uid)

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&info.ID, &info.UserName, &info.Phone, &info.Email, &info.CreateTime)
	}
	rows.Close()

	if cid := common.GetCounselorIDByUID(uid); cid != -1 {
		info.CID = &cid
	}
	return info
}

func getLogginCounselorInfo(uid int) common.Counselor {
	var c common.Counselor
	var queryStr = fmt.Sprintf("select id, u_id, name, gender, description, work_years, good_rate, motto, audio_price, video_price, ftf_price, city, topic, topic_other, create_time from counselor where u_id=%v", uid)

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		var cityID *int
		var topicID int
		rows.Scan(&c.ID, &c.UID, &c.Name, &c.Gender, &c.Description, &c.WorkYears, &c.GoodRate, &c.Motto, &c.AudioPrice, &c.VideoPrice, &c.FtfPrice, &cityID, &topicID, &c.TopicOther, &c.ApplyTime)

		if cityID != nil {
			c.City = common.GetDictInfoByID(*(cityID))
		} else {
			c.City = nil
		}
		c.Topic = *(common.GetDictInfoByID(topicID))
	}
	rows.Close()
	return c
}

// 文章草稿
func getArticleDraft(cID int) *(common.Article) {
	var queryStr = fmt.Sprintf("select id, cover, title, content, category, tags from article where is_draft=1 and c_id=%v", cID)
	var a common.Article

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&a.ID, &a.Cover, &a.Title, &a.Content, &a.Category, &a.Tags)
		rows.Close()
		return &a
	}
	rows.Close()
	return nil
}

func getAskTags() []common.AskTag {
	var queryTagParentStr = "select parent_id, parent_name from ask_tag GROUP BY `parent_id`"
	var queryTagStr string
	var at []common.AskTag

	rows1 := utils.QueryDB(queryTagParentStr)
	for rows1.Next() {
		var att common.AskTag
		rows1.Scan(&att.ID, &att.Name)
		at = append(at, att)
	}
	rows1.Close()

	for index, p := range at {
		queryTagStr = fmt.Sprintf("select id, name from ask_tag where parent_id='%v'", p.ID)
		rows2 := utils.QueryDB(queryTagStr)
		var subAt []common.AskTag
		for rows2.Next() {
			var subAtt common.AskTag
			var id int
			rows2.Scan(&id, &subAtt.Name)
			subAtt.ID = strconv.Itoa(id)
			subAt = append(subAt, subAtt)
		}
		rows2.Close()
		at[index].SubTags = &subAt
	}
	return at
}
