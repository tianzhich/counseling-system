package operation

import (
	"counseling-system/pkg/common"
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func addCounselingRecord(formData common.RecordForm, uid int) (string, bool) {
	var insertStr = "insert counseling_record set c_id=?, u_id=?, method=?, times=?, name=?, age=?, gender=?, phone=?, contact_phone=?, contact_name=?, contact_rel=?, `desc`=?, status=?"
	var resp common.Response

	// method含双引号，插入数据库前进行转义处理
	methodStr := strings.Replace(formData.Method, "\"", "\\\"", -1)
	// method name处理
	methodReg := regexp.MustCompile(`name":"(.*?)"`)
	params := methodReg.FindStringSubmatch(formData.Method)
	var methodName = params[1]

	if _, success := utils.InsertDB(insertStr, formData.CID, uid, methodStr, formData.Times, formData.Name, formData.Age, formData.Gender, formData.Phone, formData.ContactPhone, formData.ContactName, formData.ContactRel, formData.Desc, "wait_contact"); success {
		// 增加通知
		var title = fmt.Sprintf("%v向您发起了咨询预约(%v)，请及时确认", formData.Name, methodName)
		var no = common.Notification{Title: title, Desc: ""}
		common.AddNotification(common.GetUserIDByCID(*(formData.CID)), no)

		resp.Code = 1
		resp.Message = "ok"
		resJSON, _ := json.Marshal(resp)
		return string(resJSON), true
	}
	fmt.Println("新增咨询记录失败，数据库插入错误")
	return "", false
}

// 将通知或私信标记为已读
func markRead(ids []string, t string) bool {
	var updateStr string
	if t == "notification" {
		updateStr = "update notification set is_read=1 where "
	} else if t == "message" {
		updateStr = "update message set is_read=1 where "
	} else {
		fmt.Println("标注通知已读，非法类型")
		return false
	}

	for _, id := range ids {
		iid, _ := strconv.Atoi(id)
		updateStr += fmt.Sprintf("id=%v || ", iid)
	}

	updateStr = strings.TrimSuffix(updateStr, " || ")
	if ok := utils.UpdateDB(updateStr); !ok {
		fmt.Println("标注通知或私信已读，数据库更新失败")
		return false
	}
	return true
}

func addMessage(uid int, m common.Message) bool {
	var insertStr = "insert into message set sender=?, receiver=?, detail=?"
	if _, ok := utils.InsertDB(insertStr, uid, m.Receiver, m.Detail); !ok {
		fmt.Println("添加私信，插入数据库错误！")
		return false
	}
	return true
}
