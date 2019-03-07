package info

import (
	"counseling-system/pkg/utils"
	"fmt"
)

// getDictInfo return the counseling city, method, topic
func getDictInfo(dictType int) []dictInfo {
	var infos []dictInfo
	var queryStr = fmt.Sprintf("select info_code, info_name from dict_info where type_code='%v'", dictType)

	infoRows := utils.QueryDB(queryStr)
	for infoRows.Next() {
		var info dictInfo
		infoRows.Scan(&info.ID, &info.Name)
		infos = append(infos, info)
	}

	return infos
}
