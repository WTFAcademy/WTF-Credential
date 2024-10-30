package daos

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	model "wtf-credential/models"
	"wtf-credential/response"
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

// GetUserProfileByID retrieves user profile information from the database by userId.
func GetUserProfileByID(ctx context.Context, userId string) (*response.GetProfileByUserID, error) {
	var user model.TbUserCore

	// 检查 userId 是否为有效的 UUID 格式
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// 根据用户 ID 查询 tb_user_core 表中的用户信息
	if err := DB.WithContext(ctx).Where("id = ?", uid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// 构造返回结构体 GetProfileByUserID
	profile := &response.GetProfileByUserID{
		Github:   user.Github,   // 用户的 GitHub 用户名
		Email:    user.Email,    // 用户的电子邮件地址
		Username: user.UserName, // 用户名
		Nickname: user.NickName, // 如果没有对应字段，可以从其他表获取或设置为空
		Twitter:  user.Twitter,  // 用户的 Twitter 账号
		Bio:      user.Bio,      // 用户简介
		Viewer:   user.Viewer,   // 查看者与用户的关系状态
		Avatar:   user.Avatar,   // 用户头像 URL
		Wallet:   "",            // 钱包地址可从其他表（例如钱包表）获取
	}

	return profile, nil
}

func GetUserCount() (int64, error) {
	var count int64
	if err := DB.Table(model.TbUserCore{}.TableName()).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
