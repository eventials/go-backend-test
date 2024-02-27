package lib

import (
	"fmt"
	"strings"
)

type QueryString struct {
	OrderBy *string           `json:"orderBy"`
	Limit   *int              `json:"limit"`
	Query   map[string][2]any `json:"query"`
}

func (queryString *QueryString) GetLimit() int {
	if queryString.Limit != nil {
		return *queryString.Limit
	} else {
		return -1
	}
}

func (queryString *QueryString) GetQuery() (*string, []any) {
	if len(queryString.Query) == 0 {
		return nil, nil
	}

	var conditions []string
	var args []any

	for key, value := range queryString.Query {
		conditions = append(conditions, fmt.Sprintf("%v %v ?", key, value[0]))
		args = append(args, value[1])
	}

	var where *string

	*where = strings.Join(conditions, " AND ")

	return where, args
}
