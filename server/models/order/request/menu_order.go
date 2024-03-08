package request

type OrderStatusRequest struct {
	ID              uint   `json:"id" validate:"required"`
	Status          string `json:"status" validate:"required"`
	ReasonRejection string `json:"reasonRejection"`
}
