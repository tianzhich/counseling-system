package oauth

import (
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
