package common

import (
	"counseling-system/pkg/utils"
	"fmt"
	"net/http"
)

// GetAllDictInfoByTypeCode return the counseling cities, methods, topics
func GetAllDictInfoByTypeCode(dictType int) []DictInfo {
	var infos []DictInfo
	var queryStr = fmt.Sprintf("select id, info_name from dict_info where type_code='%v'", dictType)

	infoRows := utils.QueryDB(queryStr)
	for infoRows.Next() {
		var info DictInfo
		infoRows.Scan(&info.ID, &info.Name)
		infos = append(infos, info)
	}
	infoRows.Close()
	return infos
}

// GetDictInfoByID return the dictInfo by id
func GetDictInfoByID(id int) *DictInfo {
	var info DictInfo
	var queryStr = fmt.Sprintf("select info_code, info_name from dict_info where id='%v'", id)

	infoRow := utils.QueryDB(queryStr)
	if infoRow.Next() {
		infoRow.Scan(&info.ID, &info.Name)
	}
	infoRow.Close()
	return &info
}

// IsUserLogin to check the user auth and return the user id and type if user logged in
func IsUserLogin(r *http.Request) (int, int) {
	session, _ := utils.Store.Get(r, "user_session")
	if auth, ok := session.Values["auth"].(bool); !ok || !auth {
		return -1, -1
	}

	// queryUserType
	var userType = -1
	uid, _ := session.Values["uid"].(int)
	rows := utils.QueryDB(fmt.Sprintf("select type from user where id='%v'", uid))
	if rows.Next() {
		rows.Scan(&userType)
	}
	rows.Close()
	return uid, userType
}

// AddNotification to notification table
func AddNotification(no Notification) {
	var insertStr = "insert notification set u_id=?, type=?, title=?, `desc`=?"

	if _, success := utils.InsertDB(insertStr, no.UID, no.Type, no.Title, no.Desc); !success {
		fmt.Println("新增通知失败！")
	}
}

// ReadNotification set notification isRead=true
func ReadNotification(id int) {
	var updateStr = fmt.Sprintf("update notification set is_read=? where id='%v'", id)

	if success := utils.UpdateDB(updateStr, 1); !success {
		fmt.Println("更新通知（已阅读）失败！")
	}
}

// GetUserIDByCID return the user ID for counselor
func GetUserIDByCID(cID int) int {
	var queryStr = fmt.Sprintf("select u_id from counselor where id='%v'", cID)
	var uID int

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&uID)
	} else {
		fmt.Println("查找咨询师u_id失败！")
	}
	rows.Close()
	return uID
}
