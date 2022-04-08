package models

type Transaction struct{
	Amount				float64		`json:"amount" bson:"amount" `
}
