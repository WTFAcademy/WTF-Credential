package response

type QuizContent struct {
	Title   string  `json:"title"`   // 标题
	Meta    Meta    `json:"meta"`    // 元数据
	Content Content `json:"content"` // 内容
	Id      int64   `json:"id"`      // ID
}
type Meta struct {
	Type   string   `json:"type"`  // 类型
	Score  int      `json:"score"` // 分数
	Answer []string `json:"answer"`
	ID     int64    `json:"id"` // ID
}
type Content struct {
	Extends []Extend `json:"extend"`  // 扩展
	Options []Option `json:"options"` // 选项
}
type Option struct {
	Label string `json:"label"` // 标签
	Value string `json:"value"` // 值
}
type Extend struct {
	Type string `json:"type"` // 类型
	Raw  string `json:"raw"`  // 原始数据
}

type GetChapterQuizzes struct {
	GetSimpleChapterWithCourseResp
	ExerciseList []QuizContent `json:"exercise_list"` // 练习列表
}

type QuizGradeSubmitResponse struct {
	Score    int `json:"score"`
	ErrorCnt int `json:"error_cnt"`
}
