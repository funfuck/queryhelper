package queryhelper

import (
	"html/template"

	"strings"

	"github.com/jinzhu/gorm"
)

const (
	INNER_JOIN        string = "inner join "
	LEFT_JOIN         string = "left join "
	RIGHT_JOIN        string = "right join "
	ON                string = " on "
	VALUE             string = " ? "
	SPECIAL_CHARACTER string = ` 'ç' `
	ESCAPE            string = " ESCAPE "
)

type Join struct {
	Type string
	Cond string
}

type QueryInterface interface {
	GetSearchFields() []SearchField
	GetSelectFields() []SelectFields
	RelatedTables() map[string]Join
	TableName() string
	Scan(*gorm.DB) (interface{}, error)
	GenTableHtml(interface{}, int) template.HTML
}
type QueryFactory struct {
	Q      QueryInterface
	DB     *gorm.DB
	Req    *QueryString
	Flash  map[string]string
	result interface{}
}

type SelectFields struct {
	Field  string
	Label  string
	Hide   bool
	NoSort bool
}

func (fac *QueryFactory) SearchForm() template.HTML {
	searchFields := fac.Q.GetSearchFields()
	bindSearchQueryToSeconditionrchFields(fac.Req.Search, &searchFields)
	return GenSearchForm(searchFields)
}

func (fac *QueryFactory) Count() (int64, error) {
	var count int64
	tableName := fac.Q.TableName()
	relatedTables := fac.Q.RelatedTables()
	searchFields := fac.Q.GetSearchFields()

	q := joinAndWhere(fac.DB, tableName, relatedTables, searchFields, fac.Req.Search)
	if err := q.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (fac *QueryFactory) FindAll() (interface{}, error) {
	tableName := fac.Q.TableName()
	// likeFields := fac.Q.GetLikeFields()
	selectFields := fac.Q.GetSelectFields()
	relatedTables := fac.Q.RelatedTables()

	// maker select array
	arrSel := []string{}
	for _, v := range selectFields {
		arrSel = append(arrSel, v.Field)
	}

	// select
	q := fac.DB.Select(arrSel)

	// joins and where
	searchFields := fac.Q.GetSearchFields()
	q = joinAndWhere(q, tableName, relatedTables, searchFields, fac.Req.Search)

	// limit
	if fac.Req.Limit == 0 {
		fac.Req.Limit = 10
	}
	q = q.Limit(fac.Req.Limit)
	if fac.Req.Order != "" {
		q = q.Order(fac.Req.Order + " " + fac.Req.Direction)
	}
	q = q.Offset(fac.Req.Offset)

	// scan
	r, err := fac.Q.Scan(q)
	fac.result = r
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (fac *QueryFactory) GenTable() template.HTML {
	return fac.Q.GenTableHtml(fac.result, fac.Req.Offset)
}

func (fac *QueryFactory) FlashMessage() template.HTML {
	return GenAlertFlashMsg(fac.Flash)
}

func joinAndWhere(q *gorm.DB, tableName string, relatedTables map[string]Join, searchFields []SearchField, reqSearch []SearchField) *gorm.DB {
	q = q.Table(tableName)
	// joins
	for k, v := range relatedTables {
		ss := v.Type + k + ON + v.Cond
		q = q.Joins(ss)
	}

	// search
	for k, v := range reqSearch {
		if v.Value == "" {
			continue
		}
		// get search type from search array
		val := searchFields[k]
		v.Type = val.Type

		switch v.Type {
		case SEARCH_LIKE:
			q = q.Where(v.Key+LIKE+VALUE+ESCAPE+SPECIAL_CHARACTER, likeClause(v.Value))
		case SEARCH_EQUAL:
			q = q.Where(v.Key+EQUAL+VALUE, v.Value)
		}
	}
	return q
}

func likeClause(s string) string {
	replacer := strings.NewReplacer("%", `ç%`)
	condition := replacer.Replace(s)
	return "%" + condition + "%"
}

func bindSearchQueryToSeconditionrchFields(searchQuery []SearchField, searchFields *[]SearchField) {
	for k, v := range searchQuery {
		(*searchFields)[k].Value = v.Value
	}
}
