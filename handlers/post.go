package handlers

import (
	"log"
	"net/http"

	"polypixel-server/database"
	"polypixel-server/models"

	"github.com/gofiber/fiber/v3"
)

// CreatePost는 새 게시물을 생성하고 데이터베이스에 저장합니다.
func CreatePost(c fiber.Ctx) error {
	// 1. 요청 본문(Body)을 저장할 변수 (Post 모델)를 초기화합니다.
	post := new(models.Post)

	// 2. 요청 본문을 파싱하여 post 객체에 바인딩합니다.
	// Fiber의 BodyParser는 JSON, XML, Form 데이터를 처리할 수 있습니다.
	if err := c.Bind().Body(post); err != nil {
		log.Printf("Post 요청 본문 파싱 에러: %v", err)
		// 잘못된 요청 형식임을 나타내는 400 Bad Request를 반환합니다.
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "요청 본문이 유효하지 않습니다.",
		})
	}

	// 3. 데이터 유효성 검사 (아주 간단한 검사)
	if post.Title == "" || post.Contents == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "제목과 내용은 필수로 입력해야 합니다.",
		})
	}

	// 4. GORM을 사용하여 데이터베이스에 새 레코드를 생성합니다.
	// GORM은 자동으로 CreatedAt, UpdatedAt 필드를 채워줍니다.
	result := database.DB.Create(&post)

	// 5. 생성 과정에서 오류가 발생했는지 확인합니다.
	if result.Error != nil {
		log.Printf("DB Post 생성 에러: %v", result.Error)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "게시물 생성 중 서버 오류가 발생했습니다.",
		})
	}

	// 6. 성공적인 응답을 반환합니다. (201 Created)
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "게시물이 성공적으로 생성되었습니다.",
		"post":  post, // DB에서 ID와 타임스탬프가 채워진 객체를 반환합니다.
	})
}
