package service

import (
	"context"
	"github.com/json-iterator/go"
	"wtf-credential/daos"
	model "wtf-credential/models"
	"wtf-credential/request"
)

// CourseExtendedInfo 包含课程的扩展信息，
// 包括课程标题和描述的不同语言版本。
type CourseExtendedInfo struct {
	EnTitle string `json:"en_title"` // EnTitle 表示课程的英文标题
	EnDesc  string `json:"en_desc"`  // EnDesc 表示课程的英文描述。
}

// formatCourseList 格式化课程信息
func formatCourseList(lan string, courseData model.TbCourse) model.TbCourse {
	var courseExtendedInfo CourseExtendedInfo
	err := jsoniter.UnmarshalFromString(courseData.ExtendedInfo, &courseExtendedInfo)
	if err == nil && lan == "en" {
		courseData.Title = courseExtendedInfo.EnTitle
		courseData.Description = courseExtendedInfo.EnDesc
	}
	courseData.ExtendedInfo = ""
	return courseData
}

func GetAllCourse(ctx context.Context, req *request.GetAllCourse) ([]model.TbCourse, error) {

	courses, err := daos.GetCoursesByStartStatus(ctx, req.CourseStatus)
	if err != nil {
		return courses, err
	}
	// 格式化课程列表
	for i, v := range courses {
		courses[i] = formatCourseList(req.Language, v)
	}

	return courses, nil
}

func GetCourseInfo(ctx context.Context, req *request.GetCourseInfo) (*model.TbCourse, error) {

	course, err := daos.GetCourseInfoByCourseId(req.CourseID)
	if err != nil {
		return nil, err
	}
	return course, nil
}
