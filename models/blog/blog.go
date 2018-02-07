package blog

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BlogInt int

const (
	Technology BlogInt = 1001 + iota
	Article BlogInt
)

type Blog struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Type int `json:"type"`
	Labels []string `json:"labels"`
	Title string `json:"title"`
	Content string `json:"content" sql:"type:text"`
	
	UserID uint64 `json:"user_id"`
}

func (b Blog) TableName() string {
	return "blogs"
}

func (b *Blog) AddIndexs(db *gorm.DB) error {
	var err error
	
	err = db.Model(b).AddIndex("idx_blogs_deleted_at","deleted_at")
	if err != nil {
		return err
	}

	return nil
}