package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wtf-credential/configs"
)

// Repo 定义 GitHub 仓库的结构体
type Repo struct {
	ID   int    `json:"id"`       // 仓库的唯一标识符
	Name string `json:"name"`     // 仓库的名称
	URL  string `json:"html_url"` // 仓库的 GitHub 页面 URL
}

// Contributor 定义 GitHub 贡献者的结构体
type Contributor struct {
	Login         string `json:"login"`         // 贡献者的 GitHub 用户名
	AvatarURL     string `json:"avatar_url"`    // 贡献者的 GitHub 头像 URL
	Contributions int    `json:"contributions"` // 贡献者在该仓库的提交次数
}

// ContributorInfo 定义 GitHub 贡献者的详细信息结构体
type ContributorInfo struct {
	Login         string `json:"login"`         // 贡献者的 GitHub 用户名
	Id            int    `json:"id"`            // 贡献者的唯一 ID
	AvatarUrl     string `json:"avatar_url"`    // 贡献者的 GitHub 头像 URL
	Type          string `json:"type"`          // 用户类型（例如 User、Organization）
	Contributions int    `json:"contributions"` // 贡献者在该仓库的提交次数
}

// ContributorArray 定义包含用户名和头像的结构体
type ContributorArray struct {
	Contributors []ContributorInfo `json:"contributors"`
}

// 获取指定组织的所有仓库
func getRepos(org string) []Repo {
	url := fmt.Sprintf("https://api.github.com/orgs/%s/repos?per_page=1000", org)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("获取组织仓库失败，状态码: %s\n", resp.Status)
		return nil
	}

	var repos []Repo
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		fmt.Println("解析失败:", err)
		return nil
	}

	return repos
}

// 检查字符串切片中是否包含某个字符串
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// getContributorsList 获取并存储贡献者列表
func getContributorsList() {
	org := configs.Config().Org // 指定组织名
	repos := getRepos(org)

	if repos == nil {
		fmt.Println("无法获取仓库列表")
		return
	}

	ctx := context.Background() // 创建上下文

	for _, repo := range repos {
		if contains(configs.Config().BypassRepos, repo.Name) {
			fmt.Printf("跳过仓库: %s\n", repo.Name)
			continue
		}
		fmt.Printf("仓库: %s (ID: %d, URL: %s)\n", repo.Name, repo.ID, repo.URL)
		contributors := getContributors(repo.ID)

		if contributors != nil {
			// 将贡献者列表转换为 JSON 格式
			value, err := json.Marshal(contributors)
			if err != nil {
				fmt.Println("序列化贡献者失败:", err)
				continue
			}
			// 将贡献者数据存储到 Redis，以仓库名为键
			key := repo.Name // 使用仓库名作为主键
			if err := configs.Rdb.Set(ctx, key, value, 0).Err(); err != nil {
				fmt.Printf("存储到 Redis 失败: %s\n", err)
				continue
			}
			fmt.Printf("成功存储到 Redis: %s -> %s\n", key, value)
		}
	}
}

// getContributors 获取指定仓库的贡献者列表
func getContributors(repoID int) []ContributorInfo {
	url := fmt.Sprintf("https://api.github.com/repositories/%d/contributors?per_page=100", repoID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求贡献者失败:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("获取贡献者失败，状态码: %s\n", resp.Status)
		return nil
	}

	var contributors []ContributorInfo
	if err := json.NewDecoder(resp.Body).Decode(&contributors); err != nil {
		fmt.Println("解析贡献者失败:", err)
		return nil
	}

	return contributors
}

func GetContributorsJob() {
	getContributorsList() // 立即运行一次
	ticker := time.NewTicker(48 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			getContributorsList() // 每24小时执行一次
		}
	}
}
