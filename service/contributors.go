package service

import (
	"context"
	"encoding/json"
	"fmt"
	"wtf-credential/configs"
	"wtf-credential/request"
)

// Contributor 定义 GitHub 贡献者的结构体
type Contributor struct {
	Login         string `json:"login"`
	AvatarURL     string `json:"avatar_url"`
	Contributions int    `json:"contributions"`
}

// ContributorInfo 定义贡献者信息的结构体
type ContributorInfo struct {
	AvatarUrl     string `json:"avatar_url"`
	Id            int    `json:"id"`
	Login         string `json:"login"`
	Type          string `json:"type"`          // 如果需要，可以添加这个字段
	Contributions int    `json:"contributions"` // 添加贡献次数字段
}

// GetContributorsList 从 Redis 获取所有贡献者信息并返回
func GetContributorsList(ctx context.Context, req request.GetContributorsList) (map[string][]ContributorInfo, error) {
	var keys []string
	var err error

	// 如果 req.Repo 为空，获取所有键；否则获取特定 Repo 键
	if req.Repo == "" {
		keys, err = configs.Rdb.Keys(ctx, "*").Result() // 获取所有 Redis 中的键
	} else {
		keys = []string{req.Repo} // 只获取指定的 Repo 键
	}

	if err != nil {
		return nil, err
	}

	// 使用 map 存储按键分类的贡献者信息
	contributorsMap := make(map[string][]ContributorInfo)

	for _, key := range keys {
		// 跳过 specific_contributors_stars 和 contributors_stars 键
		if key == "specific_contributors_stars" || key == "contributors_stars" {
			continue //TODO:临时解决方案 后面要回来优化
		}
		// 从 Redis 获取每个键对应的值
		value, err := configs.Rdb.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		// 打印获取的值以进行调试
		fmt.Printf("Key: %s, Value: %s\n", key, value)

		// 解析 JSON 字符串到 ContributorInfo 结构体
		var contributors []ContributorInfo
		if err := json.Unmarshal([]byte(value), &contributors); err != nil {
			return nil, err
		}

		// 将解析的贡献者添加到对应键的数组中
		contributorsMap[key] = contributors
	}

	// 返回按键分类的贡献者信息
	return contributorsMap, nil
}
