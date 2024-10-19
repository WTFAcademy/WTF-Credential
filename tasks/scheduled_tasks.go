package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wtf-credential/configs"
)

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"html_url"`
}

// Contributor 定义 GitHub 贡献者的结构体
type Contributor struct {
	Login         string `json:"login"`
	AvatarURL     string `json:"avatar_url"`
	Contributions int    `json:"contributions"`
}

type ContributorInfo struct {
	Login         string `json:"login"`
	Id            int    `json:"id"`
	AvatarUrl     string `json:"avatar_url"`
	Type          string `json:"type"`
	Contributions int    `json:"contributions"`
}

// ContributorArray 定义包含用户名和头像的结构体
type ContributorArray struct {
	Contributors []ContributorInfo `json:"contributors"`
}

// 获取指定组织的所有仓库
func getRepos(org string) []Repo {
	url := fmt.Sprintf("https://api.github.com/orgs/%s/repos?per_page=100", org)
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

// getContributorslist 获取并存储贡献者列表
func getContributorslist() {
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
			key := repo.Name                                                  // 使用仓库名作为主键
			if err := configs.Rdb.Set(ctx, key, value, 0).Err(); err != nil { // 0 表示永不过期
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
	fmt.Println("项目开始运行，立即执行任务。")
	getContributorslist() // 立即运行一次
	ticker := time.NewTicker(48 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			getContributorslist() // 每24小时执行一次
		}
	}
}

//func GetWTFSolidity() {
//	fmt.Println("项目开始运行，立即执行任务。")
//	getWTFSolidity() // 立即运行一次
//	// 每24小时执行一次
//	ticker := time.NewTicker(24 * time.Hour)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			getWTFSolidity() // 每24小时执行一次
//		}
//	}
//}

//// getAA 函数用于获取指定 GitHub 仓库的贡献者列表并存储到 Redis
//func getWTFSolidity() {
//	// Connect endpoint URL
//	endpointUrl := "https://api.github.com/repositories/512209835/contributors?page=1&per_page=1000"
//
//	// Prepare request
//	req, err := http.NewRequest("GET", endpointUrl, nil)
//	if err != nil {
//		fmt.Println("创建请求失败:", err)
//		return
//	}
//
//	// Set request headers
//	req.Header.Set("Accept", "application/vnd.github+json")
//	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
//
//	// Do request
//	res, err := (&http.Client{}).Do(req)
//	if err != nil {
//		fmt.Println("发送请求失败:", err)
//		return
//	}
//	defer res.Body.Close()
//
//	// Check response
//	if res.StatusCode != http.StatusOK {
//		fmt.Printf("获取贡献者失败，状态码: %d\n", res.StatusCode)
//		return
//	}
//
//	// Decode response
//	var contributors []ContributorInfo
//	err = json.NewDecoder(res.Body).Decode(&contributors)
//	if err != nil {
//		fmt.Println("解析响应失败:", err)
//		return
//	}
//
//	// 将贡献者数据存储为数组
//	contributorArray := ContributorArray{Contributors: contributors}
//	value, err := json.Marshal(contributorArray) // 将结构体编码为 JSON 字符串
//	if err != nil {
//		fmt.Println("序列化数据失败:", err)
//		return
//	}
//	// 使用 context.Background() 创建上下文
//	ctx := context.Background()
//
//	// 将贡献者数据存储为数组
//	contributorArray = ContributorArray{Contributors: contributors}
//	value, err = json.Marshal(contributorArray) // 将结构体编码为 JSON 字符串
//	if err != nil {
//		fmt.Println("序列化数据失败:", err)
//		return
//	}
//	// 存储数据到 Redis
//	key := "WTF-Solidity"                                             // 主键
//	if err := configs.Rdb.Set(ctx, key, value, 0).Err(); err != nil { // 0 表示永不过期
//		fmt.Printf("存储到 Redis 失败: %s\n", err)
//		return
//	}
//	fmt.Printf("成功存储到 Redis: %s -> %s\n", key, value)
//}
