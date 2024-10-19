package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	config  GlobalConfig
	rConfig sync.RWMutex
)

// PostgresConfig PostgreSQL 配置参数
type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
	SslMode  string // PostgreSQL 通常需要指定 sslmode
}

// OAuthConfig 用于存储各个 OAuth 配置信息
type OAuthConfig struct {
	Github OAuthProviderConfig
}

// OAuthProviderConfig 具体的 OAuth 提供者的配置
type OAuthProviderConfig struct {
	ClientID     string
	ClientSecret string
}

// GlobalConfig 全局配置
type GlobalConfig struct {
	Port        string
	Postgres    PostgresConfig // 使用 PostgresConfig 代替 MysqlConfig
	Redis       RedisConfig
	Debug       bool
	JwtSecret   string // 添加 JWT 密钥字段
	OAuth       OAuthConfig
	GithubToken string   // 添加 GitHub Token 字段
	Org         string   // 添加组织字段
	BypassRepos []string `json:"bypass_repos"` // 新增字段，存储 bypass_repos
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// Config 返回配置文件
func Config() GlobalConfig {
	rConfig.RLock()
	configCopy := config
	rConfig.RUnlock()
	return configCopy
}

// 加载配置文件
func ParseConfig(cfg string) {
	viper.SetConfigFile(cfg)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件读取错误")
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}
