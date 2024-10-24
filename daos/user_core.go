package daos

import (
	"context"
	"github.com/google/uuid"
	model "wtf-credential/models"
)

// getStringFromPointer safely dereferences a string pointer or returns an empty string if nil
func getStringFromPointer(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func GetUserByGithubName(ctx context.Context, githubName *string) (res *model.TbUserCore, err error) {
	var user = model.TbUserCore{}
	if err = DB.WithContext(ctx).Where("github = ?", getStringFromPointer(githubName)).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, err
}

// CreateUser creates a new user record in the database.
func CreateUser(ctx context.Context, userName, gitHub, email, vieWer *string) error {
	// 生成一个新的 UUID 作为用户 ID
	newUserId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	newUser := model.TbUserCore{
		Id:       newUserId,
		UserName: getStringFromPointer(userName), // 解引用指针，提供默认值
		Bio:      "",                             // 可选项, 初始化为空字符串
		Avatar:   "",                             // 可选项, 初始化为空字符串
		Viewer:   getStringFromPointer(vieWer),   // 解引用指针，提供默认值
		Email:    getStringFromPointer(email),    // 解引用指针，提供默认值
		Github:   getStringFromPointer(gitHub),   // 解引用指针，提供默认值
		Twitter:  "",                             // 可选项, 初始化为空字符串
	}

	// 插入新用户数据到数据库
	err = DB.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		return err
	}

	return nil
}
