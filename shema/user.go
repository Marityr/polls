package shema

type (
	Users struct {
		Id       int    `json:"id,omitempty" gorm:"primary_key"`
		Name     string `json:"name,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	UserData struct {
		Id      int   `json:"id,omitempty" gorm:"primary_key"`
		User    Users `json:"username,omitempty" gorm:"foreignKey:User" `
		Count   int   `json:"count,omitempty"`
		Payment bool  `json:"payment,omitempty"`
	}

	Group struct {
		Id   int
		Name string
	}

	Permission struct {
	}
)
