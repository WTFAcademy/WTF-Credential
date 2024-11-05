package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
	"wtf-credential/configs"
	"wtf-credential/daos"
	model "wtf-credential/models"
	"wtf-credential/request"
	"wtf-credential/response"
)

const (
	BonusAmount = 1000000 // 定义一个固定的奖金常量
)

// CourseExtendedInfo 包含课程的扩展信息，
// 包括课程标题和描述的不同语言版本。
type CourseExtendedInfo struct {
	EnTitle string `json:"en_title"` // EnTitle 表示课程的英文标题
	EnDesc  string `json:"en_desc"`  // EnDesc 表示课程的英文描述。
}

func GetAllCourse(ctx context.Context) (*response.CoursesResponse, error) {
	// 获取全部课程
	allCourses, err := daos.GetAllCourses(ctx)
	if err != nil {
		return nil, err
	}

	// 创建响应对象
	response := &response.CoursesResponse{
		Published: []response.CourseDetail{ // 直接在这里添加模拟数据
			{
				Title:       "Mock Course 1",
				Path:        "mock-course-1",
				Description: "This is a mock course for testing.",
				CoverImg:    "https://example.com/mock1.jpg",
				Sort:        1,
				TotalScore:  100,
				UserCnt:     50,
				ShareUrl:    "https://example.com/mock1",
				Category:    "Mock Category",
				PassCount:   45,
			},
			{
				Title:       "Mock Course 2",
				Path:        "mock-course-2",
				Description: "Another mock course for testing.",
				CoverImg:    "https://example.com/mock2.jpg",
				Sort:        2,
				TotalScore:  90,
				UserCnt:     30,
				ShareUrl:    "https://example.com/mock2",
				Category:    "Mock Category",
				PassCount:   25,
			},
		},
		Unpublished: make([]response.CourseDetail, 0, len(allCourses)),
	}

	fillCourseDetails(allCourses, &response.Unpublished, false)

	return response, nil
}

// fillCourseDetails 将课程信息填充到指定的课程详情切片中
func fillCourseDetails(courses []model.Course, courseDetails *[]response.CourseDetail, isPublished bool) {
	for _, course := range courses {
		*courseDetails = append(*courseDetails, response.CourseDetail{
			Title:       course.Title,
			Path:        course.Path,
			Description: course.Description,
			CoverImg:    course.Cover,
			Sort:        course.Sort,
			TotalScore:  0,
			UserCnt:     0,
			ShareUrl:    "",
			Category:    course.Category,
			PassCount:   0,
		})
	}
}

func GetCoursesByType(ctx context.Context) (map[string][]response.CourseDetail, error) {
	// 获取所有课程
	allCourses, err := daos.GetAllCourses(ctx)
	if err != nil {
		return nil, err
	}

	// 按类型分类课程
	coursesByType := make(map[string][]response.CourseDetail)
	for _, course := range allCourses {
		courseType := course.Category

		// 添加课程到对应类型的切片
		coursesByType[courseType] = append(coursesByType[courseType], response.CourseDetail{
			Title:       course.Title,
			Path:        course.Path,
			Description: course.Description,
			CoverImg:    course.Cover,
			Sort:        course.Sort,
			TotalScore:  1,
			UserCnt:     1,
			ShareUrl:    "123213",
		})
	}

	return coursesByType, nil
}

func GetCourseInfo(ctx context.Context, req *request.GetCourseInfo) (*model.TbCourse, error) {
	course, err := daos.GetCourseInfoByCourseId(ctx, req.CourseID)
	if err != nil {
		return nil, err
	}
	return course, nil
}

// 辅助函数：获取路由路径中的最后一段
func getLastPathSegment(routePath string) string {
	return routePath[strings.LastIndex(routePath, "/")+1:]
}

// 辅助函数：判断用户是否完成课程单元
func isLessonFinished(loginUid string, lessonID uuid.UUID, userScoreList map[uuid.UUID]int64, allowedUIDs map[string]bool) bool {
	if userScoreList[lessonID] == 100 || allowedUIDs[loginUid] {
		return true
	}
	return false
}

