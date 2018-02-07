package user


import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Mobile            string `json:"mobile"`
	EncryptedPassword string `json:"encrpted_password"`
	Name              string `json:"name"`
	Nickname          string `json:"nickname"`

	Profile *Profile `json:"profile" sql:"ForigenKey:UserId"`
}

func (u User) TableName() string {
	return "users"
}

func (u *User) AddIndexes(db *gorm.DB) error {
	var err error

	err = db.Model(u).AddIndex("idx_users_deleted_at", "deleted_at").Error
	if err != nil {
		return err
	}

	err = db.Model(u).AddUniqueIndex("uix_users_mobile", "mobile").Error
	if err != nil {
		return err
	}

	return nil
}
