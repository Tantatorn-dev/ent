// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"fbc/ent/entc/integration/ent/fieldtype"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/__"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// FieldTypeQuery is the builder for querying FieldType entities.
type FieldTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []ent.Predicate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
func (ftq *FieldTypeQuery) Where(ps ...ent.Predicate) *FieldTypeQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit adds a limit step to the query.
func (ftq *FieldTypeQuery) Limit(limit int) *FieldTypeQuery {
	ftq.limit = &limit
	return ftq
}

// Offset adds an offset step to the query.
func (ftq *FieldTypeQuery) Offset(offset int) *FieldTypeQuery {
	ftq.offset = &offset
	return ftq
}

// Order adds an order step to the query.
func (ftq *FieldTypeQuery) Order(o ...Order) *FieldTypeQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// Get returns a FieldType entity by its id.
func (ftq *FieldTypeQuery) Get(ctx context.Context, id string) (*FieldType, error) {
	return ftq.Where(fieldtype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (ftq *FieldTypeQuery) GetX(ctx context.Context, id string) *FieldType {
	ft, err := ftq.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ft
}

// First returns the first FieldType entity in the query. Returns *ErrNotFound when no fieldtype was found.
func (ftq *FieldTypeQuery) First(ctx context.Context) (*FieldType, error) {
	fts, err := ftq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(fts) == 0 {
		return nil, &ErrNotFound{fieldtype.Label}
	}
	return fts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FieldTypeQuery) FirstX(ctx context.Context) *FieldType {
	ft, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ft
}

// FirstID returns the first FieldType id in the query. Returns *ErrNotFound when no id was found.
func (ftq *FieldTypeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{fieldtype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ftq *FieldTypeQuery) FirstXID(ctx context.Context) string {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only FieldType entity in the query, returns an error if not exactly one entity was returned.
func (ftq *FieldTypeQuery) Only(ctx context.Context) (*FieldType, error) {
	fts, err := ftq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(fts) {
	case 1:
		return fts[0], nil
	case 0:
		return nil, &ErrNotFound{fieldtype.Label}
	default:
		return nil, &ErrNotSingular{fieldtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FieldTypeQuery) OnlyX(ctx context.Context) *FieldType {
	ft, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ft
}

// OnlyID returns the only FieldType id in the query, returns an error if not exactly one id was returned.
func (ftq *FieldTypeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{fieldtype.Label}
	default:
		err = &ErrNotSingular{fieldtype.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (ftq *FieldTypeQuery) OnlyXID(ctx context.Context) string {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FieldTypes.
func (ftq *FieldTypeQuery) All(ctx context.Context) ([]*FieldType, error) {
	switch ftq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftq.sqlAll(ctx)
	case dialect.Neptune:
		return ftq.gremlinAll(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// AllX is like All, but panics if an error occurs.
func (ftq *FieldTypeQuery) AllX(ctx context.Context) []*FieldType {
	fts, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return fts
}

// IDs executes the query and returns a list of FieldType ids.
func (ftq *FieldTypeQuery) IDs(ctx context.Context) ([]string, error) {
	switch ftq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftq.sqlIDs(ctx)
	case dialect.Neptune:
		return ftq.gremlinIDs(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FieldTypeQuery) IDsX(ctx context.Context) []string {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FieldTypeQuery) Count(ctx context.Context) (int, error) {
	switch ftq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftq.sqlCount(ctx)
	case dialect.Neptune:
		return ftq.gremlinCount(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FieldTypeQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FieldTypeQuery) Exist(ctx context.Context) (bool, error) {
	switch ftq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftq.sqlExist(ctx)
	case dialect.Neptune:
		return ftq.gremlinExist(ctx)
	default:
		return false, errors.New("ent: unsupported dialect")
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FieldTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FieldType.Query().
//		GroupBy(fieldtype.FieldInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ftq *FieldTypeQuery) GroupBy(field string, fields ...string) *FieldTypeGroupBy {
	group := &FieldTypeGroupBy{config: ftq.config}
	group.fields = append([]string{field}, fields...)
	switch ftq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		group.sql = ftq.sqlQuery()
	case dialect.Neptune:
		group.gremlin = ftq.gremlinQuery()
	}
	return group
}

func (ftq *FieldTypeQuery) sqlAll(ctx context.Context) ([]*FieldType, error) {
	rows := &sql.Rows{}
	selector := ftq.sqlQuery()
	if unique := ftq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := ftq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var fts FieldTypes
	if err := fts.FromRows(rows); err != nil {
		return nil, err
	}
	fts.config(ftq.config)
	return fts, nil
}

func (ftq *FieldTypeQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := ftq.sqlQuery()
	unique := []string{fieldtype.FieldID}
	if len(ftq.unique) > 0 {
		unique = ftq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := ftq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (ftq *FieldTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ftq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ftq *FieldTypeQuery) sqlIDs(ctx context.Context) ([]string, error) {
	vs, err := ftq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (ftq *FieldTypeQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(fieldtype.Table)
	selector := sql.Select(t1.Columns(fieldtype.Columns...)...).From(t1)
	if ftq.sql != nil {
		selector = ftq.sql
		selector.Select(selector.Columns(fieldtype.Columns...)...)
	}
	for _, p := range ftq.predicates {
		p.SQL(selector)
	}
	for _, p := range ftq.order {
		p.SQL(selector)
	}
	if offset := ftq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := ftq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (ftq *FieldTypeQuery) gremlinIDs(ctx context.Context) ([]string, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	vertices, err := res.ReadVertices()
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(vertices))
	for _, vertex := range vertices {
		ids = append(ids, vertex.ID.(string))
	}
	return ids, nil
}

func (ftq *FieldTypeQuery) gremlinAll(ctx context.Context) ([]*FieldType, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().ValueMap(true).Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var fts FieldTypes
	if err := fts.FromResponse(res); err != nil {
		return nil, err
	}
	fts.config(ftq.config)
	return fts, nil
}

func (ftq *FieldTypeQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().Count().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ftq *FieldTypeQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().HasNext().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (ftq *FieldTypeQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(fieldtype.Label)
	if ftq.gremlin != nil {
		v = ftq.gremlin.Clone()
	}
	for _, p := range ftq.predicates {
		p.Gremlin(v)
	}
	if len(ftq.order) > 0 {
		v.Order()
		for _, p := range ftq.order {
			p.Gremlin(v)
		}
	}
	switch limit, offset := ftq.limit, ftq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt64)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := ftq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// FieldTypeQuery is the builder for group-by FieldType entities.
type FieldTypeGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FieldTypeGroupBy) Aggregate(fns ...Aggregate) *FieldTypeGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ftgb *FieldTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	switch ftgb.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftgb.sqlScan(ctx, v)
	case dialect.Neptune:
		return ftgb.gremlinScan(ctx, v)
	default:
		return errors.New("ftgb: unsupported dialect")
	}
}

// ScanX is like Scan, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ftgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ftgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ftgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ftgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ftgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftgb *FieldTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ftgb.sqlQuery().Query()
	if err := ftgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ftgb *FieldTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ftgb.sql
	columns := make([]string, 0, len(ftgb.fields)+len(ftgb.fns))
	columns = append(columns, ftgb.fields...)
	for _, fn := range ftgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(ftgb.fields...)
}

func (ftgb *FieldTypeGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ftgb.gremlinQuery().Query()
	if err := ftgb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ftgb.fields)+len(ftgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ftgb *FieldTypeGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ftgb.fns {
		name, tr := fn.Gremlin("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ftgb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ftgb.gremlin.Group().
		By(__.Values(ftgb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}
