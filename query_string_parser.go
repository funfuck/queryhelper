package queryhelper

import (
	"net/http"

	"github.com/gorilla/schema"
)

const (
	SEARCH_EQUAL string = "equal"
	SEARCH_LIKE  string = "like"
	EQUAL        string = " = "
	LIKE         string = " like "
)

type QueryString struct {
	Order     string
	Direction string
	Limit     int
	Offset    int
	P         int
	Search    []SearchField
}

func (q *QueryString) SearchMap() map[string]string {
	m := make(map[string]string)
	for _, v := range q.Search {
		m[v.Key] = v.Value
	}
	return m
}

type SearchField struct {
	Key      string
	Value    string
	Type     string
	Label    string
	Dropdown DropdownInterface
}

func ParseQueryString(r *http.Request) (*QueryString, error) {
	q := QueryString{}
	err := schema.NewDecoder().Decode(&q, r.URL.Query())
	q.Offset = q.P - 1
	if err != nil {
		return nil, err
	}
	return &q, nil
}
