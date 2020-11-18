package usuarios

//UsuarioInput ...
type UsuarioInput struct {
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
	EmpID    uint   `json:"idempleado"`
}

//ValidarEntrada ...
func (user *UsuarioInput) ValidarEntrada() bool {
	if user.Usuario == "" || user.Password == "" {
		return true
	}
	return false
}
