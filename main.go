package main

import (
	"log"

	"polypixel-server/database"
	"polypixel-server/handlers"
	"polypixel-server/models"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// 1. 데이터베이스 연결 (.env 로드 포함)
	database.ConnectDB()

	// 2. 테이블 마이그레이션
	err := database.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatalf("❌ Post 테이블 마이그레이션 실패: %v", err)
	}
	log.Println("✅ Post 테이블 마이그레이션 성공!")

	// 3. Fiber 앱 초기화
	app := fiber.New()

	// 4. API 라우팅 설정
	// /api/v1 그룹을 만들고 그 안에 POST /posts 경로를 정의합니다.
	api := app.Group("/api/v1")
	api.Post("/posts", handlers.CreatePost)

	// 기본 테스트 경로
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("API is running...")
	})

	// 5. 서버 실행
	log.Fatal(app.Listen(":3000"))
}
