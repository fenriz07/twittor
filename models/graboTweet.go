package models

import (
	"time"
)

/*GraboTweet Modelo del tweet*/
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
