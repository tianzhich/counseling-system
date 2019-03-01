package utils

import "net/http"

// GetUserID return the uid
func GetUserID(r *http.Request) int {
	session, _ := Store.Get(r, "user_session")

	uid, ok := session.Values["uid"].(int)
	if !ok {
		return -1
	}
	return uid
}
