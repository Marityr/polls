package shema

type (
	User struct {
		Id       int    `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	UserData struct {
		Id      int  `json:"id,omitempty"`
		User    User `json:"username,omitempty"`
		Count   int  `json:"count,omitempty"`
		Payment bool `json:"payment,omitempty"`
	}
)
