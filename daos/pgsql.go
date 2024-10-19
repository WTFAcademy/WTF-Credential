package daos

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"time"
	"wtf-credential/configs"
)

// db 全局MySQL数据库操作对象
var DB *gorm.DB

// InitPostgres 链接 PostgreSQL 数据库
func InitPostgres() {
	cfg := configs.Config().Postgres // 使用 Postgres 配置
	if cfg.Host == "" {
		panic("invalid postgres host")
	}
	// PostgreSQL 数据库的连接字符串格式
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port, cfg.SslMode)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to postgres database")
	}
	fmt.Println("Successfully connected to PostgreSQL database!")
}

// Logger is a custom logger for GORM that can be used to listen for slow queries.
type Logger struct {
	Writer io.Writer
}

// LogMode sets the logging mode for the custom logger.
func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info logs general information messages.
func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	// Implement your custom logging for Info messages here
}

// Warn logs warning messages.
func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	// Implement your custom logging for Warning messages here
}

// Error logs error messages.
func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	// Implement your custom logging for Error messages here
}

// Trace logs SQL queries.
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// Check if the query execution time exceeds a threshold (e.g., 100 milliseconds)
	threshold := 100 * time.Millisecond
	if elapsed := time.Since(begin); elapsed > threshold {
		query, rows := fc()
		// Implement your custom handling for slow queries here
		// You can log or take any other action as needed
		log.Printf("Slow query: %s [%v] %s\n", elapsed, rows, query)
	}
}
