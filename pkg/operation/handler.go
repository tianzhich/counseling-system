package operation

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// AppointHandler to handle the counseling appointment record
func AppointHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 登录验证
		var uid int
		if uid, _ = common.IsUserLogin(r); uid == -1 {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		res, _ := ioutil.ReadAll(r.Body)

		var formData common.RecordForm
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		resJSON, success := addCounselingRecord(formData, uid)
		if success {
			fmt.Fprintln(w, string(resJSON))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// MarkReadHandler mark the notification or message to isRead status
func MarkReadHandler(w http.ResponseWriter, r *http.Request) {
	var uid int
	if uid, _ = common.IsUserLogin(r); uid == -1 {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	var resp common.Response

	ids := strings.Split(r.URL.Query().Get("ids"), ",")
	t := r.URL.Query().Get("type")
	if ok := markRead(ids, t); !ok {
		resp.Code = 0
		resp.Message = "数据库操作失败"
	} else {
		resp.Code = 1
		resp.Message = "ok"
	}

	resJSON, _ := json.Marshal(resp)
	fmt.Fprintln(w, string(resJSON))
}

// AddMessageHandler add message
func AddMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var uid int
		if uid, _ = common.IsUserLogin(r); uid == -1 {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		var resp common.Response
		var m common.Message
		res, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(res, &m)
		utils.CheckErr(err)

		if m.Receiver == uid {
			resp.Code = -1
			resp.Message = "无法向自己私信！"
		} else {
			if ok := addMessage(uid, m); ok {
				resp.Code = 1
				resp.Message = "ok"
			} else {
				resp.Code = 0
				resp.Message = common.ServerErrorMessage
			}
		}
		resJSON, _ := json.Marshal(resp)
		fmt.Fprintln(w, string(resJSON))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// AppointProcessHandler 处理咨询流程进度，0表示拒绝当前流程，1表示同意
func AppointProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var uid int
		var userType int
		if uid, userType = common.IsUserLogin(r); uid == -1 {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		var pathArr = strings.Split(r.URL.Path, "/")
		if len(pathArr) != 6 {
			http.Error(w, "Invalid request", http.StatusForbidden)
			return
		}
		recordID, _ := strconv.Atoi(pathArr[4])
		operation, _ := strconv.Atoi(pathArr[5])

		var args processArgs
		form, _ := ioutil.ReadAll(r.Body)
		if len(form) > 0 {
			err := json.Unmarshal(form, &args)
			utils.CheckErr(err)
		}

		code, msg := appointProcess(uid, userType, recordID, operation, args)
		var resp = common.Response{Code: code, Message: msg}

		resJSON, _ := json.Marshal(resp)
		fmt.Fprintln(w, string(resJSON))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// UpdateInfoHandler 更新用户或咨询师信息
func UpdateInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var uid int
		var userType int
		if uid, userType = common.IsUserLogin(r); uid == -1 {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		var infoType = r.URL.Query().Get("type")
		body, err := ioutil.ReadAll(r.Body)
		utils.CheckErr(err)
		var resp common.Response

		if infoType == "1" {
			if userType == 1 {
				var data common.CounselorForm
				err := json.Unmarshal(body, &data)
				utils.CheckErr(err)
				if success := updateCounselorInfo(uid, data); !success {
					resp.Code = 0
					resp.Message = "更新咨询师信息出错，请联系管理员"
				} else {
					resp.Code = 1
					resp.Message = "ok"
				}
			} else {
				http.Error(w, "No access", http.StatusForbidden)
				return
			}
		} else if infoType == "2" {
			var data common.User
			err := json.Unmarshal(body, &data)
			utils.CheckErr(err)
			if sucess := updateUserInfo(uid, data); !sucess {
				resp.Code = 0
				resp.Message = "更新用户信息出错，请联系管理员"
			} else {
				resp.Code = 1
				resp.Message = "ok"
			}
		} else {
			http.Error(w, "Invalid request param", http.StatusBadRequest)
			return
		}
		resJSON, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(resJSON))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// AddArticleHandler 保存文章草稿和提交文章
func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var uid int
		var userType int
		if uid, userType = common.IsUserLogin(r); uid == -1 {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		} else if userType == 2 {
			http.Error(w, "Access not allowed", http.StatusForbidden)
		}

		// handler
		var cID = common.GetCounselorIDByUID(uid)
		var args common.Article
		data, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(data, &args)
		utils.CheckErr(err)
		var resp common.Response

		if success := articleProcess(cID, args); success {
			resp.Code = 1
			resp.Message = "ok"
		} else {
			resp.Code = 0
			resp.Message = common.ServerErrorMessage
		}

		resJSON, _ := json.Marshal(resp)
		fmt.Fprintln(w, string(resJSON))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
