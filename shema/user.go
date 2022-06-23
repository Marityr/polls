package shema

type (
	//Пользователь
	User struct {
		Id       int    `json:"id,omitempty" gorm:"primary_key"`
		Name     string `json:"name,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Group    Group  `json:"group,omitempty" gorm:"foreignkey:Group"`
	}

	// Группа
	Group struct {
		Id          int           `json:"id,omitempty" gorm:"primary_key"`
		Name        string        `json:"name,omitempty"`
		Permissions []Permissions `json:"permission,omitempty" gorm:"many2many:group_permissions"`
	}

	// Доступы
	Permissions struct {
		Id     int      `json:"id,omitempty" gorm:"primary_key"`
		Mark   string   `json:"mark,omitempty"`
		Group  []Group  `json:"group,omitempty" gorm:"many2many:group_permissions"`
		Models []Models `json:"models,omitempty" gorm:"many2many:permissions_models"`
	}

	// Структуры
	Models struct {
		Id          int           `json:"id,omitempty" gorm:"primary_key"`
		Name        string        `json:"name,omitempty"`
		Permissions []Permissions `json:"permission,omitempty" gorm:"many2many:permissions_models"`
	}

	// Данные пользователя
	UserData struct {
		Id      int  `json:"id,omitempty" gorm:"primary_key"`
		Users   User `json:"user,omitempty" gorm:"foreignkey:Users"`
		Count   int  `json:"count,omitempty"`
		Payment bool `json:"payment,omitempty"`
	}

	Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
)
