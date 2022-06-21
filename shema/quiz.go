package shema

// type BlockQuizInt interface {
// 	Create(DB *gorm.DB) // ned return
// 	Save(DB *gorm.DB)
// }

// func (self BlockQuiz) Create(DB *gorm.DB) {

// }
// func (self BlockQuiz) Save(DB *gorm.DB) {

// }

type (
	// Блок опроса
	BlockQuiz struct {
		Id    int    `json:"id,omitempty" gorm:"primary_key"`
		Title string `json:"title,omitempty"`
		Quiz  []Quiz `json:"quiz,omitempty" gorm:"many2many:block_quiz;"`
	}

	// Опрос
	Quiz struct {
		Id        int         `json:"id,omitempty" gorm:"primary_key"`
		BlockQuiz []BlockQuiz `json:"blockquiz,omitempty" gorm:"many2many:block_quiz;"`
		Topic     string      `json:"topic,omitempty"`
		Legend    string      `json:"legend,omitempty"`
		Questions []Questions `json:"questions,omitempty" gorm:"many2many:questions_quiz;"`
	}

	// Вопросы
	Questions struct {
		Id     int    `json:"id,omitempty" gorm:"primary_key"`
		Title  string `json:"title,omitempty"`
		Quiz   []Quiz `json:"quiz,omitempty" gorm:"many2many:questions_quiz;"`
		Legend string `json:"legend,omitempty"`
	}

	// Варианты ответов
	Answers struct {
		Id        int       `json:"id,omitempty" gorm:"primary_key"`
		Questions Questions `json:"questions,omitempty" gorm:"foreignkey:Questions"`
		Option    string    `json:"option,omitempty"`
		Value     int       `json:"value,omitempty"`
		Cause     []Cause   `json:"cause,omitempty" gorm:"many2many:causes_block"`
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
		Id        int       `json:"id,omitempty" gorm:"primary_key"`
		Users     Users     `json:"user_id,omitempty" gorm:"foreignkey:Users"`
		Cause     Cause     `json:"cause,omitempty" gorm:"foreignkey:Cause"`
		Questions Questions `json:"questions,omitempty" gorm:"foreignkey:Questions"`
		Count     int       `json:"count,omitempty"`
		Result    int       `json:"result,omitempty"`
		Color     string    `json:"color,omitempty"`
	}
)
