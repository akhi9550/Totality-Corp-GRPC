package domain

type User struct {
	ID      int64   `json:"id" gorm:"uniquekey; not null"`
	Fname   string  `json:"fname" gorm:"validate:required"`
	City    string  `json:"city" gorm:"validate:required"`
	Phone   int64   `json:"phone" gorm:"validate:required"`
	Height  float32 `json:"height" gorm:"validate:required"`
	Married bool    `json:"married" gorm:"validate:required"`
}
