// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"wedding/ent/event"
	"wedding/ent/eventrsvp"
	"wedding/ent/invitee"
	"wedding/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventRSVPQuery is the builder for querying EventRSVP entities.
type EventRSVPQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.EventRSVP
	// eager-loading edges.
	withEvent   *EventQuery
	withInvitee *InviteeQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EventRSVPQuery builder.
func (erq *EventRSVPQuery) Where(ps ...predicate.EventRSVP) *EventRSVPQuery {
	erq.predicates = append(erq.predicates, ps...)
	return erq
}

// Limit adds a limit step to the query.
func (erq *EventRSVPQuery) Limit(limit int) *EventRSVPQuery {
	erq.limit = &limit
	return erq
}

// Offset adds an offset step to the query.
func (erq *EventRSVPQuery) Offset(offset int) *EventRSVPQuery {
	erq.offset = &offset
	return erq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (erq *EventRSVPQuery) Unique(unique bool) *EventRSVPQuery {
	erq.unique = &unique
	return erq
}

// Order adds an order step to the query.
func (erq *EventRSVPQuery) Order(o ...OrderFunc) *EventRSVPQuery {
	erq.order = append(erq.order, o...)
	return erq
}

// QueryEvent chains the current query on the "event" edge.
func (erq *EventRSVPQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: erq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := erq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := erq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eventrsvp.Table, eventrsvp.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eventrsvp.EventTable, eventrsvp.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(erq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvitee chains the current query on the "invitee" edge.
func (erq *EventRSVPQuery) QueryInvitee() *InviteeQuery {
	query := &InviteeQuery{config: erq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := erq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := erq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eventrsvp.Table, eventrsvp.FieldID, selector),
			sqlgraph.To(invitee.Table, invitee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eventrsvp.InviteeTable, eventrsvp.InviteeColumn),
		)
		fromU = sqlgraph.SetNeighbors(erq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EventRSVP entity from the query.
// Returns a *NotFoundError when no EventRSVP was found.
func (erq *EventRSVPQuery) First(ctx context.Context) (*EventRSVP, error) {
	nodes, err := erq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{eventrsvp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (erq *EventRSVPQuery) FirstX(ctx context.Context) *EventRSVP {
	node, err := erq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EventRSVP ID from the query.
// Returns a *NotFoundError when no EventRSVP ID was found.
func (erq *EventRSVPQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = erq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{eventrsvp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (erq *EventRSVPQuery) FirstIDX(ctx context.Context) int {
	id, err := erq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EventRSVP entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one EventRSVP entity is not found.
// Returns a *NotFoundError when no EventRSVP entities are found.
func (erq *EventRSVPQuery) Only(ctx context.Context) (*EventRSVP, error) {
	nodes, err := erq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{eventrsvp.Label}
	default:
		return nil, &NotSingularError{eventrsvp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (erq *EventRSVPQuery) OnlyX(ctx context.Context) *EventRSVP {
	node, err := erq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EventRSVP ID in the query.
// Returns a *NotSingularError when exactly one EventRSVP ID is not found.
// Returns a *NotFoundError when no entities are found.
func (erq *EventRSVPQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = erq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = &NotSingularError{eventrsvp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (erq *EventRSVPQuery) OnlyIDX(ctx context.Context) int {
	id, err := erq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EventRSVPs.
func (erq *EventRSVPQuery) All(ctx context.Context) ([]*EventRSVP, error) {
	if err := erq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return erq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (erq *EventRSVPQuery) AllX(ctx context.Context) []*EventRSVP {
	nodes, err := erq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EventRSVP IDs.
func (erq *EventRSVPQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := erq.Select(eventrsvp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (erq *EventRSVPQuery) IDsX(ctx context.Context) []int {
	ids, err := erq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (erq *EventRSVPQuery) Count(ctx context.Context) (int, error) {
	if err := erq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return erq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (erq *EventRSVPQuery) CountX(ctx context.Context) int {
	count, err := erq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (erq *EventRSVPQuery) Exist(ctx context.Context) (bool, error) {
	if err := erq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return erq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (erq *EventRSVPQuery) ExistX(ctx context.Context) bool {
	exist, err := erq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EventRSVPQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (erq *EventRSVPQuery) Clone() *EventRSVPQuery {
	if erq == nil {
		return nil
	}
	return &EventRSVPQuery{
		config:      erq.config,
		limit:       erq.limit,
		offset:      erq.offset,
		order:       append([]OrderFunc{}, erq.order...),
		predicates:  append([]predicate.EventRSVP{}, erq.predicates...),
		withEvent:   erq.withEvent.Clone(),
		withInvitee: erq.withInvitee.Clone(),
		// clone intermediate query.
		sql:  erq.sql.Clone(),
		path: erq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (erq *EventRSVPQuery) WithEvent(opts ...func(*EventQuery)) *EventRSVPQuery {
	query := &EventQuery{config: erq.config}
	for _, opt := range opts {
		opt(query)
	}
	erq.withEvent = query
	return erq
}

// WithInvitee tells the query-builder to eager-load the nodes that are connected to
// the "invitee" edge. The optional arguments are used to configure the query builder of the edge.
func (erq *EventRSVPQuery) WithInvitee(opts ...func(*InviteeQuery)) *EventRSVPQuery {
	query := &InviteeQuery{config: erq.config}
	for _, opt := range opts {
		opt(query)
	}
	erq.withInvitee = query
	return erq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (erq *EventRSVPQuery) GroupBy(field string, fields ...string) *EventRSVPGroupBy {
	group := &EventRSVPGroupBy{config: erq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := erq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return erq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (erq *EventRSVPQuery) Select(field string, fields ...string) *EventRSVPSelect {
	erq.fields = append([]string{field}, fields...)
	return &EventRSVPSelect{EventRSVPQuery: erq}
}

func (erq *EventRSVPQuery) prepareQuery(ctx context.Context) error {
	for _, f := range erq.fields {
		if !eventrsvp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if erq.path != nil {
		prev, err := erq.path(ctx)
		if err != nil {
			return err
		}
		erq.sql = prev
	}
	return nil
}

func (erq *EventRSVPQuery) sqlAll(ctx context.Context) ([]*EventRSVP, error) {
	var (
		nodes       = []*EventRSVP{}
		withFKs     = erq.withFKs
		_spec       = erq.querySpec()
		loadedTypes = [2]bool{
			erq.withEvent != nil,
			erq.withInvitee != nil,
		}
	)
	if erq.withEvent != nil || erq.withInvitee != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, eventrsvp.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &EventRSVP{config: erq.config}
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
	if err := sqlgraph.QueryNodes(ctx, erq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := erq.withEvent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*EventRSVP)
		for i := range nodes {
			if nodes[i].event_rsvps == nil {
				continue
			}
			fk := *nodes[i].event_rsvps
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(event.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "event_rsvps" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Event = n
			}
		}
	}

	if query := erq.withInvitee; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*EventRSVP)
		for i := range nodes {
			if nodes[i].invitee_events == nil {
				continue
			}
			fk := *nodes[i].invitee_events
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(invitee.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "invitee_events" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Invitee = n
			}
		}
	}

	return nodes, nil
}

func (erq *EventRSVPQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := erq.querySpec()
	return sqlgraph.CountNodes(ctx, erq.driver, _spec)
}

func (erq *EventRSVPQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := erq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (erq *EventRSVPQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventrsvp.Table,
			Columns: eventrsvp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventrsvp.FieldID,
			},
		},
		From:   erq.sql,
		Unique: true,
	}
	if unique := erq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := erq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventrsvp.FieldID)
		for i := range fields {
			if fields[i] != eventrsvp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := erq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := erq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := erq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := erq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (erq *EventRSVPQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(erq.driver.Dialect())
	t1 := builder.Table(eventrsvp.Table)
	selector := builder.Select(t1.Columns(eventrsvp.Columns...)...).From(t1)
	if erq.sql != nil {
		selector = erq.sql
		selector.Select(selector.Columns(eventrsvp.Columns...)...)
	}
	for _, p := range erq.predicates {
		p(selector)
	}
	for _, p := range erq.order {
		p(selector)
	}
	if offset := erq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := erq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EventRSVPGroupBy is the group-by builder for EventRSVP entities.
type EventRSVPGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ergb *EventRSVPGroupBy) Aggregate(fns ...AggregateFunc) *EventRSVPGroupBy {
	ergb.fns = append(ergb.fns, fns...)
	return ergb
}

// Scan applies the group-by query and scans the result into the given value.
func (ergb *EventRSVPGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ergb.path(ctx)
	if err != nil {
		return err
	}
	ergb.sql = query
	return ergb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ergb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ergb *EventRSVPGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ergb.fields) > 1 {
		return nil, errors.New("ent: EventRSVPGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ergb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) StringsX(ctx context.Context) []string {
	v, err := ergb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ergb *EventRSVPGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ergb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = fmt.Errorf("ent: EventRSVPGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) StringX(ctx context.Context) string {
	v, err := ergb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ergb *EventRSVPGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ergb.fields) > 1 {
		return nil, errors.New("ent: EventRSVPGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ergb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) IntsX(ctx context.Context) []int {
	v, err := ergb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ergb *EventRSVPGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ergb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = fmt.Errorf("ent: EventRSVPGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) IntX(ctx context.Context) int {
	v, err := ergb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ergb *EventRSVPGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ergb.fields) > 1 {
		return nil, errors.New("ent: EventRSVPGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ergb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ergb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ergb *EventRSVPGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ergb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = fmt.Errorf("ent: EventRSVPGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ergb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ergb *EventRSVPGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ergb.fields) > 1 {
		return nil, errors.New("ent: EventRSVPGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ergb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ergb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ergb *EventRSVPGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ergb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = fmt.Errorf("ent: EventRSVPGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ergb *EventRSVPGroupBy) BoolX(ctx context.Context) bool {
	v, err := ergb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ergb *EventRSVPGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ergb.fields {
		if !eventrsvp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ergb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ergb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ergb *EventRSVPGroupBy) sqlQuery() *sql.Selector {
	selector := ergb.sql.Select()
	aggregation := make([]string, 0, len(ergb.fns))
	for _, fn := range ergb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ergb.fields)+len(ergb.fns))
		for _, f := range ergb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ergb.fields...)...)
}

// EventRSVPSelect is the builder for selecting fields of EventRSVP entities.
type EventRSVPSelect struct {
	*EventRSVPQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ers *EventRSVPSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ers.prepareQuery(ctx); err != nil {
		return err
	}
	ers.sql = ers.EventRSVPQuery.sqlQuery(ctx)
	return ers.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ers *EventRSVPSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ers.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ers *EventRSVPSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ers.fields) > 1 {
		return nil, errors.New("ent: EventRSVPSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ers.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ers *EventRSVPSelect) StringsX(ctx context.Context) []string {
	v, err := ers.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ers *EventRSVPSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ers.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = fmt.Errorf("ent: EventRSVPSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ers *EventRSVPSelect) StringX(ctx context.Context) string {
	v, err := ers.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ers *EventRSVPSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ers.fields) > 1 {
		return nil, errors.New("ent: EventRSVPSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ers.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ers *EventRSVPSelect) IntsX(ctx context.Context) []int {
	v, err := ers.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ers *EventRSVPSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ers.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = fmt.Errorf("ent: EventRSVPSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ers *EventRSVPSelect) IntX(ctx context.Context) int {
	v, err := ers.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ers *EventRSVPSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ers.fields) > 1 {
		return nil, errors.New("ent: EventRSVPSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ers.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ers *EventRSVPSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ers.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ers *EventRSVPSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ers.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = fmt.Errorf("ent: EventRSVPSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ers *EventRSVPSelect) Float64X(ctx context.Context) float64 {
	v, err := ers.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ers *EventRSVPSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ers.fields) > 1 {
		return nil, errors.New("ent: EventRSVPSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ers.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ers *EventRSVPSelect) BoolsX(ctx context.Context) []bool {
	v, err := ers.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ers *EventRSVPSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ers.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventrsvp.Label}
	default:
		err = fmt.Errorf("ent: EventRSVPSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ers *EventRSVPSelect) BoolX(ctx context.Context) bool {
	v, err := ers.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ers *EventRSVPSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ers.sqlQuery().Query()
	if err := ers.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ers *EventRSVPSelect) sqlQuery() sql.Querier {
	selector := ers.sql
	selector.Select(selector.Columns(ers.fields...)...)
	return selector
}
