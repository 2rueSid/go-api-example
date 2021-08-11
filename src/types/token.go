package types

type Token struct {
	Expiration int    `json:"expiration"`
	Token      string `json:"token"`
	UserId     int    `json:"user_id"`
}
