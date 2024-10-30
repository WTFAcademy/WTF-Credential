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
