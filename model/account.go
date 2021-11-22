package model

type InsertAccount struct {
	ID     int32  `json:"id"  binding:"required"`
	Name   string `json:"name" binding:"required"`
	Number string `json:"number" binding:"required,len=10"`
}

type QueryAccount struct {
	Account []InsertAccount `json:"account" binding:"required"`
}

type DeleteAccounts struct {
	Name string `json:"name" binding:"required"`
}
