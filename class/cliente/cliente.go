package cliente

import (
	"github.com/biezhi/gorm-paginator/pagination"
	persona "github.com/frank1995alfredo/api/class/persona"
	database "github.com/frank1995alfredo/api/database"
	clientes "github.com/frank1995alfredo/api/models/clientes"
)

//Cliente ...
type Cliente struct {
	persona.Persona
	clienteID     string
	codigoCliente string
}

//ObtenerClientes ...
func (cliente *Cliente) ObtenerClientes(page int, limit int) *pagination.Paginator {

	var clie []clientes.Cliente

	db := database.DB.Where("estado=?", true).Find(&clie)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"cliente_id asc"},
		ShowSQL: false,
	}, &clie)

	return paginator

}

//CrearCliente ...
func (cliente *Cliente) CrearCliente(discID uint, ciuID uint,
	priNombre string, segNombre string, priApellido string,
	segApellido string, fechaNac string, numCedula string,
	direccion string, email string, telefono string,
	genero string, estado bool, nivelDis string, cargoEmpID uint,
	codigoClie string) string {

	cliente.PersonaConstructor("", discID, ciuID, priNombre, segNombre,
		priApellido, segApellido, fechaNac, numCedula, direccion, email, telefono,
		genero, estado, nivelDis)

	cliente.ClienteConstructor("", codigoClie)

	discapacidadID := cliente.GetDiscID()
	ciudadID := cliente.GetCiuID()
	prinombre := cliente.GetPriNombre()
	segnombre := cliente.GetSegNombre()
	priapellido := cliente.GetPriApellido()
	segapellido := cliente.GetSegApellido()
	fechanac := cliente.GetFechaNac()
	numcedula := cliente.GetNumCedula()
	direcc := cliente.GetDireccion()
	emai := cliente.GetEmail()
	telef := cliente.GetTelefono()
	gene := cliente.GetGenero()
	niveldis := cliente.GetNivelDisacapacidad()
	codigocli := cliente.GetCodigoCli()

	var clien clientes.Cliente

	if cliente.validarEntrada(ciudadID, discapacidadID, prinombre, segnombre, priapellido,
		segapellido, numcedula, codigocli) {
		return "Por favor ingrese los datos que son obligatorios."
	}

	if err := database.DB.Where("num_cedula=?", numCedula).First(&clien).Error; err == nil {
		return "Ya existe este empleado con el mismo número de cédula, ingrese otro."
	}

	datos := clientes.Cliente{DiscID: discapacidadID, CiuID: ciudadID,
		PriNombre: prinombre, SegNombre: segnombre, PriApellido: priapellido, SegApellido: segapellido,
		FechaNac: fechanac, NumCedula: numcedula, CodigoCli: codigocli, Direccion: direcc,
		Email: emai, Telefono: telef, Genero: gene, Estado: true, NivelDis: niveldis}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&datos).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	return "Datos ingresados correctamente."

}

//ActualizarCliente ...
func (cliente *Cliente) ActualizarCliente(ID string, discID uint, ciuID uint,
	priNombre string, segNombre string, priApellido string,
	segApellido string, fechaNac string, numCedula string,
	direccion string, email string, telefono string,
	genero string, estado bool, nivelDis string,
	codigoClie string) string {

	cliente.PersonaConstructor("", discID, ciuID, priNombre, segNombre,
		priApellido, segApellido, fechaNac, numCedula, direccion, email, telefono,
		genero, estado, nivelDis)

	cliente.ClienteConstructor("", codigoClie)

	cliente.SetDiscID(discID)
	cliente.SetCiuID(ciuID)
	cliente.SetPriNombre(priNombre)
	cliente.SetSegNombre(segNombre)
	cliente.SetPriApellido(priApellido)
	cliente.SetSegApellido(segApellido)
	cliente.SetFechaNac(fechaNac)
	cliente.SetNumCedula(numCedula)
	cliente.SetDireccion(direccion)
	cliente.SetEmail(email)
	cliente.SetTelefono(telefono)
	cliente.SetGenero(genero)
	cliente.SetEstado(estado)
	cliente.SetNivelDisacapacidad(nivelDis)
	cliente.SetCodigoCli(codigoClie)

	discapacidadID := cliente.GetDiscID()
	ciudadID := cliente.GetCiuID()
	prinombre := cliente.GetPriNombre()
	segnombre := cliente.GetSegNombre()
	priapellido := cliente.GetPriApellido()
	segapellido := cliente.GetSegApellido()
	fechanac := cliente.GetFechaNac()
	numcedula := cliente.GetNumCedula()
	direcc := cliente.GetDireccion()
	emai := cliente.GetEmail()
	telef := cliente.GetTelefono()
	gene := cliente.GetGenero()
	estad := cliente.GetEstado()
	niveldis := cliente.GetNivelDisacapacidad()
	codigocli := cliente.GetCodigoCli()

	var clie clientes.Cliente

	if cliente.validarEntrada(ciudadID, discapacidadID, prinombre, segnombre, priapellido,
		segapellido, numcedula, codigocli) {
		return "Por favor ingrese los datos que son obligatorios."
	}

	datos := clientes.Cliente{DiscID: discapacidadID, CiuID: ciudadID,
		PriNombre: prinombre, SegNombre: segnombre, PriApellido: priapellido, SegApellido: segapellido,
		FechaNac: fechanac, NumCedula: numcedula, CodigoCli: codigocli, Direccion: direcc,
		Email: emai, Telefono: telef, Genero: gene, Estado: estad, NivelDis: niveldis}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&clie).Where("cliente_id=?", ID).Update(datos).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	return "Registro modificado correctamente."

}

//EliminarCliente ...
func (cliente *Cliente) EliminarCliente(ID string) string {

	cliente.ClienteConstructor(ID, "")

	cliente.SetCliID(ID)
	id := cliente.GetCliID()

	var clie clientes.Cliente

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&clie).Where("cliente_id=?", id).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	return "Registro eliminado correctamente."
}

//BuscarCliente ...
func (cliente *Cliente) BuscarCliente(valor string) ([]clientes.Cliente, string) {

	var cli []clientes.Cliente

	if err := database.DB.Where("num_cedula=?", valor).
		Or("pri_apellido=?", valor).
		Or("seg_apellido=?", valor).
		First(&cli).Error; err != nil {
		return nil, "No existe este empleado."
	}

	return cli, ""

}

//ClienteConstructor ... constructor de la estructura empleado
func (cliente *Cliente) ClienteConstructor(clienteID string, codigoCliente string) {
	cliente.clienteID = clienteID
	cliente.codigoCliente = codigoCliente
}

//GetCliID ...
func (cliente *Cliente) GetCliID() string {
	return cliente.clienteID
}

//SetCliID ...
func (cliente *Cliente) SetCliID(id string) {
	cliente.clienteID = id
}

//GetCodigoCli ...
func (cliente *Cliente) GetCodigoCli() string {
	return cliente.codigoCliente
}

//SetCodigoCli ...
func (cliente *Cliente) SetCodigoCli(codigocli string) {
	cliente.codigoCliente = codigocli
}

func (cliente *Cliente) validarEntrada(ciudadID uint, discapacidadID uint,
	priNombre string, segNombre string, priApellido string,
	segApellido string, numCedula string, codigoClie string) bool {

	if ciudadID == 0 || discapacidadID == 0 || priNombre == "" ||
		segNombre == "" || priApellido == "" || segApellido == "" || numCedula == "" ||
		codigoClie == "" {
		return true
	}

	return false
}
