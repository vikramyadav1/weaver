package product

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:generate mockery --name=ProductRepository

type ProductRepository interface {
	All() ([]Product, error)
	Get(id int) (*Product, error)
	Create(p Product) (*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) error
}

var _ ProductRepository = &SqlProductRepository{}

type SqlProductRepository struct {
	db sqlx.DB
}

func (pr *SqlProductRepository) All() ([]Product, error) {
	people := []Product{}

	err := pr.db.Select(&people, "SELECT * FROM products")

	if err != nil {
		return nil, err
	} else {

		return people, nil
	}
}

func (pr *SqlProductRepository) Get(id int) (*Product, error) {
	person := Product{}
	err := pr.db.Get(&person, "SELECT * from products WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (pr *SqlProductRepository) Create(p Product) (*Product, error) {
	_, err := pr.db.NamedExec("INSERT INTO products (first_name,last_name,email,group_id) VALUES(:first_name, last_name,:email,:group_id)", p)

	if err != nil {
		fmt.Printf("Error creating person: %v", err)
		return nil, err
	}

	return new(Product), nil
}

func (pr *SqlProductRepository) Update(id int, p Product) (*Product, error) {
	p.Id = id
	_, err := pr.db.NamedExec("UPDATE products SET first_name=:first_name, last_name=:last_name, email=:email, group_id=:group_id WHERE id=:id", p)

	if err != nil {
		fmt.Printf("Error creating person: %v", err)
		return nil, err
	}

	return new(Product), nil
}

func (pr *SqlProductRepository) Delete(id int) error {
	_, err := pr.db.NamedExec("DELETE FROM products WHERE id=:id", id)

	if err != nil {
		fmt.Printf("Error deleting person: %v", err)
		return err
	}

	return nil
}
