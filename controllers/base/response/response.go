package response

import "time"

type ProductResponse struct {
	Id       uint      `json:"id"`
	Product  string    `json:"product"`
	CreateAt time.Time `json:"createAt`
	UpdateAt time.Time `json:"updateAt`
}

type GraphSales struct {
	Label []string `json:"label"`
	Data  []int32  `json:"data"`
}

func FromLabelString(label []string, data []int32) GraphSales {
	return GraphSales{
		Label: label,
		Data:  data,
	}
}
