package inputsmantenimiento

//CargoInput ...
type CargoInput struct {
	Descripcion string `json:"descripcion"`
}

//VerificarDescripcion ...
func (cargo *CargoInput) VerificarDescripcion() bool {
	if cargo.Descripcion == "" {
		return true
	}
	return false
}
