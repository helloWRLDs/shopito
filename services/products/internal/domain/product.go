package domain

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ImageURL   string `json:"img_url"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}
