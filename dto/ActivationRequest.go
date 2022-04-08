package dto

type ActivationRequest struct{
	Active          bool	`json:"active" validate:"required" bson:"active"`
}
