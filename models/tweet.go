package models

/*Tweet captura el body que nos envian*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
