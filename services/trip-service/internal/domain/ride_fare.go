package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID                primitive.ObjectID
	UserId            string
	packageSlug       string // van luxury etc ..
	TotalePriceInCent float32
}
