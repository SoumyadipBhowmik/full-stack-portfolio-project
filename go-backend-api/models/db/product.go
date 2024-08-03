package models

import "github.com/google/uuid"

type Product struct {
	ID            uuid.UUID     `db:"id"`
	Name          string        `db:"name"`
	ImageUrl      string        `db:"description"`
	Stock         int64         `db:"stock"`
	Price         int32         `db:"price"`
	Description   string        `db:"description"`
	Specification Specification `db:"specification"`
}
