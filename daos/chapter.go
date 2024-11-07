package daos

import (
	"context"
	"gorm.io/gorm"
	model "wtf-credential/models"
)

// FetchChapterDetailsByID 根据章节 ID 获取章节详情
func FetchChapterDetailsByID(ctx context.Context, chapterID int64) (*model.Chapter, error) {
	var chapter model.Chapter
	err := DB.WithContext(ctx).Where("id = ?", chapterID).First(&chapter).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 若未找到记录，可以在此处返回 nil 和自定义错误，表示章节不存在
			return nil, nil
		}
		// 其他数据库错误
		return nil, err
	}
	return &chapter, nil
}

func GetChapterByPathAndRoutePath(ctx context.Context, path, routePath string) (*model.Chapter, error) {
	var chapter model.Chapter
	if err := DB.Where("path = ? AND route_path = ?", path, routePath).First(&chapter).Error; err != nil {
		return nil, err
	}
	return &chapter, nil
}
