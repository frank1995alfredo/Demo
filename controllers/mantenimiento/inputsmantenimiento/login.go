package inputsmantenimiento

//UserInput ...
type UserInput struct {
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
	EmpID    uint   `json:"idempleado"`
	Estado   bool   `json:"estado"`
}

//ValidarEntrada ...
func (user *UserInput) ValidarEntrada() bool {
	if user.Usuario == "" || user.Password == "" {
		return true
	}
	return false
}
