package shema

type (
	User struct {
		Id       int    `json:"id,omitempty" gorm:"primary_key"`
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
		// uid
	}

	UserData struct {
		Id      int  `json:"id,omitempty" gorm:"primary_key"`
		User    User `json:"username,omitempty"`
		Count   int  `json:"count,omitempty"`
		Payment bool `json:"payment,omitempty"`
	}
)
