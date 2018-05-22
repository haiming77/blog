package blog

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Flag struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Flag string `json:"flag" sql:"type:text"`
}

func (f Flag) TableName() string {
	return "flags"
}

func (f *Flag) AddIndexs(db *gorm.DB) error {
	var err error

	err = db.Model(f).AddIndex("idx_flags_deleted_at", "deleted_at").Error
	if err != nil {
		return err
	}

	return nil
}
