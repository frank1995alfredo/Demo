package persona

//Persona ... esta  estructura tiene atributos privados
type Persona struct {
	personaID   uint
	discID      uint
	ciuID       uint
	priNombre   string
	segNombre   string
	priApellido string
	segApellido string
	fechaNac    string
	numCedula   string
	direccion   string
	email       string
	telefono    string
	genero      string
	estado      bool
	nivelDis    string
}

//PersonaConstructor ...  constructor
func (persona *Persona) PersonaConstructor(personaID uint, discID uint, ciuID uint,
	priNombre string, segNombre string, priApellido string, segApellido string, fechaNac string,
	numCedula string, direccion string, email string, telefono string,
	genero string, estado bool, nivelDis string) {

	persona.personaID = personaID
	persona.discID = discID
	persona.ciuID = ciuID
	persona.priNombre = priNombre
	persona.segNombre = segNombre
	persona.priApellido = priApellido
	persona.segApellido = segApellido
	persona.fechaNac = fechaNac
	persona.numCedula = numCedula
	persona.direccion = direccion
	persona.email = email
	persona.telefono = telefono
	persona.genero = genero
	persona.nivelDis = nivelDis
}

//GetPersonaID ...
func (persona *Persona) GetPersonaID() uint {
	return persona.personaID
}

//SetPersonaID ...
func (persona *Persona) SetPersonaID(id uint) {
	persona.personaID = id
}

//GetDiscID ...
func (persona *Persona) GetDiscID() uint {
	return persona.discID
}

//SetDiscID ...
func (persona *Persona) SetDiscID(id uint) {
	persona.discID = id
}

//GetCiuID ...
func (persona *Persona) GetCiuID() uint {
	return persona.ciuID
}

//SetCiuID ...
func (persona *Persona) SetCiuID(id uint) {
	persona.ciuID = id
}

//GetPriNombre ...
func (persona *Persona) GetPriNombre() string {
	return persona.priNombre
}

//SetPriNombre ...
func (persona *Persona) SetPriNombre(nombre string) {
	persona.priNombre = nombre
}

//GetSegNombre ...
func (persona *Persona) GetSegNombre() string {
	return persona.segNombre
}

//SetSegNombre ...
func (persona *Persona) SetSegNombre(nombre string) {
	persona.segNombre = nombre
}

//GetPriApellido ...
func (persona *Persona) GetPriApellido() string {
	return persona.priApellido
}

//SetPriApellido ...
func (persona *Persona) SetPriApellido(apellido string) {
	persona.priApellido = apellido
}

//GetSegApellido ...
func (persona *Persona) GetSegApellido() string {
	return persona.segApellido
}

//SetSegApellido ...
func (persona *Persona) SetSegApellido(apellido string) {
	persona.segApellido = apellido
}

//GetFechaNac ...
func (persona *Persona) GetFechaNac() string {
	return persona.fechaNac
}

//SetFechaNac ...
func (persona *Persona) SetFechaNac(fecha string) {
	persona.fechaNac = fecha
}

//GetNumCedula ...
func (persona *Persona) GetNumCedula() string {
	return persona.numCedula
}

//SetNumCedula ...
func (persona *Persona) SetNumCedula(numcedula string) {
	persona.numCedula = numcedula
}

//GetDireccion ...
func (persona *Persona) GetDireccion() string {
	return persona.direccion
}

//SetDireccion ...
func (persona *Persona) SetDireccion(direccion string) {
	persona.direccion = direccion
}

//GetEmail ...
func (persona *Persona) GetEmail() string {
	return persona.email
}

//SetEmail ...
func (persona *Persona) SetEmail(email string) {
	persona.email = email
}

//GetTelefono ...
func (persona *Persona) GetTelefono() string {
	return persona.telefono
}

//SetTelefono ...
func (persona *Persona) SetTelefono(telefono string) {
	persona.telefono = telefono
}

//GetGenero ...
func (persona *Persona) GetGenero() string {
	return persona.genero
}

//SetGenero ...
func (persona *Persona) SetGenero(genero string) {
	persona.genero = genero
}

//GetEstado ...
func (persona *Persona) GetEstado() bool {
	return persona.estado
}

//SetEstado ...
func (persona *Persona) SetEstado(estado bool) {
	persona.estado = estado
}

//GetNivelDisacapacidad ...
func (persona *Persona) GetNivelDisacapacidad() string {
	return persona.nivelDis
}

//SetNivelDisacapacidad ...
func (persona *Persona) SetNivelDisacapacidad(nivel string) {
	persona.nivelDis = nivel
}
