package dto

type ShortURLRequest struct {
	LongURL string `json:"long_url"`
	UserID  uint64 `json:"user_id"`
}
