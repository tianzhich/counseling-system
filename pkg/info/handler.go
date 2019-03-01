package info

import (
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// PreHandler to get the pre infomation
func PreHandler(w http.ResponseWriter, r *http.Request) {
	var preInfo PreInfo
	preInfo.Cities = getAllCity()

	resJSON, _ := json.Marshal(preInfo)

	fmt.Fprintln(w, string(resJSON))
}

func getAllCity() []City {
	var cities []City

	rows := utils.QueryDB("select id, name from city")
	for i := 0; rows.Next(); i++ {
		var city City

		err := rows.Scan(&city.CityID, &city.CityName)
		utils.CheckErr(err)
		cities = append(cities, city)
	}

	return cities
}
