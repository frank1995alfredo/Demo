package practica

//Persona ...
type Persona struct {
	PersonaID uint       `json:"idpersona" gorm:"primary_key"`
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Direccion string     `json:"direccion"`
	Ciudad    string     `json:"ciudad"`
	Contactos []Contacto `json:"contactos"  gorm:"foreignkey:PerID"`
}

//Contacto ...
type Contacto struct {
	ContactoID uint   `json:"idcontacto" gorm:"primary_key"`
	Telefono   string `json:"telefono"`
	Email      string `json:"email"`
	PerID      uint   `json:"idpersona" gorm:"not null"`
}
