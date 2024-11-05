package daos

import (
	"context"
	"github.com/google/uuid"
	model "wtf-credential/models"
)

func GetLessonByCourseId(ctx context.Context, courseId string) ([]*model.TbCourseLesson, error) {
	var lessons []*model.TbCourseLesson

	// 将 courseId 转换为 UUID
	courseUUID, err := uuid.Parse(courseId)
	if err != nil {
		return nil, err // 如果转换失败，返回错误信息
	}
	// 查询数据库获取课程单元信息
	err = DB.WithContext(ctx).Where("course_id = ?", courseUUID).Find(&lessons).Error
	if err != nil {
		return nil, err // 返回错误信息
	}

	return lessons, nil // 返回课程单元信息
}

// FetchLessonsByPath 根据课程路径查询所有课程单元，并按照 Sort 字段排序
func FetchLessonsByPath(ctx context.Context, path string) ([]model.TbCourseLesson, error) {
	var lessons []model.TbCourseLesson

	// 查找所有与给定路径匹配的课程单元，并按照 Sort 字段排序
	err := DB.Where("path = ?", path).Order("sort ASC").Find(&lessons).Error
	if err != nil {
		return nil, err
	}

	return lessons, nil
}
