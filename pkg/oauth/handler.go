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

	queryStr = fmt.Sprintf("select * from user where phone='%v' or email='%v'", phone, email)
	existRows := utils.QueryDB(queryStr)

	queryStr = fmt.Sprintf("select * from user where username='%v'", username)
	repeatRows := utils.QueryDB(queryStr)

	if existRows.Next() {
		registerRes.Code = 0
		registerRes.Message = "邮箱或手机号已被注册，可直接登录"
	} else if repeatRows.Next() {
		registerRes.Code = 0
		registerRes.Message = "用户名已被注册"
	} else {
		if utils.InsertDB(insertStr, username, pwd, phone, email) {
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

func signin(form SigninForm) (result, int) {
	var expectedPwd string
	var username = form.Username
	var pwd = form.Password
	var res result

	var uid int

	queryStr := fmt.Sprintf("select password, id from user where username='%v'", username)
	rows := utils.QueryDB(queryStr)

	if !rows.Next() {
		res.Code = 0
		res.Message = "用户名或密码错误"
		uid = -1
	} else {
		err := rows.Scan(&expectedPwd, &uid)
		utils.CheckErr(err)
		if expectedPwd != pwd {
			res.Code = 0
			res.Message = "用户名或密码错误"
			uid = -1
		} else {
			res.Code = 1
			res.Message = "login success!"
		}
	}
	return res, uid
}

// AuthHandler to check if a user is logged in
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var res result

	// Check if user is authenticated
	if !utils.IsUserLogin(r) {
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

// ApplyHandler to handle the apply for consultant
func ApplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		res, _ := ioutil.ReadAll(r.Body)

		var formData ApplyForm
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		if !utils.IsUserLogin(r) {
			var result = result{Code: -1, Message: "用户未登录，无法执行操作！"}
			resJSON, _ := json.Marshal(result)
			fmt.Fprintf(w, string(resJSON))
			return
		}

		resJSON := applyConsultant(formData, utils.GetUserID(r))
		fmt.Fprintln(w, string(resJSON))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func applyConsultant(form ApplyForm, uid int) string {
	var queryStr = fmt.Sprintf("select * from `consultant` where `u_id` ='%v'", uid)
	var insertStr = "insert consultant set name=?, gender=?, description=?, work_years=?, motto=?, audio_price=?, video_price=?, ftf_price=?, u_id=?"
	var applyRes result

	existRows := utils.QueryDB(queryStr)
	if existRows.Next() {
		applyRes.Code = 0
		applyRes.Message = "账户已绑定咨询师，可直接登录"
	} else {
		if utils.InsertDB(insertStr, form.Name, form.Gender, form.Description, form.WorkYears, form.Motto, form.AudioPrice, form.VideoPrice, form.FtfPrice, uid) {
			applyRes.Code = 1
			applyRes.Message = "注册成功"
			handleApplyCity(form.City)
		} else {
			applyRes.Code = -1
			applyRes.Message = "数据库连接错误，请稍后重试！"
		}
	}

	resJSON, _ := json.Marshal(applyRes)
	return string(resJSON)
}

// 面对面咨询城市的判断处理
func handleApplyCity(city string) {
	var queryStr = fmt.Sprintf("select * from dict_info where `type_code`=8 and `info_name`='%v'", city)
	existRows := utils.QueryDB(queryStr)

	if existRows.Next() {
		return
	}

	cid := utils.QueryDBRow("select count(*) from dict_info where `type_code`=8") + 1
	if utils.InsertDB("insert dict_info set typecode=?, info_code=?, info_name=?", 8, cid, city) {
		return
	}

	fmt.Println("新增咨询城市出错！")
}
