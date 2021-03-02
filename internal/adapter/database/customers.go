package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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
// If customer doesn't exists, returns bikerental.ErrNotFound error.
func (r *CustomersRepository) GetInTx(ctx context.Context, tx *sqlx.Tx, id string) (*bikerental.Customer, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	var m customerModel
	if err := tx.GetContext(ctx, &m, `select * from customers where id = $1`, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrNotFound
		}
		return nil, fmt.Errorf("querying postgres: %w", err)
	}

	result := m.ToAppCustomer()
	return &result, nil
}

// Get returns a customer by id. If it doesn't exists, returns bikerental.ErrNotFound error.
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

// AddInTx creates new customer in db using existing db transaction.
func (r *CustomersRepository) AddInTx(ctx context.Context, tx *sqlx.Tx, c bikerental.Customer) (id string, err error) {
	if c.ID == "" {
		c.ID = uuid.NewString()
	}

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
		return "", fmt.Errorf("building sql query: %w", err)
	}

	if _, err = tx.NamedExecContext(ctx, q, newCustmerModel(c)); err != nil {
		return "", fmt.Errorf("inserting customer row into postgres: %w", err)
	}

	app.AugmentLogFromCtx(ctx, r.log).WithField("id", c.ID).Info("customer created in db")

	return c.ID, nil
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
	case bikerental.CustomerTypeBuisiness:
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
		c.Type = bikerental.CustomerTypeBuisiness
	case customerTypeIndividual:
		c.Type = bikerental.CustomerTypeIndividual
	}
	return c
}
