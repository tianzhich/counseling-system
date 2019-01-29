package feature1

// Person defines to be exported
type Person struct {
	Name   string
	Gender string
	Age    int
}

// Res defines the fake get request type
type Res struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}
