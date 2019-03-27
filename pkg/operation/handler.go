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
		if isLogin, _ := common.IsUserLogin(r); !isLogin {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		res, _ := ioutil.ReadAll(r.Body)

		var formData RecordForm
		err := json.Unmarshal(res, &formData)
		utils.CheckErr(err)

		resJSON, success := addCounselingRecord(formData, common.GetUserID(r))
		if success {
			fmt.Fprintln(w, string(resJSON))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
