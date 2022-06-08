package common

import (
	"context"
	"example/util"
	"time"
)

func createData(ctx context.Context, bottleName string) string {
	id := util.GenUuidV4()
	createdAt := time.Now()
	updatedAt := time.Now()

	return id
}
