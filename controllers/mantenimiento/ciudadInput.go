package mantenimiento

//CiudadInput ...
type CiudadInput struct {
	ProID       uint   `json:"idprovincia"`
	Descripcion string `json:"descripcion"`
	Estado      bool   `json:"estado"`
}

//VerificarDescripcion ...
func (ciudad *CiudadInput) VerificarDescripcion() bool {
	if ciudad.Descripcion == "" || ciudad.ProID == 0 {
		return true
	}
	return false
}
