package daos

import (
	"context"
	"errors"
	"fmt"
	model "wtf-credential/models"

	"github.com/google/uuid"
)

// 将 int 转换为 UUID 的函数
func intToUUID(id int) (uuid.UUID, error) {
	// 使用 id 的低位 8 字节生成 UUID
	return uuid.NewSHA1(uuid.Nil, []byte(fmt.Sprintf("%d", id))), nil
}

// GetLessonQuizByCourseId 根据课程 ID 获取对应的测验信息
func GetLessonQuizByCourseId(ctx context.Context, courseId string) (map[uuid.UUID]model.TbLessonQuiz, error) {
	var quizzes []model.TbLessonQuiz

	// 查询数据库获取测验信息
	err := DB.WithContext(ctx).Where("course_id = ?", courseId).Find(&quizzes).Error
	if err != nil {
		return nil, err // 返回错误信息
	}

	// 如果没有找到测验，返回错误
	if len(quizzes) == 0 {
		return nil, errors.New("no quizzes found for this course")
	}

	// 将测验信息转换为以 UUID 为键的映射
	quizMap := make(map[uuid.UUID]model.TbLessonQuiz)
	for _, quiz := range quizzes {
		quizID, err := intToUUID(quiz.Id) // 将 int 转换为 UUID
		if err != nil {
			return nil, err // 返回错误信息
		}
		quizMap[quizID] = quiz // 使用转换后的 UUID 作为键
	}

	return quizMap, nil // 返回测验信息映射
}

// FindQuizListByChapterId 根据章节 ID 获取对应的测验列表 sort 都是0 这里按照 id 正序
func FindQuizListByChapterId(ctx context.Context, chapterId int64) ([]model.Quiz, error) {
	var quizzes []model.Quiz
	err := DB.WithContext(ctx).Where("chapter_id = ?", chapterId).Order("id asc").Find(&quizzes).Error
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}
