package model

// Todo 用户待办事项  {TodoID: 1, text: 'Stay hungry，Stay foolish.', done: false }
type Todo struct {
	Model

	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID uint   `json:"userID"`
}
