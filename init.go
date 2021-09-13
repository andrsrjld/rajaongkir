package rajaongkir

import "log"

var (
	Agent RajaongkirAgent
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("Rajaongkir initializing.")
	Agent = NewRajaongkirService()
	log.Println("Rajaongkir initialized.")
}
