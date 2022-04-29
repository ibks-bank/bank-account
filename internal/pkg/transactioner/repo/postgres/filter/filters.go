package filter

import (
	"fmt"
	"strings"
	"time"
)

type Filter func(query string, args []interface{}) (string, []interface{})

func ByDateFrom(date time.Time) Filter {
	return func(query string, args []interface{}) (string, []interface{}) {
		args = append(args, date)
		query = withCondition(query) + fmt.Sprintf(" created_at >= $%d ", len(args))
		return query, args
	}
}

func ByDateTo(date time.Time) Filter {
	return func(query string, args []interface{}) (string, []interface{}) {
		args = append(args, date)
		query = withCondition(query) + fmt.Sprintf(" created_at <= $%d ", len(args))
		return query, args
	}
}

func withCondition(query string) string {
	if !strings.Contains(query, "where") {
		return query + " where "
	}

	return query + " and "
}
