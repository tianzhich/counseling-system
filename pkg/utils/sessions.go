package utils

import (
	"os"

	"github.com/gorilla/sessions"
)

var (
	key = []byte(os.Getenv("SESSION_KEY"))
	// Store for session management
	Store = sessions.NewCookieStore(key)
)
