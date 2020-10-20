package atlas

type (

	//User
	Users struct {
		ID       int64   `json:"id"`
		Name     string  `json:"name"`
		Email    string  `json:"email"`
		Password *string `json:"password,omitempty"`
	}

	//TokenData estrutura claims  do JWT
	TokenData struct {
		IDUser    int64
		IDCompany int64
		NameUser  string
		Password  string
		Role      int
	}

	//Paginate estrutura paginação
	Paginate struct {
		Query      string
		Total      int64
		Offset     int
		Limit      int
		NamedQuery map[string]interface{} //sqlx
	}

	//JWTService serviço para validação e geração do jwt
	JwtServices interface {
		GenerateToken(data TokenData) (string, error)
		ParseAndVerifyToken(token string) (*TokenData, error)
	}

	//UsersServices  serviços para busca, criação e autenticação do usuário
	UsersServices interface {
		Find(page *Paginate) ([]Users, error)
		Create(payload *Users) error
		AuthToken(email string) (*TokenData, error)
	}
)
