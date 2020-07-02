package bd

import (
	"context"
	"time"

	"github.com/fenriz07/twittor/models"
)

/*InsertoRelacion inserta un nuevo registro en la tabla relación*/
func InsertoRelacion(t models.Relation) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
