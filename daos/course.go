package daos

import (
	"context"
	model "wtf-credential/models"
)

// 根据 startStatus 获取课程的 DAO 方法
func GetCoursesByStartStatus(ctx context.Context, startStatus int) ([]model.TbCourse, error) {
	var res []model.TbCourse
	var err error

	if startStatus > 0 {
		err = DB.WithContext(ctx).
			Where("start_status = ?", startStatus).
			Order("sort ASC"). // 使用引号确保 SQL 语法正确
			Find(&res).Error
	} else {
		err = DB.WithContext(ctx).
			Order("sort ASC"). // 使用引号确保 SQL 语法正确
			Find(&res).Error
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetCourseInfoByCourseId(courseId string) (courseInfo *model.TbCourse, err error) {
	var course = model.TbCourse{}
	if err := DB.Where("id = ?", courseId).First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}
