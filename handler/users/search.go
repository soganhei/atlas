package users

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/soganhei.com.br/atlas"
)

func search(c *gin.Context) (*atlas.Paginate, error) {

	//Pegar token JWT
	//token := c.MustGet("token").(atlas.TokenData)

	type Uri struct {
		Limit int    `form:"limit"`
		Page  int    `form:"page"`
		Name  string `form:"name"`
	}

	var uri Uri

	if err := c.ShouldBind(&uri); err != nil {
		return nil, err
	}

	if uri.Limit == 0 {
		uri.Limit = 50
	}

	if uri.Page == 0 {
		uri.Page = 1
	}

	paginate := atlas.Paginate{
		Offset: ((uri.Page - 1) * uri.Limit),
		Limit:  uri.Limit,
	}

	var and strings.Builder
	var query string

	namedQuery := map[string]interface{}{}

	if uri.Name != "" {

		uri.Name = strings.ToLower(uri.Name)

		query = " AND LOWER(UNACCENT(u.name)) LIKE '%' || LOWER( UNACCENT(:usernName) ) || '%'"
		and.WriteString(query)

		namedQuery["usernName"] = uri.Name
	}

	paginate.Query = and.String()
	paginate.NamedQuery = namedQuery

	return &paginate, nil
}
