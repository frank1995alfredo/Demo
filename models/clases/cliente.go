package clases

//persona "github.com/frank1995alfredo/api/models/clases"

//ClienteM ...
type ClienteM struct {
	ClienteID uint   `json:"idcliente" gorm:"primary_key"`
	CodigoCli string `json:"codigoclientepapu"`
}

//ClientePersona ...
type ClientePersona struct {
	PersonaM
	ClienteM
}

func (c ClientePersona) getDiscapacidad() string {
	return c.GetNombres() + " " + c.CodigoCli + " " + c.NivelDis
}
