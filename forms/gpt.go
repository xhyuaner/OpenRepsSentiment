package forms

// UserQuestionForm 用户询问的问题文本---请求参数结构体
type UserQuestionForm struct {
	UserQuestion string `form:"user_question" json:"user_question" binding:"required"`
}
