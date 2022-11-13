package response

import "time"

type ProductResponse struct {
	Id       uint      `json:"id"`
	Product  string    `json:"product"`
	CreateAt time.Time `json:"createAt`
	UpdateAt time.Time `json:"updateAt`
}
