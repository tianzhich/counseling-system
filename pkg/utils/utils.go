package utils

import (
	"net/http"
)

// CheckErr to check err like db, file operation
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// IsUserLogin to check the user auth
func IsUserLogin(r *http.Request) bool {
	sessions, _ := Store.Get(r, "user_session")
	if auth, ok := sessions.Values["auth"].(bool); !ok || !auth {
		return false
	}
	return true
}
