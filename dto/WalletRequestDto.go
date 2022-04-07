package dto

type WalletRequest struct{
	FirstName          string `json:"firstName" validate:"required" bson:"firstName"`
	LastName          string	`json:"lastName" validate:"required" bson:"lastName"`
}
