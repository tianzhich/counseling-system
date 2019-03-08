package oauth

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
)

func signup(form SignupForm) (common.Response, int) {
	var username = form.Username
	var pwd = form.Password
	var phone = form.Phone
	var email = form.Email
	var r common.Response

	var uid int64 = -1

	var queryStr string
	var insertStr = "insert user set username=?, password=?, phone=?, email=?"

	queryStr = fmt.Sprintf("select * from user where phone='%v' or email='%v'", phone, email)
	existRows := utils.QueryDB(queryStr)

	queryStr = fmt.Sprintf("select * from user where username='%v'", username)
	repeatRows := utils.QueryDB(queryStr)

	if existRows.Next() {
		r.Code = 0
		r.Message = "邮箱或手机号已被注册，可直接登录"
	} else if repeatRows.Next() {
		r.Code = 0
		r.Message = "用户名已被注册"
	} else {
		var success bool
		uid, success = utils.InsertDB(insertStr, username, pwd, phone, email)
		if success {
			r.Code = 1
			r.Message = "注册成功"
		} else {
			fmt.Println("新增用户失败，数据库插入错误！")
			r.Code = -1
			r.Message = "数据库错误"
			uid = -1
		}
	}

	return r, int(uid)
}

func signin(form SigninForm) (common.Response, int) {
	var expectedPwd string
	var username = form.Username
	var pwd = form.Password
	var r common.Response

	var uid int

	queryStr := fmt.Sprintf("select password, id from user where username='%v'", username)
	rows := utils.QueryDB(queryStr)

	if !rows.Next() {
		r.Code = 0
		r.Message = "用户名或密码错误"
		uid = -1
	} else {
		err := rows.Scan(&expectedPwd, &uid)
		utils.CheckErr(err)
		if expectedPwd != pwd {
			r.Code = 0
			r.Message = "用户名或密码错误"
			uid = -1
		} else {
			r.Code = 1
			r.Message = "login success!"
		}
	}
	return r, uid
}

func applyCounselor(form ApplyForm, uid int) (string, bool) {
	var queryStr = fmt.Sprintf("select * from `counselor` where `u_id` ='%v'", uid)
	var insertStr = "insert counselor set name=?, gender=?, description=?, work_years=?, motto=?, audio_price=?, video_price=?, ftf_price=?, u_id=?"
	var applyRes common.Response

	existRows := utils.QueryDB(queryStr)
	if existRows.Next() {
		applyRes.Code = 0
		applyRes.Message = "账户已绑定咨询师，可直接登录"
	} else {
		if _, status := utils.InsertDB(insertStr, form.Name, form.Gender, form.Description, form.WorkYears, form.Motto, form.AudioPrice, form.VideoPrice, form.FtfPrice, uid); status {
			applyRes.Code = 1
			applyRes.Message = "注册成功"
			handleApplyCity(form.City, uid)
			handleApplyUserType(uid)
			handleApplyTopic(form.Topic, form.OtherTopic, uid)
		} else {
			fmt.Println("新增咨询师失败，数据库插入错误！")
			return "", false
		}
	}

	resJSON, _ := json.Marshal(applyRes)
	return string(resJSON), true
}

// 咨询师入驻 -> 面对面咨询城市的判断处理
func handleApplyCity(city string, uid int) {
	if city == "" {
		return
	}

	var queryStr = fmt.Sprintf("select * from dict_info where `type_code`=8 and `info_name`='%v'", city)
	existRows := utils.QueryDB(queryStr)

	if existRows.Next() {
		return
	}

	cid := utils.QueryDBRow("select count(*) from dict_info where `type_code`=8") + 1
	if rowID, status := utils.InsertDB("insert dict_info set type_code=?, info_code=?, info_name=?", 8, cid, city); status {
		updateStr := fmt.Sprintf("update counselor set city=? where u_id='%v'", uid)
		if updateStatus := utils.UpdateDB(updateStr, rowID); updateStatus {
			return
		}
	} else {
		fmt.Println("新增咨询城市出错！")
	}
}

// 咨询师入驻 -> 修改用户类型
func handleApplyUserType(uid int) {
	var updateStr = fmt.Sprintf("update user set type=? where id='%v'", uid)
	if !utils.UpdateDB(updateStr, 1) {
		fmt.Println("新增咨询师，修改用户类型出错！")
	}
}

// 咨询师入驻 -> 咨询主题(领域)处理
func handleApplyTopic(topic string, otherTopic string, uid int) {
	var queryStr = fmt.Sprintf("select id from dict_info where type_code=4 and info_name='%v'", topic)
	var tid int

	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&tid)
		updateStr := fmt.Sprintf("update counselor set topic=?, topic_other=? where u_id='%v'", uid)
		if success := utils.UpdateDB(updateStr, tid, otherTopic); success {
			return
		}
		fmt.Println("新增咨询师，修改咨询主题出错！")
	}
}