// getHighestUserScores 接受用户测验成绩的切片，并返回一个映射，其中键为课程单元 ID，值为最高分数
func getHighestUserScores(userQuizGrades []*model.TbLessonQuizGrade) map[uuid.UUID]int64 {
	scoreMap := make(map[uuid.UUID]int64) // 用于存储每个课程单元的最高分数

	for _, grade := range userQuizGrades {
		// 检查当前成绩是否高于映射中已存储的分数
		if currentScore, exists := scoreMap[grade.LessonId]; !exists || grade.Score > currentScore {
			scoreMap[grade.LessonId] = grade.Score // 更新最高分数
		}
	}

	return scoreMap // 返回包含最高分数的映射
}

func GetCourseQuizzes(ctx context.Context, req *request.GetCourseQuizzes, loginUid string) (*response.CourseQuizList, error) {
	var (
		lessonList    []*model.TbCourseLesson          // 存储课程单元的切片
		userQuizGrade []*model.TbLessonQuizGrade       // 存储用户测验成绩的切片
		courseInfo    *model.TbCourse                  // 存储课程信息
		lessonQuiz    map[uuid.UUID]model.TbLessonQuiz //存储课程单元对应测验的映射
	)

	// 使用通道并行获取数据
	errCh := make(chan error, 4)

	// 获取课程单元列表
	go func() {
		var err error
		lessonList, err = daos.GetLessonByCourseId(ctx, req.CourseID)
		errCh <- err
	}()

	// 获取用户测验成绩
	go func() {
		var err error
		userQuizGrade, err = daos.GetUserQuizGradeByUserIdCourseId(ctx, loginUid, req.CourseID)
		errCh <- err
	}()

	// 获取课程信息
	go func() {
		var err error
		courseInfo, err = daos.GetCourseInfoByCourseId(ctx, req.CourseID)
		errCh <- err
	}()

	// 获取课程测验信息
	go func() {
		var err error
		lessonQuiz, err = daos.GetLessonQuizByCourseId(ctx, req.CourseID)
		errCh <- err
	}()

	// 检查所有通道中的错误
	for i := 0; i < 4; i++ {
		if err := <-errCh; err != nil {
			return nil, err // 返回第一个错误
		}
	}

	// 处理用户分数
	userScoreList := getHighestUserScores(userQuizGrade)
	allowedUIDs := map[string]bool{
		"5fd2ddf9-f4d7-4150-a459-a3e291eae68f": true,
		"875dab8c-0c58-4068-a4ac-025dab1e1b94": true,
	}

	var (
		lessonQuizInfoList []*response.LessonQuizInfo // 使用 response 结构体
		canGraduate        = true
	)

	// 格式化测验信息
	for _, lesson := range lessonList {
		isFinish := isLessonFinished(loginUid, lesson.Id, userScoreList, allowedUIDs)
		if !isFinish {
			canGraduate = false
		}

		lessonQuizInfoList = append(lessonQuizInfoList, &response.LessonQuizInfo{
			Id:            lesson.Id,
			RoutePath:     lesson.RoutePath,
			EnTitle:       getLastPathSegment(lesson.RoutePath),
			CourseId:      lesson.CourseId,
			EstimatedTime: lesson.EstimatedTime,
			Sort:          lesson.Sort,
			Title:         lesson.Title,
			QuizId:        lessonQuiz[lesson.Id].Id,
			IsFinish:      isFinish,
			ScorePercent:  userScoreList[lesson.Id],
		})
	}

	// 返回响应结果
	return &response.CourseQuizList{
		List:        lessonQuizInfoList, // 返回格式化后的测验信息列表
		CanGraduate: canGraduate,
		Course:      courseInfo,
	}, nil
}

func GetUserCourseLesson(ctx context.Context, req *request.GetUserCourseLesson, loginUid string) (string, error) {

	return "nil", nil
}

