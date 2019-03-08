package model

import "time"

//RPeticion struct
type RPeticion struct {
	ID      int
	Palabra string
}

//RIdioma struct
type RIdioma struct {
	ID     int
	Nombre string
}

//RTraduccion struct
type RTraduccion struct {
	ID                int
	Fecha             time.Time
	PalabraTraduccion string
	PeticionID        int
	IdiomaID          int
}
