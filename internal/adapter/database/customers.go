package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nglogic/go-example-project/internal/app"
)

// CustomersRepository manages customers in db.
type CustomersRepository struct {
	db *sqlx.DB
}

// Get returns a customer by id. If it doesn't exists, returns app.ErrNotFound error.
func (r *CustomersRepository) Get(ctx context.Context, id string) (*app.Customer, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	var m customerModel
	err := r.db.GetContext(
		ctx,
		&m,
		`select * from customers where id = ?`,
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrNotFound
		}
		return nil, fmt.Errorf("querying postgres: %w", err)
	}

	result := m.ToAppCustomer()
	return &result, nil
}

// AddInTx creates new customer in db using existing db transaction.
func (r *CustomersRepository) AddInTx(ctx context.Context, tx *sqlx.Tx, c app.Customer) (id string, err error) {
	if c.ID == "" {
		c.ID = uuid.NewString()
	}

	_, err = tx.NamedExecContext(
		ctx,
		`insert into customers (id, type, first_name, surname, email)
		values (:id, :type, :first_name, :surname, :email)`,
		newCustmerModel(c),
	)
	if err != nil {
		return "", fmt.Errorf("inserting customer row into postgres: %w", err)
	}

	return c.ID, nil
}

type customerModel struct {
	ID        string `db:"id"`
	Type      string `db:"type"`
	FirstName string `db:"first_name"`
	Surname   string `db:"surname"`
	Email     string `db:"email"`
}

func newCustmerModel(ac app.Customer) customerModel {
	c := customerModel{
		ID:        ac.ID,
		FirstName: ac.FirstName,
		Surname:   ac.Surname,
		Email:     ac.Email,
	}
	switch ac.Type {
	case app.CustomerTypeBuisiness:
		c.Type = customerTypeBusiness
	case app.CustomerTypeIndividual:
		c.Type = customerTypeIndividual
	}

	return c
}

func (m *customerModel) ToAppCustomer() app.Customer {
	c := app.Customer{
		ID:        m.ID,
		Type:      0,
		FirstName: m.FirstName,
		Surname:   m.Surname,
		Email:     m.Email,
	}
	switch m.Type {
	case customerTypeBusiness:
		c.Type = app.CustomerTypeBuisiness
	case customerTypeIndividual:
		c.Type = app.CustomerTypeIndividual
	}
	return c
}
