package db

import (
	"fmt"
	"logger/internal/helper/formatter"
	"sync"
)

var once sync.Once
var singleton *QueryBuilder

// QueryBuilder ...
type QueryBuilder struct {
	state string
}

// ExecUpdate ...
func (qb *QueryBuilder) ExecUpdate(table string, columns []string, conditions []Condition) string {
	var i, j int
	field := ""
	length := len(columns)
	query := fmt.Sprintf("UPDATE \"%s\" SET ", table)
	for i = 1; i < length; i++ {
		field = fmt.Sprintf("%s\"%s\" = $%d, ", field, columns[i-1], i)
	}
	field = fmt.Sprintf("%s\"%s\" = $%d ", field, columns[length-1], length)

	//condition
	query = fmt.Sprintf("%s %s WHERE ", query, field)
	lengthCondition := len(conditions) + i
	for j = i + 1; j < lengthCondition; j++ {
		query = fmt.Sprintf("%s\"%s\"=$%d AND ", query, conditions[j-i-1].Key, j)
	}
	query = fmt.Sprintf("%s\"%s\"=$%d", query, conditions[lengthCondition-i-1].Key, lengthCondition)
	return query
}

// ExecInsertAutoIncrement ...
func (qb *QueryBuilder) ExecInsertAutoIncrement(table string, columns []string, primaryKey string) string {
	var i int
	field := ""
	length := len(columns)
	query := fmt.Sprintf("INSERT INTO \"%s\" (", table)
	for i = 0; i < length-1; i++ {
		field = fmt.Sprintf("%s\"%s\", ", field, columns[i])
	}
	field = fmt.Sprintf("%s\"%s\" ", field, columns[length-1])

	query = fmt.Sprintf("%s%s) VALUES(", query, field)
	field = ""
	for i = 1; i < length; i++ {
		field = fmt.Sprintf("%s$%d, ", field, i)
	}
	query = fmt.Sprintf("%s%s$%d) RETURNING \"%s\"", query, field, length, primaryKey)

	return query
}

// ExecDelete ...
func (qb *QueryBuilder) ExecDelete(tables []Table, value []int) string {
	var query string
	for k, v := range tables {
		query = fmt.Sprintf("%s DELETE FROM \"%s\" WHERE \"%s\" = %d; ", query, v.Table, v.PrimaryKey, value[k])
	}
	return query
}

// GenerateCondition ...
func (qb *QueryBuilder) GenerateCondition(table string, conditions []Condition) string {
	query := ""
	if len(conditions) > 0 {
		query = fmt.Sprintf("%s WHERE", query)

		for k, cond := range conditions {
			//init
			if cond.Type != "json" {
				cond.Value = formatter.CleanString(&cond.Value)
			}
			start := ""
			end := ""
			if cond.Operator == "" {
				cond.Operator = "="
			}
			if cond.Table == nil {
				cond.Table = &table
			}
			if cond.Connector == "" {
				cond.Connector = "AND"
			}
			if k == 0 {
				cond.Connector = ""
			}
			if cond.StartBrackets {
				start = "("
			}
			if cond.EndBrackets {
				end = ")"
			}

			//generate
			switch cond.Type {
			case "int":
				val := formatter.StringToInteger(cond.Value)
				query = fmt.Sprintf("%s %s %s \"%s\".\"%s\" %s %d %s", query, cond.Connector, start, *cond.Table, cond.Key, cond.Operator, val, end)
			case "timestamps":
				query = fmt.Sprintf("%s %s %s \"%s\".\"%s\" %s '%s'::date %s", query, cond.Connector, start, *cond.Table, cond.Key, cond.Operator, cond.Value, end)
			case "json":
				query = fmt.Sprintf("%s %s %s %s %s %s %s", query, cond.Connector, start, cond.Key, cond.Operator, cond.Value, end)
			case "array":
				query = fmt.Sprintf("%s %s %s \"%s\".\"%s\" %s %s %s", query, cond.Connector, start, *cond.Table, cond.Key, cond.Operator, cond.Value, end)
			default:
				query = fmt.Sprintf("%s %s %s \"%s\".\"%s\" %s '%s' %s", query, cond.Connector, start, *cond.Table, cond.Key, cond.Operator, cond.Value, end)
			}
		}
	}
	return query
}

// Query ...
func (qb *QueryBuilder) Query(table string, conditions []Condition) string {
	query := fmt.Sprintf("SELECT * FROM \"%s\" %s", table, qb.GenerateCondition(table, conditions))
	return query
}

// QueryOrder ...
func (qb *QueryBuilder) QueryOrder(table string, conditions []Condition, order string) string {
	var ord string
	if order != "" {
		ord = fmt.Sprintf(" ORDER BY %s", order)
	}
	query := fmt.Sprintf("SELECT * FROM \"%s\" %s %s", table, qb.GenerateCondition(table, conditions), ord)
	return query
}

// QueryWithLimit ...
func (qb *QueryBuilder) QueryWithLimit(table string, conditions []Condition, limit int, offset int, order string) string {
	query := ""
	if len(conditions) > 0 {
		query = qb.Query(table, conditions)
		if order != "" {
			query = fmt.Sprintf("%s ORDER BY %s", query, order)
		}
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
	}
	return query
}

// NewQueryBuilder ...
func NewQueryBuilder() *QueryBuilder {
	once.Do(func() {
		singleton = &QueryBuilder{state: "off"}
	})
	return singleton
}
