package mantenimiento

//Provincia ...
type Provincia struct {
	ProvinciaID uint     `json:"idprovincia" gorm:"primary_key"`
	Descripcion string   `json:"descripcion" gorm:"size:100" gorm:"not null;unique"`
	Estado      bool     `json:"estado"`
	Ciudades    []Ciudad `json:"ciudades" gorm:"foreignkey:ProID"`
}
