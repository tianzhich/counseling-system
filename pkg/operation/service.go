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

	if _, success := utils.InsertDB(insertStr, *(formData.CID), uid, methodStr, formData.Times, formData.Name, formData.Age, formData.Gender, formData.Phone, formData.ContactPhone, formData.ContactName, formData.ContactRel, formData.Desc, "wait_contact"); success {
		// 增加通知
		var title = fmt.Sprintf("%v向您发起了咨询预约(%v)，请及时确认", formData.Name, methodName)
		var no = common.Notification{Title: title, Desc: "", Type: "counseling"}
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

// 增加私信
func addMessage(uid int, m common.Message) bool {
	var insertStr = "insert into message set sender=?, receiver=?, detail=?"
	if _, ok := utils.InsertDB(insertStr, uid, m.Receiver, m.Detail); !ok {
		fmt.Println("添加私信，插入数据库错误！")
		return false
	}
	return true
}

// 咨询流程进度处理, -2表示权限错误, -1表示数据库错误, 0表示处理失败(参数不足,非法操作等), 1表示成功
func appointProcess(uID int, userType int, recordID int, operation int, args processArgs) (int, string) {
	var prevStatus string
	var uuID int
	var cID int
	var updateStr = "update counseling_record set status=?"

	var queryStr = fmt.Sprintf("select c_id, u_id, status from counseling_record where id=%v", recordID)
	rows := utils.QueryDB(queryStr)
	if rows.Next() {
		rows.Scan(&cID, &uuID, &prevStatus)
		if userType == 1 && cID != common.GetCounselorIDByUID(uID) || (userType == 2 && uID != uuID) {
			fmt.Println("咨询流程进度处理，非有效UID")
			return -2, "非有效UID"
		}
	} else {
		fmt.Println("查找咨询记录失败")
		return -1, "查找咨询记录失败"
	}
	rows.Close()

	// 逻辑判断处理
	switch prevStatus {
	case "wait_contact":
		// 咨询师操作
		if userType == 1 {
			if operation == 0 {
				updateStr += fmt.Sprintf(", cancel_reason1=? where id=%v", recordID)
				if args.CancelReason2 != nil {
					utils.UpdateDB(updateStr, "cancel", *(args.CancelReason2))
				} else {
					return 0, "取消理由不能为空"
				}
			} else if operation == 1 {
				updateStr += ", start_time=?"
				if args.StartTime != nil {
					if args.Location != nil {
						updateStr += fmt.Sprintf(", location=? where id=%v", recordID)
						utils.UpdateDB(updateStr, "wait_confirm", *(args.StartTime), *(args.Location))
					} else {
						updateStr += fmt.Sprintf(" where id=%v", recordID)
						utils.UpdateDB(updateStr, "wait_confirm", *(args.StartTime))
					}
				} else {
					return 0, "确认时间不能为空"
				}
			} else {
				return 0, "非法操作，必须为同意(1)或拒绝(0)"
			}
		}

		// 咨询者操作
		if userType == 2 {
			if operation == 1 {
				return 0, "非法操作，咨询者无法主动协商咨询时间"
			} else if operation == 0 {
				if args.CancelReason1 != nil {
					updateStr += fmt.Sprintf(", cancel_reason1=? where id=%v", recordID)
					utils.UpdateDB(updateStr, "cancel", *(args.CancelReason1))
				} else {
					return 0, "取消理由不能为空" // 咨询者理由为选择项
				}
			} else {
				return 0, "非法操作，必须为同意(1)或拒绝(0)"
			}
		}

	case "wait_confirm":
		if userType == 2 {
			if operation == 1 {
				if args.StartTime != nil {
					updateStr += ", start_time=?"
					if args.Location != nil {
						updateStr += fmt.Sprintf(", location=? where id=%v", recordID)
						utils.UpdateDB(updateStr, "wait_counseling", *(args.StartTime), *(args.Location))
					} else {
						updateStr += fmt.Sprintf(" where id=%v", recordID)
						utils.UpdateDB(updateStr, "wait_counseling", *(args.StartTime))
					}
				} else {
					updateStr += fmt.Sprintf(" where id=%v", recordID)
					utils.UpdateDB(updateStr, "wait_counseling")
				}
			} else if operation == 0 {
				if args.CancelReason1 != nil {
					updateStr += fmt.Sprintf(", cancel_reason1=? where id=%v", recordID)
					utils.UpdateDB(updateStr, "cancel", *(args.CancelReason1))
				} else {
					return 0, "取消理由不能为空" // 咨询者理由为选择项
				}
			} else {
				return 0, "非法操作，必须为同意(1)或拒绝(0)"
			}
		} else {
			return -2, "咨询师无法确认协商结果"
		}

	case "wait_counseling":
		if userType == 2 {
			if operation == 1 {
				updateStr += fmt.Sprintf(" where id=%v", recordID)
				utils.UpdateDB(updateStr, "wait_comment")
			} else if operation == 0 {
				if args.CancelReason1 != nil {
					updateStr += fmt.Sprintf(", cancel_reason1=? where id=%v", recordID)
					utils.UpdateDB(updateStr, "wait_comment", args.CancelReason1)
				} else {
					return 0, "取消理由不能为空" // 此时取消不退款
				}
			} else {
				return 0, "非法操作，必须为同意(1)或拒绝(0)"
			}
		} else {
			return -2, "咨询师无法操作"
		}

	case "wait_comment":
		if userType == 2 {
			if operation == 1 {
				updateStr += fmt.Sprintf(", rating_score=?, rating_text=?, letters=? where id=%v", recordID)
				utils.UpdateDB(updateStr, "finish", args.RatingScore, args.RatingText, args.Letters)
			} else {
				return 0, "非法操作，只能确认评价"
			}
		} else {
			return -2, "咨询师无法操作"
		}

	case "finish":
		// 更新感谢信
		if userType == 2 {
			if args.Letters != nil {
				updateStr += fmt.Sprintf(", letters=? where id=%v", recordID)
				utils.UpdateDB(updateStr, "finish", args.Letters)
			} else {
				return 0, "感谢信不能为空"
			}
		} else {
			return -2, "咨询师无法操作"
		}

	default:
		return -1, "咨询状态错误"
	}

	return 1, "ok"
}
