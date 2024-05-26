package users

import (
	"database/sql"
	"fmt"

	"github.com/jaydto/goApiMyql/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}

}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * from user WHERE email=?", email)

	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return u, nil

}

func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *Store) GetUserById(id int) (*types.User, error) {
	return nil, nil

}

func (s *Store) CreateUser(user types.User) error {
	return nil

}