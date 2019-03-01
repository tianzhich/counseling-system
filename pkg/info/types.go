package info

// City struct
type City struct {
	CityID   int    `json:"cityId"`
	CityName string `json:"cityName"`
}

// PreInfo xxx
type PreInfo struct {
	Cities []City `json:"cities"`
}
