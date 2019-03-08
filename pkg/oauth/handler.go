package oauth

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

// SignupHandler to handle the req for signup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		res, _ := ioutil.ReadAll(r.Body)

		var formData SignupForm
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		result, success := registerUser(formData)
		if success {
			fmt.Fprintf(w, result)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// SigninHandler to handle the request form signin
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		res, _ := ioutil.ReadAll(r.Body)

		var formData SigninForm
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		// handle user login
		result, uid := signin(formData)

		if result.Code == 1 {
			// session
			// Get a session. Get() always returns a session, even if empty.
			session, err := utils.Store.Get(r, "user_session")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Set expires after one week
			session.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400 * 7,
				HttpOnly: true,
			}

			// Set some session values.
			session.Values["auth"] = true
			session.Values["username"] = formData.Username
			session.Values["uid"] = uid
			// Save it before we write to the response/return from the handler.
			session.Save(r, w)
		}

		resJSON, _ := json.Marshal(result)
		fmt.Fprintf(w, string(resJSON))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// AuthHandler to check if a user is logged in
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var res common.Response

	// Check if user is authenticated
	isLoggedin, userType := common.IsUserLogin(r)
	if !isLoggedin {
		res.Code = 0
		res.Message = "用户未登录"
	} else {
		res.Code = 1
		res.Message = "用户已登录"
		res.Data = authData{UserType: userType}
	}

	resJSON, _ := json.Marshal(res)

	// Print secret message
	fmt.Fprintln(w, string(resJSON))
}

// SignoutHandler to handle the signout
func SignoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "user_session")

	// Revoke user authentication
	session.Values["auth"] = false
	session.Save(r, w)
}

// ApplyHandler to handle the apply for counselor
func ApplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		res, _ := ioutil.ReadAll(r.Body)

		var formData ApplyForm
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		if isLoggedin, _ := common.IsUserLogin(r); !isLoggedin {
			var result = common.Response{Code: -1, Message: "用户未登录，无法执行操作！"}
			resJSON, _ := json.Marshal(result)
			fmt.Fprintf(w, string(resJSON))
			return
		}

		resJSON, success := applyCounselor(formData, common.GetUserID(r))
		if success {
			fmt.Fprintln(w, string(resJSON))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
