package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type QueryOwner struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id" omitempty`
	Name            string             `json:"name"  binding:"required"`
	Address         string             `json:"address"  binding:"required"`
	WaterBill       int32              `json:"waterBill"  binding:"required"`
	ElectricityBill int32              `json:"electricityBill"  binding:"required"`
	ServicePerDate  int32              `json:"servicePerDate"  binding:"required"`
	Tel             string             `json:"tel"  binding:"required,len=10" `
}

type UpdateOwner struct {
	Name            string `json:"name" bson:"name" binding:"required"`
	Address         string `json:"address" bson:"address"  binding:"required"`
	WaterBill       int32  `json:"waterBill"  bson:"waterBill"  binding:"required"`
	ElectricityBill int32  `json:"electricityBill" bson:"electricityBill"  binding:"required"`
	ServicePerDate  int32  `json:"servicePerDate" bson:"servicePerDate"  binding:"required"`
	Tel             string `json:"tel" bson:"tel" binding:"required,len=10"`
}
