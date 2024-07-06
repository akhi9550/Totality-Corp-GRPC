package models

type User struct {
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   string  `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}

type SearchUser struct {
	City    string `json:"city"`
	Phone   string `json:"phone"`
	Married bool   `json:"married"`
}

type Users struct {
	ID      int64   `json:"id"`
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   string  `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}
