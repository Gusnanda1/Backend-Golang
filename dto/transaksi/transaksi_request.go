package transaksidto

type CreateTransaksiRequest struct {
	CounterQTY int    `json:"counter_qty" form:"counter_qty" `
	Total      int    `json:"total" form:"total" `
	Status     string `json:"status" form:"status" `
	Image      string `json:"image" form:"image" `
	Trip_id    int    `json:"trip_id" form:"trip_id" `
	UserID     int    `json:"user_id" form:"user_id"`
}

type UpdateTransaksiRequest struct {
	CounterQTY int    `json:"counter_qty" form:"counter_qty"`
	Total      int    `json:"total" form:"total" `
	Status     string `json:"status" form:"status" `
	Image      string `json:"image" form:"image" `
	Trip_id    int    `json:"trip_id" form:"trip_id" `
	User_id    int    `json:"user_id" form:"user_id" `
}
