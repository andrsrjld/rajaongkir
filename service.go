package rajaongkir

type RajaongkirService struct {
	Repository *RajaongkirRepository
}

type RajaongkirAgent interface {
}

func NewRajaongkirService() RajaongkirAgent {
	db := InitDB()

	repo := NewRajaongkirRepository(db)
	return &RajaongkirService{
		Repository: repo,
	}
}
