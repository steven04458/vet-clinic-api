package models

type CatResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Breed  string  `json:"breed"`
	Weight float32 `json:"weight"`
}

type CatCreateRequest struct {
	Name   string  `json:"name" validate:"required"`
	Age    int     `json:"age" validate:"gte=0"`
	Breed  string  `json:"breed"`
	Weight float32 `json:"weight" validate:"gte=0"`
}
