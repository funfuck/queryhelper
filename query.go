package helper

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	INNER_JOIN string = "inner join "
	LEFT_JOIN  string = "left join "
	RIGHT_JOIN string = "right join "
	ON         string = " on "
	VALUE      string = " ? "
)

type Join struct {
	Type string
	Cond string
}

type QueryInterface interface {
	GetSearchFields() []SearchField
	GetSelectFields() []string
	RelatedTables() map[string]Join
	TableName() string
	Scan(*gorm.DB) (interface{}, error)
}
type QueryFactory struct {
	Q   QueryInterface
	DB  *gorm.DB
	Req *QueryString
}

func (fac *QueryFactory) FindAll() (interface{}, error) {

	tableName := fac.Q.TableName()
	// searchFields := fac.M.GetSearchFields()
	// likeFields := fac.M.GetLikeFields()
	selectFields := fac.Q.GetSelectFields()
	relatedTables := fac.Q.RelatedTables()

	// select
	q := fac.DB.Select(selectFields).Table(tableName)

	// joins
	for k, v := range relatedTables {
		ss := v.Type + k + ON + v.Cond
		q = q.Joins(ss)
	}

	// search
	// queryString, err := ParseQueryString(fac.Req)
	// if err != nil {
	// 	return nil, err
	// }
	for _, v := range fac.Req.Search {
		fmt.Println(v.Type, SEARCH_EQUAL, SEARCH_LIKE)
		switch v.Type {
		case SEARCH_LIKE:
			q = q.Where(v.Key+LIKE+VALUE, likeClause(v.Value))
		case SEARCH_EQUAL:
			q = q.Where(v.Key+EQUAL+VALUE, v.Value)
		}
	}

	// limit
	if fac.Req.Limit == "" {
		fac.Req.Limit = "10"
	}
	q = q.Limit(fac.Req.Limit)
	if fac.Req.Order != "" {
		q = q.Order(fac.Req.Order + " " + fac.Req.Direction)
	}

	// scan
	r, err := fac.Q.Scan(q)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func likeClause(s string) string {
	return "%" + s + "%"
}
