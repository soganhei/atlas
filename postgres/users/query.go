package users

import (
	"fmt"

	"github.com/soganhei.com.br/atlas"
)

func queryFind(page *atlas.Paginate) string {

	query := fmt.Sprintf(`
		SELECT 
			u.id, 
			u.name, 
			u.email			 
			FROM users u		 
		%s
	`, page.Query)
	return query
}

func queryCreate() string {
	return `
		INSERT 
			INTO users(					
				name,
				email,
				password,					
				created_at,
				updated_at) 
		VALUES($1,$2,$3,$4,$4)`
}

func queryAuthToken() string {
	return "SELECT id, name, password FROM users WHERE email=$1"
}
