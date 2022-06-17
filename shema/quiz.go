package shema

type (
	// Блок опроса
	BlockQuiz struct {
		Id    int    `json:"id,omitempty" gorm:"primary_key"`
		Title string `json:"title,omitempty"`
	}

	// Опрос
	Quiz struct {
		Id        int       `json:"id,omitempty" gorm:"primary_key"`
		BlockQuiz BlockQuiz `json:"block_quiz,omitempty" gorm:"foreignkey:BlockQuiz"`
		Topic     string    `json:"topic,omitempty"`
		Legend    string    `json:"legend,omitempty"`
	}

	// Вопросы
	Questions struct {
		Id     int    `json:"id,omitempty" gorm:"primary_key"`
		Title  string `json:"title,omitempty"`
		Quiz   Quiz   `json:"quiz,omitempty" gorm:"foreignkey:Topic"`
		Legend string `json:"legend,omitempty"`
	}

	// Варианты ответов
	Answers struct {
		Id        int       `json:"id,omitempty" gorm:"primary_key"`
		Questions Questions `json:"questions,omitempty" gorm:"foreignkey:Questions"`
		Option    string    `json:"option,omitempty"`
		Value     int       `json:"value,omitempty"`
		Cause     Cause     `json:"cause,omitempty" gorm:"foreignkey:Cause"`
	}

	// Причины болезней
	Cause struct {
		Id          int    `json:"id,omitempty" gorm:"primary_key"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Addition    string `json:"addition,omitempty"`
	}

	// Результаты опроса
	Result struct {
		Id int `json:"id,omitempty" gorm:"primary_key"`
		//UserId    User      `json:"user_id,omitempty" gorm:"foreignkey:Username"`
		Cause     Cause     `json:"cause,omitempty" gorm:"foreignkey:Cause"`
		Questions Questions `json:"questions,omitempty" gorm:"foreignkey:Questions"`
		Count     int       `json:"count,omitempty"`
		Result    int       `json:"result,omitempty"`
		Color     string    `json:"color,omitempty"`
	}
)
