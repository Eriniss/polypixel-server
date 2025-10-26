package models

import (
	"gorm.io/gorm"
)

// Post는 데이터베이스의 게시물 테이블 구조를 정의합니다.
// gorm.Model은 ID, CreatedAt, UpdatedAt, DeletedAt 필드를 자동으로 포함합니다.
type Post struct {
	gorm.Model // ID, CreatedAt(게시일), UpdatedAt, DeletedAt 필드 자동 포함

	Title    string `gorm:"type:varchar(255);not null"` // 게시물 제목
	Contents string `gorm:"type:text;not null"`         // 게시물 내용
}
