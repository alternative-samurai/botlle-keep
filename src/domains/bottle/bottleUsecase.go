package bottle

import (
	"context"
	"example/src/infrastructures/entities"
	"example/src/infrastructures/model"
	"example/util"
	"time"
)

func CreateBottle(ctx context.Context, bottleName string) string {

	id := util.GenUuidV4()
	createdAt := ctx.Value("reqDate").(time.Time)
	entities.Insert(&model.Bottle{
		BottleID:   id,
		BottleName: bottleName,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	})
	return id
}
