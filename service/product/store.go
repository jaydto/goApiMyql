package product

import (
	"database/sql"
	"fmt"

	"github.com/jaydto/goApiMyql/types"
)

type Store struct {
	db *sql.DB
}

// CreateProductPayload implements types.ProductStore.
func (s *Store) CreateProductPayload(p types.CreateProductPayload) error {
	sql := `INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)`
	result, err := s.db.Exec(sql, p.Name, p.Description, p.Image, p.Price, p.Quantity)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	_, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}
	return nil
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProductById(id int) (*types.Product, error) {

	rows, err := s.db.Query("SELECT * FROM products WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	p := new(types.Product)
	for rows.Next() {
		p, err = ScanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}
	if p.ID == 0 {
		return nil, fmt.Errorf("product not found")
	}
	return p, nil

}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := ScanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)

	}
	return products, nil

}

// CreateProductPayload implements types.ProductStore.

// func (s *Store) CreateProductPayload(p types.Product) error {
// 	sql := `INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)`
// 	_, err := s.db.Exec(sql, p.Name, p.Description, p.Image, p.Price, p.Quantity)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func ScanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.ID,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.Name,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
