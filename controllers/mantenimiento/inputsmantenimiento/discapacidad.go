package inputsmantenimiento

//DiscapacidadInput ...
type DiscapacidadInput struct {
	Descripcion string `json:"descripcion"`
}

//ValidarEntrada ...
func (discapacidad *DiscapacidadInput) ValidarEntrada() bool {
	if discapacidad.Descripcion == "" {
		return true
	}
	return false
}
