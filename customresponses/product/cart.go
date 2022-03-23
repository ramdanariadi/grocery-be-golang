package product

type CartResponse struct {
	Id       string `json:"id"`
	Price    int64  `json:"price"`
	Weight   uint   `json:"weight"`
	Category string `json:"category"`
	PerUnit  int    `json:"perUnit"`
	ImageUrl string `json:"imageUrl"`
	Name     string `json:"name"`
	Total    int    `json:"total"`
}