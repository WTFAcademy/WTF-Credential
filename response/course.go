package response

import (
	"github.com/google/uuid"
	model "wtf-credential/models"
)

// LessonQuizInfo 表示课程单元的测验信息
type LessonQuizInfo struct {
	Id            uuid.UUID `json:"id"`             // 课程单元的唯一标识符
	RoutePath     string    `json:"route_path"`     // 课程单元的路由路径
	EnTitle       string    `json:"en_title"`       // 课程单元的英文标题
	CourseId      uuid.UUID `json:"course_id"`      // 关联的课程 ID
	EstimatedTime string    `json:"estimated_time"` // 预估完成时间
	Sort          int       `json:"sort"`           // 课程单元排序
	Title         string    `json:"title"`          // 课程单元的标题
	QuizId        int       `json:"quiz_id"`        // 关联的测验 ID
	IsFinish      bool      `json:"is_finish"`      // 是否完成测验
	ScorePercent  int64     `json:"score_percent"`  // 用户得分百分比
}

// CourseQuizList 表示课程测验信息的列表和毕业状态
type CourseQuizList struct {
	List        []*LessonQuizInfo `json:"list"`         // 存储课程测验信息的切片
	CanGraduate bool              `json:"can_graduate"` // 是否可以毕业
	Course      *model.TbCourse   `json:"course"`       // 关联的课程信息
}

type CourseDetail struct {
	Title       string `json:"title"`       // 课程标题
	Path        string `json:"path"`        // 课程路由路径
	Description string `json:"description"` // 课程描述
	CoverImg    string `json:"cover_img"`   // 课程封面图片链接
	Sort        int64  `json:"sort"`        // 课程排序编号
	TotalScore  int    `json:"total_score"` // 课程总分
	UserCnt     int    `json:"user_cnt"`    // 参与人数
	ShareUrl    string `json:"share_url"`   // 课程分享链接
	Category    string `json:"category"`    // 课程分类
	PassCount   int    `json:"pass_count"`  // 通过总数
}

type CoursesResponse struct {
	Published   []CourseDetail `json:"published"`   // 已发布课程列表
	Unpublished []CourseDetail `json:"unpublished"` // 未发布课程列表
}

type GetCoursesByType struct {
	CoursesByType map[string][]CourseDetail `json:"courses_by_type"`
}

type GetStatistics struct {
	LearnerCount     int64 `json:"learner_count"`     // 学习人数
	ContributorCount int   `json:"contributor_count"` // 贡献人数
	StarCount        int   `json:"star_count"`        // Stars数量
	BonusAmount      int   `json:"bonus_amount"`      // 奖金
}

type GetCourseInfoByPath struct {
	Title           string `json:"title"`            // 课程标题
	Path            string `json:"path"`             // 课程路由路径
	Description     string `json:"description"`      // 课程描述
	CoverImg        string `json:"cover_img"`        // 课程封面图片链接
	Category        string `json:"category"`         // 课程分类
	Level           string `json:"level"`            // 难易程度
	Repo            string `json:"repo"`             // 仓库地址
	CurrentLearners int    `json:"current_learners"` // 当前课程学习者的数量
	StudyTime       int    `json:"study_time"`       // 学习时间（单位：分钟）
	LastUpdated     string `json:"last_updated"`     // 课程更新时间（格式：ISO 8601，例如 "2024-11-03T14:00:00Z"）
}

// GetCourseChapters 表示课程章节的信息
type GetCourseChapters struct {
	Id              int64   `json:"id"`
	Title           string  `json:"title,omitempty"`         // 章节标题
	RoutePath       string  `json:"route_path,omitempty"`    // 章节路由路径
	Sort            int64   `json:"sort,omitempty"`          // 排序编号
	CurrentLearners int     `json:"current_learners"`        // 当前课程学习者的数量
	QuizProgress    float64 `json:"quiz_progress,omitempty"` // 章节中测验的完成进度（以百分比表示即0-1小数）
	CodeProgress    float64 `json:"code_progress,omitempty"` // 章节中代码练习的完成进度（以百分比表示0-1小数）
}

// GetChapterDetailsByID 结构体，用于存储章节详情
type GetChapterDetailsByID struct {
	Title     string `json:"title"`      // 标题
	Sort      int64  `json:"sort"`       // 排序
	Content   string `json:"content"`    // 文档内容
	StudyTime int64  `json:"study_time"` // 章节学习时间，单位为分钟
	Score     int    `json:"score"`      // 分数，若已登录为用户最高分数，未登录为 0
}
