package daos

import (
	"context"
	"github.com/google/uuid"
	model "wtf-credential/models"
)

// GetUserQuizGradeByUserIdCourseId 根据用户ID和课程ID获取用户的测验成绩
func GetUserQuizGradeByUserIdCourseId(ctx context.Context, userId string, courseId string) ([]*model.TbLessonQuizGrade, error) {
	var grades []*model.TbLessonQuizGrade

	// 将 userId 和 courseId 转换为 UUID
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err // 如果转换失败，返回错误信息
	}
	courseUUID, err := uuid.Parse(courseId)
	if err != nil {
		return nil, err // 如果转换失败，返回错误信息
	}

	// 查询数据库获取用户测验成绩
	err = DB.WithContext(ctx).
		Where("uid = ? AND course_id = ?", userUUID, courseUUID).
		Find(&grades).Error

	if err != nil {
		return nil, err // 返回错误信息
	}

	return grades, nil // 返回用户测验成绩
}
