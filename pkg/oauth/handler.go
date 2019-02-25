package oauth

import (
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

type result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SignupHandler to handle the req for signup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		res, _ := ioutil.ReadAll(r.Body)

		var formData SignupForm
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		result := registerUser(formData)
		fmt.Fprintf(w, result)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func registerUser(form SignupForm) string {
	var username = form.Username
	var pwd = form.Password
	var phone = form.Phone
	var email = form.Email
	var registerRes result

	var queryStr string
	var insertStr = "insert user set username=?, password=?, phone=?, email=?"

	db := utils.InitialDb()
	defer db.Close()

	queryStr = fmt.Sprintf("select * from user where phone='%v' or email='%v'", phone, email)
	existRows, existErr := db.Query(queryStr)
	utils.CheckErr(existErr)

	queryStr = fmt.Sprintf("select * from user where username='%v'", username)
	repeatRows, repeatErr := db.Query(queryStr)
	utils.CheckErr(repeatErr)

	if existRows.Next() {
		registerRes.Code = 0
		registerRes.Message = "邮箱或手机号已被注册，可直接登录"
	} else if repeatRows.Next() {
		registerRes.Code = 0
		registerRes.Message = "用户名已被注册"
	} else {
		stmt, err := db.Prepare(insertStr)
		utils.CheckErr(err)
		defer stmt.Close()

		res, err := stmt.Exec(username, pwd, phone, email)
		utils.CheckErr(err)

		rows, _ := res.RowsAffected()
		if rows == 1 {
			registerRes.Code = 1
			registerRes.Message = "注册成功"
		} else {
			registerRes.Code = -1
			registerRes.Message = "数据库连接错误！"
		}
	}

	resJSON, _ := json.Marshal(registerRes)
	return string(resJSON)
}

// SigninHandler to handle the request form signin
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		res, _ := ioutil.ReadAll(r.Body)

		var formData SigninForm
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		// handle user login
		result := signin(formData)

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
			// Save it before we write to the response/return from the handler.
			session.Save(r, w)
		}

		resJSON, _ := json.Marshal(result)
		fmt.Fprintf(w, string(resJSON))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func signin(form SigninForm) result {
	var expectedPwd string
	var username = form.Username
	var pwd = form.Password
	var res result
	var db = utils.InitialDb()
	defer db.Close()

	var queryStr = fmt.Sprintf("select password from user where username='%v'", username)
	rows, err := db.Query(queryStr)
	utils.CheckErr(err)

	if !rows.Next() {
		res.Code = 0
		res.Message = "用户名或密码错误"
	} else {
		err := rows.Scan(&expectedPwd)
		utils.CheckErr(err)
		if expectedPwd != pwd {
			res.Code = 0
			res.Message = "用户名或密码错误"
		} else {
			res.Code = 1
			res.Message = "login success!"
		}
	}
	return res
}

// AuthHandler to check if a user is logged in
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var res result
	session, _ := utils.Store.Get(r, "user_session")

	// Check if user is authenticated
	if auth, ok := session.Values["auth"].(bool); !ok || !auth {
		res.Code = 0
		res.Message = "用户未登录"
	} else {
		res.Code = 1
		res.Message = "用户已登录"
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
