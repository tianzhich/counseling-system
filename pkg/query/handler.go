package query

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// CounselorListHandler return the Counselor list
func CounselorListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// pagination
		pageNum, err := strconv.Atoi(r.URL.Query().Get("pageNum"))
		pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
		var p = pagination{
			PageNum:  pageNum,
			PageSize: pageSize,
		}

		like := r.URL.Query().Get("like")

		var pp pagination
		var counselors []counselor

		// filters data body
		res, _ := ioutil.ReadAll(r.Body)
		if string(res) == "" {
			pp, counselors = queryCounselors(p, nil, "ORDER BY create_time", like)
		} else {
			var option *filterOption
			err = json.Unmarshal(res, &option)
			utils.CheckErr(err)
			pp, counselors = queryCounselors(p, option, "ORDER BY create_time", like)
		}

		var resp common.Response
		resp.Code = 1
		resp.Message = "ok"
		resp.Data = counselorRespData{
			pagination: pp,
			List:       counselors,
		}

		resJSON, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(resJSON))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// NewlyCounselorsHandler return the Counselor list ordered by join_time
func NewlyCounselorsHandler(w http.ResponseWriter, r *http.Request) {
	// pagination
	pageNum, _ := strconv.Atoi(r.URL.Query().Get("pageNum"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	var p = pagination{
		PageNum:  pageNum,
		PageSize: pageSize,
	}

	var pp pagination
	var counselors []counselor

	pp, counselors = queryCounselors(p, nil, "ORDER BY create_time DESC", "")

	var resp common.Response
	resp.Code = 1
	resp.Message = "ok"
	resp.Data = counselorRespData{
		pagination: pp,
		List:       counselors,
	}

	resJSON, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(resJSON))
}

// CounselorInfoHandler return the info about a single Counselor
func CounselorInfoHandler(w http.ResponseWriter, r *http.Request) {
	var idStr = r.URL.Query().Get("id")
	var resp common.Response
	if idStr == "" {
		resp.Code = 0
		resp.Message = "咨询师不存在"
	} else {
		id, _ := strconv.Atoi(idStr)
		csl := queryCounselor(id)
		if csl == nil {
			resp.Code = 0
			resp.Message = "咨询师不存在"
		} else {
			resp.Code = 1
			resp.Message = "ok"
			resp.Data = *csl
		}
	}

	resJSON, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(resJSON))
}

// NotificationHandler return the messages and notifications of user
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	var uid int
	if uid, _ = common.IsUserLogin(r); uid == -1 {
		http.Error(w, "Not Loggin", http.StatusUnauthorized)
		return
	}

	var resp common.Response
	resp.Code = 1
	resp.Message = "ok"
	resp.Data = queryNotifications(uid)

	resJSON, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(resJSON))
}

// CounselingRecordHandler return all the counseling record
func CounselingRecordHandler(w http.ResponseWriter, r *http.Request) {
	var uid int
	var cid int
	var userType int
	var resp common.Response

	if uid, userType = common.IsUserLogin(r); uid == -1 {
		http.Error(w, "Not Loggin", http.StatusUnauthorized)
		return
	}

	// query by id
	rid, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if rid != 0 {
		var uuid int
		var record *(common.RecordForm)
		if userType == 1 {
			uuid = common.GetCounselorIDByUID(uid)
		} else {
			uuid = uid
		}
		if record = queryCounselingRecordByID(userType, uuid, rid); record != nil {
			resp.Code = 1
			resp.Message = "ok"
			resp.Data = *(record)
		} else {
			resp.Code = 0
			resp.Message = "非法查询ID"
		}
	} else {
		var records []common.RecordForm
		if userType == 1 {
			cid = common.GetCounselorIDByUID(uid)
			records = queryCounselingRecords(userType, cid)
		} else {
			records = queryCounselingRecords(userType, uid)
		}
		resp.Code = 1
		resp.Message = "ok"
		resp.Data = records
	}

	resJSON, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(resJSON))
}

// MessageHandler return all the message
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	var uid int
	if uid, _ = common.IsUserLogin(r); uid == -1 {
		http.Error(w, "Not Loggin", http.StatusUnauthorized)
		return
	}

	var resp common.Response
	resp.Code = 1
	resp.Message = "ok"
	resp.Data = queryMessage(uid)
	resJSON, _ := json.Marshal(resp)

	fmt.Fprintf(w, string(resJSON))
}
