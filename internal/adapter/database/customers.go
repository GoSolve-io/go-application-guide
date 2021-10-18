package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/nglogic/go-application-guide/internal/app"
	"github.com/nglogic/go-application-guide/internal/app/bikerental"
	"github.com/sirupsen/logrus"
)

// CustomersRepository manages customers in db.
type CustomersRepository struct {
	db  *sqlx.DB
	log logrus.FieldLogger
}

// GetInTx returns a customer by id using existing transaction.
// If customer doesn't exists, returns app.ErrNotFound error.
func (r *CustomersRepository) GetInTx(ctx context.Context, tx *sqlx.Tx, id string) (*bikerental.Customer, error) {
	var m customerModel
	if err := tx.GetContext(ctx, &m, `select * from customers where id = $1`, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, app.ErrNotFound
		}
		return nil, fmt.Errorf("querying postgres: %w", err)
	}

	result := m.ToAppCustomer()
	return &result, nil
}

// Get returns a customer by id. If it doesn't exists, returns app.ErrNotFound error.
func (r *CustomersRepository) Get(ctx context.Context, id string) (*bikerental.Customer, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("creating postgresql transaction: %w", err)
	}
	defer func() {
		_ = tx.Commit()
	}()

	return r.GetInTx(ctx, tx, id)
}

// CreateInTx creates new customer in db using existing db transaction.
func (r *CustomersRepository) CreateInTx(ctx context.Context, tx *sqlx.Tx, c bikerental.Customer) error {
	sqlq := sqlBuilder.Insert("customers").
		Columns("id", "type", "first_name", "surname", "email").
		Values(
			squirrel.Expr(":id"),
			squirrel.Expr(":type"),
			squirrel.Expr(":first_name"),
			squirrel.Expr(":surname"),
			squirrel.Expr(":email"),
		)
	q, _, err := sqlq.ToSql()
	if err != nil {
		return fmt.Errorf("building sql query: %w", err)
	}

	if _, err = tx.NamedExecContext(ctx, q, newCustmerModel(c)); err != nil {
		return fmt.Errorf("inserting customer row into postgres: %w", err)
	}

	app.AugmentLogFromCtx(ctx, r.log).WithField("id", c.ID).Info("customer created in db")

	return nil
}

// Create creates new customer in db.
func (r *CustomersRepository) Create(ctx context.Context, c bikerental.Customer) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("creating postgresql transaction: %w", err)
	}
	defer func() {
		_ = tx.Commit()
	}()

	return r.CreateInTx(ctx, tx, c)
}

// Delete removes customer from db.
func (r *CustomersRepository) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, `delete from customers where id=$1`, id)
	if err != nil {
		return fmt.Errorf("deleting customer row from postgres: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return app.ErrNotFound
	}

	app.AugmentLogFromCtx(ctx, r.log).WithField("id", id).Info("customer deleted from db")

	return nil
}

type customerModel struct {
	ID        string `db:"id"`
	Type      string `db:"type"`
	FirstName string `db:"first_name"`
	Surname   string `db:"surname"`
	Email     string `db:"email"`
}

func newCustmerModel(ac bikerental.Customer) customerModel {
	c := customerModel{
		ID:        ac.ID,
		FirstName: ac.FirstName,
		Surname:   ac.Surname,
		Email:     ac.Email,
	}
	switch ac.Type {
	case bikerental.CustomerTypeBusiness:
		c.Type = customerTypeBusiness
	case bikerental.CustomerTypeIndividual:
		c.Type = customerTypeIndividual
	}

	return c
}

func (m *customerModel) ToAppCustomer() bikerental.Customer {
	c := bikerental.Customer{
		ID:        m.ID,
		Type:      0,
		FirstName: m.FirstName,
		Surname:   m.Surname,
		Email:     m.Email,
	}
	switch m.Type {
	case customerTypeBusiness:
		c.Type = bikerental.CustomerTypeBusiness
	case customerTypeIndividual:
		c.Type = bikerental.CustomerTypeIndividual
	}
	return c
}
