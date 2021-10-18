package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/nglogic/go-application-guide/internal/app"
	"github.com/sirupsen/logrus"
)

func rollbackTx(ctx context.Context, tx *sqlx.Tx, l logrus.FieldLogger) {
	err := tx.Rollback()
	if err != nil && !errors.Is(err, sql.ErrTxDone) {
		app.AugmentLogFromCtx(ctx, l).Errorf("rolling back postgres transaction: %v", err)
	}
}

func commitTx(ctx context.Context, tx *sqlx.Tx, l logrus.FieldLogger) error {
	if err := tx.Commit(); err != nil {
		return err
	}

	app.AugmentLogFromCtx(ctx, l).Info("postgres tx committed")
	return nil
}
