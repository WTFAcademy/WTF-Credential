package daos

import (
	"context"
	model "wtf-credential/models"
)

func GetUserByGithubName(ctx context.Context, githubName *string) (res *model.TbUserCore, err error) {
	var user = model.TbUserCore{}
	if err = DB.WithContext(ctx).Where("github = ?", githubName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, err
}
