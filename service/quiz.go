package service

import (
	"context"
	"fmt"
	"wtf-credential/daos"
	model "wtf-credential/models"
	"wtf-credential/request"
	"wtf-credential/response"
	"wtf-credential/util"

	"github.com/samber/lo"

	jsoniter "github.com/json-iterator/go"
)

// QuizContentParser 解析练习内容
type QuizContentParser struct {
	quiz model.Quiz
}

func NewQuizContentParser(quiz model.Quiz) *QuizContentParser {
	return &QuizContentParser{quiz: quiz}
}

func (p *QuizContentParser) Parse() (response.QuizContent, error) {
	var content response.QuizContent
	if err := jsoniter.UnmarshalFromString(p.quiz.Content, &content); err != nil {
		return content, fmt.Errorf("parse quiz content failed: %w", err)
	}

	content.Id = p.quiz.ID
	content.Meta.ID = p.quiz.ID
	return content, nil
}

type quizService struct{}

var QuizService = &quizService{}

func (s *quizService) GetChapterQuizzes(ctx context.Context, req *request.GetChapterQuizzes) (*response.GetChapterQuizzes, error) {
	// 获取章节课程基本信息
	chapter, err := GetSimpleChapterWithCourse(ctx, req.Path, req.RoutePath)
	if err != nil {
		return nil, err
	}
	exerciseList, err := s.GetQuizzesByChapterId(ctx, chapter.SimpleChapter.Id, false)
	if err != nil {
		return nil, err
	}
	return &response.GetChapterQuizzes{
		GetSimpleChapterWithCourseResp: *chapter,
		ExerciseList:                   exerciseList,
	}, nil
}

// GetQuizzesByChapterId 根据章节ID获取练习列表
func (s *quizService) GetQuizzesByChapterId(ctx context.Context, chapterId int64, withAnswer bool) ([]response.QuizContent, error) {
	// 获取练习列表
	quizzes, err := daos.FindQuizListByChapterId(ctx, chapterId)
	if err != nil {
		return nil, fmt.Errorf("find quiz list failed: %w", err)
	}

	quizContents := make([]response.QuizContent, 0, len(quizzes))

	// 解析每个练习内容
	for _, quiz := range quizzes {
		parser := NewQuizContentParser(quiz)
		content, err := parser.Parse()
		if err != nil {
			return nil, err
		}
		// 根据需要隐藏答案
		if !withAnswer {
			content.Meta.Answer = nil
		}
		quizContents = append(quizContents, content)
	}

	return quizContents, nil
}

func (s *quizService) GradeSubmit(ctx context.Context, req *request.QuizGradeSubmitReq) (*response.QuizGradeSubmitResponse, error) {
	// 获取章节课程基本信息
	quizList, err := s.GetQuizzesByChapterId(ctx, req.ChapterId, true)
	if err != nil {
		return nil, err
	}
	// 计算成绩 错误数量，总进度
	gradeResult, err := s.calculateGrade(req, quizList)
	if err != nil {
		return nil, err
	}
	fmt.Println("gradeResult: ", gradeResult)
	// 记录成绩到表 TODO:
	// gradeId, err := daos.AddGrade(ctx, req.CourseId, req.ChapterId, grade, errorCnt, totalScore)
	// if err != nil {
	// 	return nil, err
	// }
	return &response.QuizGradeSubmitResponse{
		Score:    gradeResult.Score,
		ErrorCnt: gradeResult.ErrorCnt,
	}, nil
}

type GradeResult struct {
	Score    int // 得分
	ErrorCnt int // 错误数量
	Progress int // 进度
}

// validateAnswer 验证单个答案是否正确
func (s *quizService) validateAnswer(expected []string, actual []string, score int) (int, bool) {
	if util.EqualSliceIgnoreOrder(expected, actual) {
		return score, true
	}
	return 0, false
}

// calculateGrade 计算成绩
func (s *quizService) calculateGrade(req *request.QuizGradeSubmitReq, exercises []response.QuizContent) (*GradeResult, error) {
	// 将提交的答案转换为map便于查找
	answerMap := lo.KeyBy(req.Answers, func(ans request.AnswerRequest) int64 {
		return ans.Id
	})

	result := &GradeResult{}
	totalPossibleScore := 0

	// 遍历所有习题进行评分
	for _, exercise := range exercises {
		totalPossibleScore += exercise.Meta.Score
		// 检查是否有提交该题的答案
		submission, attempted := answerMap[exercise.Id]
		if !attempted {
			result.ErrorCnt++
			continue
		}
		score, correct := s.validateAnswer(exercise.Meta.Answer, submission.Answers, exercise.Meta.Score)
		result.Score += score
		if !correct {
			result.ErrorCnt++
		}
	}

	// 避免除以零
	if totalPossibleScore > 0 {
		result.Progress = int(float64(result.Score) / float64(totalPossibleScore) * 100)
	}

	return result, nil
}
