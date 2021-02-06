package product

import (
	"database/sql"
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
	Db *sqlx.DB
}

func (pr *SqlProductRepository) All() ([]Product, error) {
	people := []Product{}

	err := pr.Db.Select(&people, "SELECT * FROM products;")

	if err != nil {
		return nil, err
	} else {

		return people, nil
	}
}

func (pr *SqlProductRepository) Get(id int) (*Product, error) {
	person := Product{}
	err := pr.Db.Get(&person, "SELECT * from products WHERE id = $1;", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &person, nil
}

func (pr *SqlProductRepository) Create(p Product) (*Product, error) {
	_, err := pr.Db.NamedExec("INSERT INTO products (id,name,description,price,group_id) VALUES(:id, :name,:description,:price,:group_id);", p)

	if err != nil {
		fmt.Printf("Error creating person: %v", err)
		return nil, err
	}

	return new(Product), nil
}

func (pr *SqlProductRepository) Update(id int, p Product) (*Product, error) {
	p.Id = id
	_, err := pr.Db.NamedExec("UPDATE products SET id=:id, name=:name, description=:description, price=:price, group_id=:group_id WHERE id=:id;", p)

	if err != nil {
		fmt.Printf("Error creating person: %v", err)
		return nil, err
	}

	return new(Product), nil
}

func (pr *SqlProductRepository) Delete(id int) error {
	_, err := pr.Db.Exec("DELETE FROM products WHERE id=$1;", id)

	if err != nil {
		fmt.Printf("Error deleting person: %v", err)
		return err
	}

	return nil
}
