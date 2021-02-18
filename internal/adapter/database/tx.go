package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nglogic/go-example-project/internal/app"
	"github.com/sirupsen/logrus"
)

func rollbackTx(ctx context.Context, tx *sqlx.Tx, l logrus.FieldLogger) {
	if err := tx.Rollback(); err != nil {
		app.AugmentLogFromCtx(ctx, l).Errorf("rolling back postgres transaction: %v", err)
	}
}
