package models

import (
//"time"
)

type Product struct {
	Id   int64  `json:"id"`
	Code string `json:"code"       binding:"required"`
	Name string `json:"name"       binding:"required"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
}

type Products struct {
	Collection []Product `json:"products"`
}

func (p Product) TableName() string {
	return "products"
}

// Timestamps

//func (p *Product) PreInsert(s gorp.SqlExecutor) error {
//p.CreatedAt = time.Now()
//p.UpdatedAt = p.CreatedAt
//return nil
//}

//func (p *Product) PreUpdate(s gorp.SqlExecutor) error {
//p.UpdatedAt = time.Now()
//return nil
//}
