package users

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/soganhei.com.br/atlas"
)

type Services struct {
	DB *sqlx.DB
}

func NewServices(DB *sqlx.DB) *Services {
	return &Services{
		DB: DB,
	}
}

func (s *Services) Find(page *atlas.Paginate) ([]atlas.Users, error) {

	users := []atlas.Users{}

	query := queryFind(page)

	rows, err := s.DB.NamedQuery(query, page.NamedQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var data atlas.Users
		err := rows.Scan(
			&data.ID,    //id
			&data.Name,  //name
			&data.Email, //email
		)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	return users, nil
}

func (s *Services) Create(payload *atlas.Users) error {

	query := queryCreate()

	createdAt := time.Now().Format(time.RFC3339Nano)

	_, err := s.DB.Exec(query,
		payload.Name,     //name
		payload.Email,    //email
		payload.Password, //password
		createdAt,        //created_at
	)
	if err != nil {
		return err
	}
	return nil

}

func (s *Services) AuthToken(email string) (*atlas.TokenData, error) {

	var token atlas.TokenData

	query := queryAuthToken()

	err := s.DB.QueryRow(query, email).Scan(
		&token.IDUser,
		&token.NameUser,
		&token.Password,
	)
	if err != nil {
		return nil, nil
	}
	return &token, nil

}
