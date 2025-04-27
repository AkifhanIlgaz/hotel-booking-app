package queries

import (
	"fmt"
	"strings"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/utils"
)

type QueryBuilder struct {
	strings.Builder
}

type ConditionType int

const (
	ConditionEqual ConditionType = iota
	ConditionLike
	ConditionBetween
	ConditionGt
	ConditionIn
)

type Condition struct {
	Type     ConditionType
	Field    string
	Value    string
	Min, Max int
	Values   []string
}

func BuildHotelsQueryWithParams(params models.HotelFilterParams) string {
	var query QueryBuilder
	conditions := []Condition{
		{
			Type:  ConditionLike,
			Field: "city",
			Value: params.City,
		},
		{
			Type:  ConditionLike,
			Field: "country",
			Value: params.Country,
		},
		{
			Type:  ConditionLike,
			Field: "name",
			Value: params.Search,
		},
		{
			Type:  ConditionBetween,
			Field: "price_per_night",
			Min:   params.MinPrice,
			Max:   params.MaxPrice,
		},
		{
			Type:  ConditionGt,
			Field: "rating",
			Min:   int(params.MinRating),
		},
	}

	for _, f := range params.Features {
		conditions = append(conditions, Condition{
			Type:  ConditionLike,
			Field: "features",
			Value: f,
		})
	}

	query.WriteString("SELECT * FROM hotels ")
	query.buildWhereClause(conditions...)
	query.buildOrderByClause(params.SortBy, params.SortOrder)
	query.buildPagination(1, 3)

	return query.String()
}

func (qb *QueryBuilder) buildOrderByClause(sortBy, sortOrder string) {
	sortBy = utils.CamelToSnakeCase(sortBy)

	qb.WriteString(fmt.Sprintf("ORDER BY %v %v ", sortBy, sortOrder))
}

func (qb *QueryBuilder) buildWhereClause(conditions ...Condition) {
	qb.WriteString("WHERE ")

	for i, c := range conditions {
		if i > 0 {
			qb.WriteString(" AND ")
		}

		switch c.Type {
		case ConditionLike:
			qb.likeClause(c)
		case ConditionBetween:
			qb.betweenClause(c)
		case ConditionGt:
			qb.gtClause(c)
		case ConditionIn:
			qb.inClause(c)
		}
	}
}

func (qb *QueryBuilder) likeClause(c Condition) {
	qb.WriteString(fmt.Sprintf("%v LIKE '%%%v%%' ", c.Field, c.Value))
}

func (qb *QueryBuilder) betweenClause(c Condition) {
	qb.WriteString(fmt.Sprintf("%s BETWEEN %d AND %d ", c.Field, c.Min, c.Max))
}

func (qb *QueryBuilder) gtClause(c Condition) {
	qb.WriteString(fmt.Sprintf("%v >= %v ", c.Field, c.Min))
}

func (qb *QueryBuilder) inClause(c Condition) {
	var sb strings.Builder

	for i, v := range c.Values {
		if i > 0 {
			sb.WriteRune(',')
		}

		sb.WriteString(fmt.Sprintf("'%v'", v))
	}

	qb.WriteString(fmt.Sprintf("%v IN (%v) ", c.Field, sb.String()))
}

func (qb *QueryBuilder) buildPagination(page, pageSize int) {
	qb.WriteString(fmt.Sprintf("OFFSET %v ROWS ", (page-1)*pageSize))
	qb.WriteString(fmt.Sprintf("FETCH NEXT %v ROWS ONLY ", pageSize))
}

const InsertHotelQuery = `
INSERT INTO hotels (
			id,
			name,
			description,
			city,
			country,
			image_url,
			price_per_night,
			rating,
			phone_number,
			features,
			created_at
		)
		ValueS (
			@id, @name, @description, @city, @country, @image_url, @price_per_night, @rating, @phone_number, @features, @created_at
		);
`

const SelectHotelById = `
SELECT * FROM hotels WHERE id = @id;
`
