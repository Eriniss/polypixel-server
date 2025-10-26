package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB는 애플리케이션 전체에서 사용할 gorm.DB 인스턴스입니다.
var DB *gorm.DB

// loadEnv는 .env 파일을 로드하는 헬퍼 함수입니다.
func loadEnv() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		// .env 파일이 없는 경우에도 오류로 처리하지 않으려면 log.Fatal 대신 log.Println 등을 사용할 수 있습니다.
		// 하지만 여기서는 DB 정보가 필수이므로 Fatal로 처리합니다.
		log.Fatal("❌ .env 파일을 로드하는 데 실패했습니다:", err)
	}
}

// ConnectDB는 PostgreSQL 데이터베이스에 연결합니다.
func ConnectDB() {
	// 1. .env 파일 로드
	loadEnv()

	// 2. 환경 변수에서 DB 정보 가져오기 (os.Getenv 사용)
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// DSN (Data Source Name) 문자열을 구성합니다.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	// GORM을 사용하여 DB 연결을 시도합니다.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ 데이터베이스 연결에 실패했습니다: ", err)
		os.Exit(1)
	}

	log.Println("✅ 데이터베이스 연결 성공!")

	// 전역 DB 변수에 연결 인스턴스를 저장합니다.
	DB = db
}
