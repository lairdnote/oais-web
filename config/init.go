package config

import (
	"github.com/joho/godotenv"
)

func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	/*
		// 连接数据库
		model.Database(os.Getenv("POSTGRES_DSN"))
		cache.Redis()
	 */
}