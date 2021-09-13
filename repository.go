package rajaongkir

import "database/sql"

type RajaongkirRepository struct {
	DB *sql.DB
}

func NewRajaongkirRepository(DB *sql.DB) *RajaongkirRepository {
	return &RajaongkirRepository{
		DB: DB,
	}
}
