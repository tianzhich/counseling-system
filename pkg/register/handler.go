package register

import (
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handler to handle the req for registration
func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.AllowCors(&w)

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method == "POST" {
		res, _ := ioutil.ReadAll(r.Body)

		var formData Form
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		result := registerUser(formData)
		fmt.Fprintf(w, result)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func registerUser(form Form) string {
	var username = form.Username
	var pwd = form.Password
	var phone = form.Phone
	var email = form.Email

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
		return "邮箱或手机号已被注册，可直接登录"
	} else if repeatRows.Next() {
		return "用户名已被注册"
	} else {
		stmt, err := db.Prepare(insertStr)
		utils.CheckErr(err)
		defer stmt.Close()

		res, err := stmt.Exec(username, pwd, phone, email)
		utils.CheckErr(err)

		rows, _ := res.RowsAffected()
		if rows == 1 {
			return "注册成功"
		}
		return "注册失败，请稍后重试"
	}
}
