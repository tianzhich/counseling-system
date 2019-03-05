package utils

import "net/http"

// Response declare the struct of API response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// GetUserID return the uid
func GetUserID(r *http.Request) int {
	session, _ := Store.Get(r, "user_session")

	uid, ok := session.Values["uid"].(int)
	if !ok {
		return -1
	}
	return uid
}
