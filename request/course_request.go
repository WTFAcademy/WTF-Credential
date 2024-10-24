package request

import "github.com/gin-gonic/gin"

type GetAllCourse struct {
	CourseStatus int    `form:"course_status" binding:"required"` // 课程状态（必填），1 是进行中，2 是即将到来
	Language     string `form:"language" binding:"required"`      // 语言（必填），空表示中文，"en" 表示英文
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
