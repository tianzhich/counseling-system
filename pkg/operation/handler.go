package operation

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
