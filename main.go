package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	db  *gorm.DB
	rdb *redis.Client
	ctx = context.Background()
)

func main() {
	// 1. 初始化 PostgreSQL
	dsn := "host=localhost user=postgres password=yourpassword dbname=todo_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Todo{})

	// 2. 初始化 Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	r := gin.Default()

	// 示例：获取 Todo 并尝试从缓存读取
	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		cacheKey := "todo:" + id

		// --- 尝试从 Redis 读取缓存 ---
		val, err := rdb.Get(ctx, cacheKey).Result()
		if err == nil {
			// 缓存命中
			var todo Todo
			json.Unmarshal([]byte(val), &todo)
			c.JSON(http.StatusOK, gin.H{"data": todo, "source": "cache"})
			return
		}

		// --- 缓存未命中，查数据库 ---
		var todo Todo
		if err := db.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到"})
			return
		}

		// --- 写入 Redis 缓存 (有效期 10 分钟) ---
		todoJSON, _ := json.Marshal(todo)
		rdb.Set(ctx, cacheKey, todoJSON, 10*time.Minute)

		c.JSON(http.StatusOK, gin.H{"data": todo, "source": "database"})
	})

	r.Run(":8080")
}
