package bd

import (
	"context"
	"time"

	"github.com/fenriz07/twittor/models"
)

/*BorroRelacion borra una relacion en la coleccion*/
func BorroRelacion(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
