// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"
	"wedding/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// InviteePartyQuery is the builder for querying InviteeParty entities.
type InviteePartyQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.InviteeParty
	// eager-loading edges.
	withInvitees *InviteeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InviteePartyQuery builder.
func (ipq *InviteePartyQuery) Where(ps ...predicate.InviteeParty) *InviteePartyQuery {
	ipq.predicates = append(ipq.predicates, ps...)
	return ipq
}

// Limit adds a limit step to the query.
func (ipq *InviteePartyQuery) Limit(limit int) *InviteePartyQuery {
	ipq.limit = &limit
	return ipq
}

// Offset adds an offset step to the query.
func (ipq *InviteePartyQuery) Offset(offset int) *InviteePartyQuery {
	ipq.offset = &offset
	return ipq
}

// Order adds an order step to the query.
func (ipq *InviteePartyQuery) Order(o ...OrderFunc) *InviteePartyQuery {
	ipq.order = append(ipq.order, o...)
	return ipq
}

// QueryInvitees chains the current query on the "invitees" edge.
func (ipq *InviteePartyQuery) QueryInvitees() *InviteeQuery {
	query := &InviteeQuery{config: ipq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ipq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ipq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inviteeparty.Table, inviteeparty.FieldID, selector),
			sqlgraph.To(invitee.Table, invitee.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, inviteeparty.InviteesTable, inviteeparty.InviteesColumn),
		)
		fromU = sqlgraph.SetNeighbors(ipq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InviteeParty entity from the query.
// Returns a *NotFoundError when no InviteeParty was found.
func (ipq *InviteePartyQuery) First(ctx context.Context) (*InviteeParty, error) {
	nodes, err := ipq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{inviteeparty.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ipq *InviteePartyQuery) FirstX(ctx context.Context) *InviteeParty {
	node, err := ipq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InviteeParty ID from the query.
// Returns a *NotFoundError when no InviteeParty ID was found.
func (ipq *InviteePartyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ipq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{inviteeparty.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ipq *InviteePartyQuery) FirstIDX(ctx context.Context) int {
	id, err := ipq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InviteeParty entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one InviteeParty entity is not found.
// Returns a *NotFoundError when no InviteeParty entities are found.
func (ipq *InviteePartyQuery) Only(ctx context.Context) (*InviteeParty, error) {
	nodes, err := ipq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{inviteeparty.Label}
	default:
		return nil, &NotSingularError{inviteeparty.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ipq *InviteePartyQuery) OnlyX(ctx context.Context) *InviteeParty {
	node, err := ipq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InviteeParty ID in the query.
// Returns a *NotSingularError when exactly one InviteeParty ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ipq *InviteePartyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ipq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = &NotSingularError{inviteeparty.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ipq *InviteePartyQuery) OnlyIDX(ctx context.Context) int {
	id, err := ipq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InviteeParties.
func (ipq *InviteePartyQuery) All(ctx context.Context) ([]*InviteeParty, error) {
	if err := ipq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ipq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ipq *InviteePartyQuery) AllX(ctx context.Context) []*InviteeParty {
	nodes, err := ipq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InviteeParty IDs.
func (ipq *InviteePartyQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ipq.Select(inviteeparty.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ipq *InviteePartyQuery) IDsX(ctx context.Context) []int {
	ids, err := ipq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ipq *InviteePartyQuery) Count(ctx context.Context) (int, error) {
	if err := ipq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ipq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ipq *InviteePartyQuery) CountX(ctx context.Context) int {
	count, err := ipq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ipq *InviteePartyQuery) Exist(ctx context.Context) (bool, error) {
	if err := ipq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ipq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ipq *InviteePartyQuery) ExistX(ctx context.Context) bool {
	exist, err := ipq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InviteePartyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ipq *InviteePartyQuery) Clone() *InviteePartyQuery {
	if ipq == nil {
		return nil
	}
	return &InviteePartyQuery{
		config:       ipq.config,
		limit:        ipq.limit,
		offset:       ipq.offset,
		order:        append([]OrderFunc{}, ipq.order...),
		predicates:   append([]predicate.InviteeParty{}, ipq.predicates...),
		withInvitees: ipq.withInvitees.Clone(),
		// clone intermediate query.
		sql:  ipq.sql.Clone(),
		path: ipq.path,
	}
}

// WithInvitees tells the query-builder to eager-load the nodes that are connected to
// the "invitees" edge. The optional arguments are used to configure the query builder of the edge.
func (ipq *InviteePartyQuery) WithInvitees(opts ...func(*InviteeQuery)) *InviteePartyQuery {
	query := &InviteeQuery{config: ipq.config}
	for _, opt := range opts {
		opt(query)
	}
	ipq.withInvitees = query
	return ipq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InviteeParty.Query().
//		GroupBy(inviteeparty.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ipq *InviteePartyQuery) GroupBy(field string, fields ...string) *InviteePartyGroupBy {
	group := &InviteePartyGroupBy{config: ipq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ipq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ipq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.InviteeParty.Query().
//		Select(inviteeparty.FieldName).
//		Scan(ctx, &v)
//
func (ipq *InviteePartyQuery) Select(field string, fields ...string) *InviteePartySelect {
	ipq.fields = append([]string{field}, fields...)
	return &InviteePartySelect{InviteePartyQuery: ipq}
}

func (ipq *InviteePartyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ipq.fields {
		if !inviteeparty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ipq.path != nil {
		prev, err := ipq.path(ctx)
		if err != nil {
			return err
		}
		ipq.sql = prev
	}
	return nil
}

func (ipq *InviteePartyQuery) sqlAll(ctx context.Context) ([]*InviteeParty, error) {
	var (
		nodes       = []*InviteeParty{}
		_spec       = ipq.querySpec()
		loadedTypes = [1]bool{
			ipq.withInvitees != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &InviteeParty{config: ipq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ipq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ipq.withInvitees; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*InviteeParty)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Invitees = []*Invitee{}
		}
		query.withFKs = true
		query.Where(predicate.Invitee(func(s *sql.Selector) {
			s.Where(sql.InValues(inviteeparty.InviteesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.invitee_party_invitees
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "invitee_party_invitees" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "invitee_party_invitees" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Invitees = append(node.Edges.Invitees, n)
		}
	}

	return nodes, nil
}

func (ipq *InviteePartyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ipq.querySpec()
	return sqlgraph.CountNodes(ctx, ipq.driver, _spec)
}

func (ipq *InviteePartyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ipq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ipq *InviteePartyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   inviteeparty.Table,
			Columns: inviteeparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inviteeparty.FieldID,
			},
		},
		From:   ipq.sql,
		Unique: true,
	}
	if fields := ipq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inviteeparty.FieldID)
		for i := range fields {
			if fields[i] != inviteeparty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ipq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ipq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ipq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ipq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, inviteeparty.ValidColumn)
			}
		}
	}
	return _spec
}

func (ipq *InviteePartyQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ipq.driver.Dialect())
	t1 := builder.Table(inviteeparty.Table)
	selector := builder.Select(t1.Columns(inviteeparty.Columns...)...).From(t1)
	if ipq.sql != nil {
		selector = ipq.sql
		selector.Select(selector.Columns(inviteeparty.Columns...)...)
	}
	for _, p := range ipq.predicates {
		p(selector)
	}
	for _, p := range ipq.order {
		p(selector, inviteeparty.ValidColumn)
	}
	if offset := ipq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ipq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InviteePartyGroupBy is the group-by builder for InviteeParty entities.
type InviteePartyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ipgb *InviteePartyGroupBy) Aggregate(fns ...AggregateFunc) *InviteePartyGroupBy {
	ipgb.fns = append(ipgb.fns, fns...)
	return ipgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ipgb *InviteePartyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ipgb.path(ctx)
	if err != nil {
		return err
	}
	ipgb.sql = query
	return ipgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ipgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ipgb *InviteePartyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ipgb.fields) > 1 {
		return nil, errors.New("ent: InviteePartyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ipgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) StringsX(ctx context.Context) []string {
	v, err := ipgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ipgb *InviteePartyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ipgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = fmt.Errorf("ent: InviteePartyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) StringX(ctx context.Context) string {
	v, err := ipgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ipgb *InviteePartyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ipgb.fields) > 1 {
		return nil, errors.New("ent: InviteePartyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ipgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) IntsX(ctx context.Context) []int {
	v, err := ipgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ipgb *InviteePartyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ipgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = fmt.Errorf("ent: InviteePartyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) IntX(ctx context.Context) int {
	v, err := ipgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ipgb *InviteePartyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ipgb.fields) > 1 {
		return nil, errors.New("ent: InviteePartyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ipgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ipgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ipgb *InviteePartyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ipgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = fmt.Errorf("ent: InviteePartyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ipgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ipgb *InviteePartyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ipgb.fields) > 1 {
		return nil, errors.New("ent: InviteePartyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ipgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ipgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ipgb *InviteePartyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ipgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = fmt.Errorf("ent: InviteePartyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ipgb *InviteePartyGroupBy) BoolX(ctx context.Context) bool {
	v, err := ipgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ipgb *InviteePartyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ipgb.fields {
		if !inviteeparty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ipgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ipgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ipgb *InviteePartyGroupBy) sqlQuery() *sql.Selector {
	selector := ipgb.sql
	columns := make([]string, 0, len(ipgb.fields)+len(ipgb.fns))
	columns = append(columns, ipgb.fields...)
	for _, fn := range ipgb.fns {
		columns = append(columns, fn(selector, inviteeparty.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ipgb.fields...)
}

// InviteePartySelect is the builder for selecting fields of InviteeParty entities.
type InviteePartySelect struct {
	*InviteePartyQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ips *InviteePartySelect) Scan(ctx context.Context, v interface{}) error {
	if err := ips.prepareQuery(ctx); err != nil {
		return err
	}
	ips.sql = ips.InviteePartyQuery.sqlQuery()
	return ips.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ips *InviteePartySelect) ScanX(ctx context.Context, v interface{}) {
	if err := ips.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ips *InviteePartySelect) Strings(ctx context.Context) ([]string, error) {
	if len(ips.fields) > 1 {
		return nil, errors.New("ent: InviteePartySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ips.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ips *InviteePartySelect) StringsX(ctx context.Context) []string {
	v, err := ips.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ips *InviteePartySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ips.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = fmt.Errorf("ent: InviteePartySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ips *InviteePartySelect) StringX(ctx context.Context) string {
	v, err := ips.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ips *InviteePartySelect) Ints(ctx context.Context) ([]int, error) {
	if len(ips.fields) > 1 {
		return nil, errors.New("ent: InviteePartySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ips.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ips *InviteePartySelect) IntsX(ctx context.Context) []int {
	v, err := ips.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ips *InviteePartySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ips.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = fmt.Errorf("ent: InviteePartySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ips *InviteePartySelect) IntX(ctx context.Context) int {
	v, err := ips.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ips *InviteePartySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ips.fields) > 1 {
		return nil, errors.New("ent: InviteePartySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ips.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ips *InviteePartySelect) Float64sX(ctx context.Context) []float64 {
	v, err := ips.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ips *InviteePartySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ips.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = fmt.Errorf("ent: InviteePartySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ips *InviteePartySelect) Float64X(ctx context.Context) float64 {
	v, err := ips.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ips *InviteePartySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ips.fields) > 1 {
		return nil, errors.New("ent: InviteePartySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ips.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ips *InviteePartySelect) BoolsX(ctx context.Context) []bool {
	v, err := ips.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ips *InviteePartySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ips.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inviteeparty.Label}
	default:
		err = fmt.Errorf("ent: InviteePartySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ips *InviteePartySelect) BoolX(ctx context.Context) bool {
	v, err := ips.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ips *InviteePartySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ips.sqlQuery().Query()
	if err := ips.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ips *InviteePartySelect) sqlQuery() sql.Querier {
	selector := ips.sql
	selector.Select(selector.Columns(ips.fields...)...)
	return selector
}