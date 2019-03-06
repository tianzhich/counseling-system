package utils

import (
	"fmt"
	"net/http"
)

// CheckErr to check err like db, file operation
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// IsUserLogin to check the user auth and return the user type if user logged in
func IsUserLogin(r *http.Request) (bool, int) {
	sessions, _ := Store.Get(r, "user_session")
	if auth, ok := sessions.Values["auth"].(bool); !ok || !auth {
		return false, -1
	}

	// queryUserType
	var userType = -1
	uid := GetUserID(r)
	rows := QueryDB(fmt.Sprintf("select type from user where id='%v'", uid))
	if rows.Next() {
		rows.Scan(&userType)
	}
	return true, userType
}
