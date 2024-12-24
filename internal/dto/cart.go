package dto

type AddToCartRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

type AddToCartResponse struct {
	SuccessTakeStock bool
	Id               int64
	ProductName      string
	Price            int64
}

type DeleteCartRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
}
