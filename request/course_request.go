package request

import "github.com/gin-gonic/gin"

type GetAllCourse struct {
	Language string `form:"language"` // 语言（必填），空表示中文，"en" 表示英文
}

// BinGetAllCourse 从查询参数绑定 GetAllCourse 请求结构体
// @Param   course_status  query    int     true  "课程状态（必填），1 是进行中，2 是即将到来"
// @Param   language       query    string  true  "语言（必填），空表示中文，en 表示英文"
func BinGetAllCourse(c *gin.Context) (*GetAllCourse, error) {
	var req GetAllCourse
	if err := c.ShouldBindQuery(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// GetCourseInfo 用于获取单个课程的信息，包含课程的唯一标识符
type GetCourseInfo struct {
	CourseID string `form:"course_id" binding:"required"` // 课程ID（必填），用于获取指定的课程
}

func BinGetCourseInfo(c *gin.Context) (*GetCourseInfo, error) {
	var req GetCourseInfo
	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

type GetCourseByPath struct {
	Path string `uri:"path" binding:"required"` // 使用 uri 标签绑定路径参数
}

func BinGetCourseByPath(c *gin.Context) (*GetCourseByPath, error) {
	var req GetCourseByPath
	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

type GetCourseChaptersByPath struct {
	Path string `uri:"path" binding:"required"` // 使用 uri 标签绑定路径参数
}

func BinGetCourseChaptersByPath(c *gin.Context) (*GetCourseChaptersByPath, error) {
	var req GetCourseChaptersByPath
	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

type GetCourseQuizzes struct {
	CourseID string `form:"course_id" binding:"required"` // 课程ID（必填），用于获取指定的课程
	Lan      string `form:"lan"`                          // 语言（可选），用于指定语言版本
}

// BinGetCourseQuizzes
func BinGetCourseQuizzes(c *gin.Context) (*GetCourseQuizzes, error) {
	var req GetCourseQuizzes
	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}
	req.Lan = c.Query("lan")
	return &req, nil
}

type GetUserCourseLesson struct {
	CourseID string `form:"course_id" binding:"required"` // 课程ID（必填），用于获取指定的课程
	LessonID string `form:"lesson_id" binding:"required"` // 单元ID（必填），用于获取指定课程的单元
	Lan      string `form:"lan"`                          // 语言（可选），用于指定语言版本
}

// BinGetUserCourseLesson
func BinGetUserCourseLesson(c *gin.Context) (*GetUserCourseLesson, error) {
	var req GetUserCourseLesson
	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}
	req.Lan = c.Query("lan")
	return &req, nil
}
