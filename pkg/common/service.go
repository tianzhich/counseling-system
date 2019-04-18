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
func AddNotification(uid int, no Notification) {
	var insertStr = fmt.Sprintf("insert notification set u_id=%v, type=?, title=?, `desc`=?, payload=?", uid)

	if _, success := utils.InsertDB(insertStr, no.Type, no.Title, no.Desc, no.Payload); !success {
		fmt.Println("新增通知失败！")
	}
}

// GetUserIDByCID return the user ID for counselor
func GetUserIDByCID(cID int) int {
	var queryStr = fmt.Sprintf("select u_id from counselor where id=%v", cID)
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

// GetCounselorIDByUID return the c_id for counselor
func GetCounselorIDByUID(id int) int {
	var queryStr = fmt.Sprintf("select id from counselor where u_id=%v", id)
	var cID int

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&cID)
	} else {
		cID = -1
	}
	rows.Close()
	return cID
}

// GetCounselorNameByCID return counselor name by CID
func GetCounselorNameByCID(id int) string {
	var queryStr = fmt.Sprintf("select name from counselor where id=%v", id)
	var name string

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&name)
	} else {
		fmt.Println("查找咨询师姓名失败！")
	}
	rows.Close()
	return name
}

// GetUserNameByID return the username
func GetUserNameByID(id int) string {
	var queryStr = fmt.Sprintf("select username from user where id=%v", id)
	var username string

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&username)
	}
	rows.Close()
	return username
}

// GetNameByUID return the username or counselor name according to userID
func GetNameByUID(uid int) string {
	var name string
	var queryStr = fmt.Sprintf("select username, type from user where id=%v", uid)

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		var userType int
		rows.Scan(&name, &userType)
		rows.Close()

		if userType == 1 {
			queryStr = fmt.Sprintf("select name from counselor where u_id=%v", uid)
			crows := utils.QueryDB(queryStr)
			if crows.Next() {
				crows.Scan(&name)
				crows.Close()
			} else {
				fmt.Println("根据UID查找咨询师姓名失败！")
			}
		}
	} else {
		fmt.Println("根据UID查找用户名失败！")
	}
	return name
}

// HandleApplyCity 面对面咨询城市的插入更新处理
func HandleApplyCity(city string, uid int) {
	if city == "" {
		return
	}
	var cityID int
	var queryStr = fmt.Sprintf("select id from dict_info where `type_code`=8 and `info_name`='%v'", city)
	existRows := utils.QueryDB(queryStr)

	if existRows.Next() {
		existRows.Scan(&cityID)
		existRows.Close()
	} else {
		infoCode := utils.QueryDBRow("select count(*) from dict_info where `type_code`=8") + 1
		if cID, success := utils.InsertDB("insert dict_info set type_code=?, info_code=?, info_name=?", 8, infoCode, city); success {
			cityID = int(cID)
		} else {
			fmt.Println("新增咨询城市出错！")
			return
		}
	}

	updateStr := fmt.Sprintf("update counselor set city=? where u_id='%v'", uid)
	if success := utils.UpdateDB(updateStr, cityID); !success {
		fmt.Println("更新咨询师所在城市出错")
	}
}
