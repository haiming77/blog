package user


type Profile struct {
	Avatar    string `json:"avatar"`
	BrithDate string `json:"brith_date"`

	UserID uint64 `json:"user_id"`
}

func (p Profile) TableName() string {
	return "profile"
}