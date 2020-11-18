package mantenimiento

//ProvinciaInput ...
type ProvinciaInput struct {
	Descripcion string `json:"descripcion"`
	Estado      bool   `json:"estado"`
}

//ValidarEntrada ...
func (provincia *ProvinciaInput) ValidarEntrada() bool {
	if provincia.Descripcion == "" {
		return true
	}
	return false
}
