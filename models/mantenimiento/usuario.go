package mantenimiento

//User ...
type User struct {
	UsuarioID uint64 `json:"idusuario" gorm:"primary_key"`
	Usuario   string `json:"usuario"`
	Password  string `json:"password"`
	EmpID     uint   `json:"idempleado" gorm:"not null"`
	Estado    bool   `json:"estado"`
}
