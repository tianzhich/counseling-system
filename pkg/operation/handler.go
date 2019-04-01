package operation

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// MarkReadNotification mark the notification to isRead status
func MarkReadNotification(w http.ResponseWriter, r *http.Request) {
	var uid int
	if uid, _ = common.IsUserLogin(r); uid == -1 {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	var resp common.Response

	ids := strings.Split(r.URL.Query().Get("ids"), ",")
	if ok := markReadNotification(ids); !ok {
		resp.Code = 0
		resp.Message = "数据库操作失败"
	} else {
		resp.Code = 1
		resp.Message = "ok"
	}

	resJSON, _ := json.Marshal(resp)
	fmt.Fprintln(w, string(resJSON))
}
