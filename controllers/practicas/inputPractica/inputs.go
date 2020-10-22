package inputs

//PersonaInput ...
type PersonaInput struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Direccion string `json:"direccion"`
	Ciudad    string `json:"ciudad"`
}

//ContactoInput ...
type ContactoInput struct {
	Telefono string `json:"telefono"`
	Email    string `json:"email"`
	PerID    uint   `json:"idpersona"`
}
