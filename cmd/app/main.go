package main

import (
	"counseling-system/pkg/info"
	"counseling-system/pkg/oauth"
	"counseling-system/pkg/operation"
	"counseling-system/pkg/query"

	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// API handler
	oauthHandlers(mux)
	infoHandlers(mux)
	queryHandlers(mux)
	operationHandlers(mux)

	log.Println("Listening on port 8081 ...")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}

func oauthHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api/oauth/signup", oauth.SignupHandler)
	mux.HandleFunc("/api/oauth/signin", oauth.SigninHandler)
	mux.HandleFunc("/api/oauth/auth", oauth.AuthHandler)
	mux.HandleFunc("/api/oauth/signout", oauth.SignoutHandler)
	mux.HandleFunc("/api/oauth/apply", oauth.ApplyHandler)
}

func infoHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api/info/counselingFilters", info.CounselorFilterHandler)
}

func queryHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api/query/counselorList", query.CounselorListHandler)
	mux.HandleFunc("/api/query/newlyCounselors", query.NewlyCounselorsHandler)
	mux.HandleFunc("/api/query/counselor", query.CounselorInfoHandler)
	mux.HandleFunc("/api/query/notifications", query.NotificationHandler)
	mux.HandleFunc("/api/query/messages", query.MessageHandler)
	mux.HandleFunc("/api/query/counselingRecords", query.CounselingRecordHandler)
}

func operationHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api/operation/appoint", operation.AppointHandler)
	mux.HandleFunc("/api/operation/markRead", operation.MarkReadHandler)
	mux.HandleFunc("/api/operation/addMessage", operation.AddMessageHandler)
}
