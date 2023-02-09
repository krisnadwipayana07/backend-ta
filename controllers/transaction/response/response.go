package response

type GraphSales struct {
	Label []string `json:"label"`
	Data  []uint   `json:"data"`
}

func FromLabelString(label []string, data []uint) GraphSales {
	return GraphSales{
		Label: label,
		Data:  data,
	}
}
