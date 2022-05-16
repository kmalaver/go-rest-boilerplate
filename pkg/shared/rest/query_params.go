package rest

import (
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
)

// strategy to use for query params with complex queries

// where
// name__in=a,b,c and
// (age__gt=20 or hasPets=true)

// /users/?filter=[{name__in:a,b,c},{$or:[{age__gt:20},{hasPets:true}]}]&order_by=name__asc&limit=10&offset=0

// /users/?filter={
// 	name:{_in:[a,b,c]},
// 	age:{_gt:20},
// 	order_by: {name:'asc'},
// }
// 	&limit:10,
// 	&offset:0
//

type Filter struct {
	Name     string
	Operator string
	Value    interface{}
}

type Filters []Filter

func GetFilters(c echo.Context) Filters {
	var filters Filters

	queryParam := c.QueryParam("filter")
	if queryParam == "" {
		return filters
	}

	return filters
}

type OrderBy struct {
	Name      string
	Direction string
}

type Params struct {
	Filters Filters
	OrderBy OrderBy
	Limit   uint64
	Offset  uint64
}

var UserParams = Params{
	Filters: Filters{
		{
			Name:     "name",
			Operator: "in",
			Value:    []string{"a", "b", "c"},
		},
		{
			Name:     "age",
			Operator: "gt",
			Value:    20,
		},
	},
	OrderBy: OrderBy{
		Name:      "name",
		Direction: "asc",
	},
	Limit:  10,
	Offset: 0,
}

// should transform to sql query

func (f *Filter) ToBuilder() dbr.Builder {
	switch f.Operator {
	case "eq":
		return dbr.Eq(f.Name, f.Value)
	case "gt":
		return dbr.Gt(f.Name, f.Value)
	case "gte":
		return dbr.Gte(f.Name, f.Value)
	default:
		return dbr.Eq(f.Name, f.Value)
	}
}

func (f Filters) ToBuilder() dbr.Builder {
	var filters []dbr.Builder
	for _, filter := range f {
		filters = append(filters, filter.ToBuilder())
	}
	return dbr.And(filters...)
}

// func ListUsers(db dbr.Session, params Params) ([]user, error){
// 	db.Select("*").
// 		From("users").
// 		Where(params.Filters.ToBuilder()).
// 		OrderBy(params.OrderBy.Name).
// 		Limit(params.Limit).
// 		Offset(params.Offset).

// }
