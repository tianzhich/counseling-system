package common

import (
	"counseling-system/pkg/utils"
	"fmt"
	"net/http"
)

// GetUserID return the uid
func GetUserID(r *http.Request) int {
	session, _ := utils.Store.Get(r, "user_session")

	uid, ok := session.Values["uid"].(int)
	if !ok {
		return -1
	}
	return uid
}

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

// IsUserLogin to check the user auth and return the user type if user logged in
func IsUserLogin(r *http.Request) (bool, int) {
	sessions, _ := utils.Store.Get(r, "user_session")
	if auth, ok := sessions.Values["auth"].(bool); !ok || !auth {
		return false, -1
	}

	// queryUserType
	var userType = -1
	uid := GetUserID(r)
	rows := utils.QueryDB(fmt.Sprintf("select type from user where id='%v'", uid))
	if rows.Next() {
		rows.Scan(&userType)
	}
	rows.Close()
	return true, userType
}
