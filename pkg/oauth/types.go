package oauth

// SignupForm to be posted
type SignupForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// SigninForm to be exported
type SigninForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConsultantDetail xxx
type ConsultantDetail struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ApplyForm xxx
type ApplyForm struct {
	Name        string             `json:"name"`
	Gender      int                `json:"gender"`
	WorkYears   int                `json:"workYears"`
	Description string             `json:"description"`
	Motto       string             `json:"motto"`
	Detail      []ConsultantDetail `json:"detail"`
	AudioPrice  int                `json:"audioPrice"`
	VideoPrice  int                `json:"videoPrice"`
	FtfPrice    int                `json:"ftfPrice"`
	City        string             `json:"city"`
}