func GetStatistics(ctx context.Context) (*response.GetStatistics, error) {
	// 从 Redis 获取 Star 数量
	totalStars, err := configs.Rdb.Get(ctx, "contributors_stars").Int()
	if err != nil {
		return nil, fmt.Errorf("获取 Star 数量失败: %v", err)
	}

	// 从 Redis 获取贡献人数
	totalContributors, err := configs.Rdb.Get(ctx, "total_contributors").Int()
	if err != nil {
		return nil, fmt.Errorf("获取贡献人数失败: %v", err)
	}

	//从数据库取用户数量    //TODO:计划放到redis里面
	lenarnerCount, err := daos.GetUserCount()
	if err != nil {
		return nil, fmt.Errorf("获取 lenarnerCount 数量失败: %v", err)
	}
	// 创建并返回统计信息
	statistics := &response.GetStatistics{
		LearnerCount:     lenarnerCount,
		ContributorCount: totalContributors,
		StarCount:        totalStars,
		BonusAmount:      BonusAmount,
	}
	return statistics, nil
}

// GetCourseByPath 根据路径获取课程信息
func GetCourseByPath(ctx context.Context, req *request.GetCourseByPath) (*response.GetCourseInfoByPath, error) {
	// 从数据库中获取课程信息
	course, err := daos.GetCourseInfoByPath(ctx, req.Path)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, nil // 没有找到对应的课程
	}

	// 将数据库查询结果映射到 GetCourseInfoByPath 响应结构体
	courseInfo := &response.GetCourseInfoByPath{
		Title:           course.Title,                          // 映射课程标题
		Path:            course.Path,                           // 映射课程的路由路径
		Description:     course.Description,                    // 映射课程的详细描述
		CoverImg:        course.Cover,                          // 映射课程封面图片链接
		Category:        course.Category,                       // 映射课程分类
		Level:           course.Level,                          // 映射课程难易程度
		Repo:            course.Repo,                           // 映射课程代码仓库地址
		CurrentLearners: 1000,                                  // 映射当前学习者的数量
		StudyTime:       20,                                    // 映射当前课程学习时间
		LastUpdated:     course.UpdatedAt.Format(time.RFC3339), // 映射课程更新时间
	}

	return courseInfo, nil
}

// GetCourseChaptersByPath 根据课程路径获取课程章节
func GetCourseChaptersByPath(ctx context.Context, req *request.GetCourseChaptersByPath) ([]response.GetCourseChapters, error) {
	// 查询课程单元
	lessons, err := daos.FetchLessonsByPath(ctx, req.Path)
	if err != nil {
		return nil, err
	}

	// 构建返回的章节信息
	var courseChapters []response.GetCourseChapters // 直接使用 Chapter 类型的切片

	for _, lesson := range lessons {
		courseChapters = append(courseChapters, response.GetCourseChapters{
			Title:           lesson.Title,
			RoutePath:       lesson.RoutePath,
			Sort:            lesson.Sort,
			CurrentLearners: 1000,
			QuizProgress:    0,
			CodeProgress:    0,
		})
	}

	return courseChapters, nil
}

// GetUserCourseChaptersByPath 根据课程路径获取用户的课程章节
func GetUserCourseChaptersByPath(ctx context.Context, req *request.GetCourseChaptersByPath, loginUid string) ([]response.GetCourseChapters, error) {
	// 查询课程单元
	lessons, err := daos.FetchLessonsByPath(ctx, req.Path)
	if err != nil {
		return nil, err
	}

	// 构建返回的章节信息
	var courseChapters []response.GetCourseChapters

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	for _, lesson := range lessons {
		// 随机选择 QuizProgress 和 CodeProgress 的值
		progress := 0.5 + rand.Float64() // 生成 0.5 或 1

		// 如果随机值大于 0.75 则设为 1，否则设为 0.5
		if progress > 0.75 {
			progress = 1
		} else {
			progress = 0.5
		}

		courseChapters = append(courseChapters, response.GetCourseChapters{
			Title:        lesson.Title,
			RoutePath:    lesson.RoutePath,
			Sort:         lesson.Sort,
			QuizProgress: progress, // 随机值
			CodeProgress: progress, // 随机值
		})
	}

	return courseChapters, nil
}
