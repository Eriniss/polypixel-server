package main

import (
	"log"

	"polypixel-server/database"
	models "polypixel-server/models"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// 1. 데이터베이스 연결 (.env 로드 포함)
	database.ConnectDB()

	// 2. 테이블 마이그레이션 (Post 모델로 변경)
	err := database.DB.AutoMigrate(&models.Post{}) // 👈 models.Post로 변경
	if err != nil {
		log.Fatalf("❌ Post 테이블 마이그레이션 실패: %v", err)
	}
	log.Println("✅ Post 테이블 마이그레이션 성공!")

	// 3. Fiber 앱 초기화 및 라우팅 (생략)
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World! (DB 연결 및 Post 마이그레이션 완료)")
	})

	log.Fatal(app.Listen(":3000"))
}
