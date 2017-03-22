package helper

import (
	"net/http"

	"github.com/gorilla/schema"
)

const (
	SEARCH_EQUAL string = "equal"
	SEARCH_LIKE  string = "like"
	EQUAL               = " = "
	LIKE                = " like "
)

type QueryString struct {
	Order     string
	Direction string
	Limit     string
	Search    []SearchField
}

type SearchField struct {
	Key      string
	Value    string
	Type     string
	Dropdown DropdownInterface
}

func ParseQueryString(r *http.Request) (*QueryString, error) {
	q := QueryString{}
	err := schema.NewDecoder().Decode(&q, r.URL.Query())
	if err != nil {
		return nil, err
	}
	return &q, nil
}
