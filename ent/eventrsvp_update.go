// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"wedding/ent/event"
	"wedding/ent/eventrsvp"
	"wedding/ent/invitee"
	"wedding/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventRSVPUpdate is the builder for updating EventRSVP entities.
type EventRSVPUpdate struct {
	config
	hooks    []Hook
	mutation *EventRSVPMutation
}

// Where adds a new predicate for the EventRSVPUpdate builder.
func (eru *EventRSVPUpdate) Where(ps ...predicate.EventRSVP) *EventRSVPUpdate {
	eru.mutation.predicates = append(eru.mutation.predicates, ps...)
	return eru
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (eru *EventRSVPUpdate) SetEventID(id int) *EventRSVPUpdate {
	eru.mutation.SetEventID(id)
	return eru
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (eru *EventRSVPUpdate) SetNillableEventID(id *int) *EventRSVPUpdate {
	if id != nil {
		eru = eru.SetEventID(*id)
	}
	return eru
}

// SetEvent sets the "event" edge to the Event entity.
func (eru *EventRSVPUpdate) SetEvent(e *Event) *EventRSVPUpdate {
	return eru.SetEventID(e.ID)
}

// SetInviteeID sets the "invitee" edge to the Invitee entity by ID.
func (eru *EventRSVPUpdate) SetInviteeID(id int) *EventRSVPUpdate {
	eru.mutation.SetInviteeID(id)
	return eru
}

// SetNillableInviteeID sets the "invitee" edge to the Invitee entity by ID if the given value is not nil.
func (eru *EventRSVPUpdate) SetNillableInviteeID(id *int) *EventRSVPUpdate {
	if id != nil {
		eru = eru.SetInviteeID(*id)
	}
	return eru
}

// SetInvitee sets the "invitee" edge to the Invitee entity.
func (eru *EventRSVPUpdate) SetInvitee(i *Invitee) *EventRSVPUpdate {
	return eru.SetInviteeID(i.ID)
}

// Mutation returns the EventRSVPMutation object of the builder.
func (eru *EventRSVPUpdate) Mutation() *EventRSVPMutation {
	return eru.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (eru *EventRSVPUpdate) ClearEvent() *EventRSVPUpdate {
	eru.mutation.ClearEvent()
	return eru
}

// ClearInvitee clears the "invitee" edge to the Invitee entity.
func (eru *EventRSVPUpdate) ClearInvitee() *EventRSVPUpdate {
	eru.mutation.ClearInvitee()
	return eru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eru *EventRSVPUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eru.hooks) == 0 {
		affected, err = eru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventRSVPMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eru.mutation = mutation
			affected, err = eru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eru.hooks) - 1; i >= 0; i-- {
			mut = eru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eru *EventRSVPUpdate) SaveX(ctx context.Context) int {
	affected, err := eru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eru *EventRSVPUpdate) Exec(ctx context.Context) error {
	_, err := eru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eru *EventRSVPUpdate) ExecX(ctx context.Context) {
	if err := eru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eru *EventRSVPUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventrsvp.Table,
			Columns: eventrsvp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventrsvp.FieldID,
			},
		},
	}
	if ps := eru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if eru.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventrsvp.EventTable,
			Columns: []string{eventrsvp.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eru.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventrsvp.EventTable,
			Columns: []string{eventrsvp.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eru.mutation.InviteeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventrsvp.InviteeTable,
			Columns: []string{eventrsvp.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: invitee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eru.mutation.InviteeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventrsvp.InviteeTable,
			Columns: []string{eventrsvp.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: invitee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventrsvp.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EventRSVPUpdateOne is the builder for updating a single EventRSVP entity.
type EventRSVPUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventRSVPMutation
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (eruo *EventRSVPUpdateOne) SetEventID(id int) *EventRSVPUpdateOne {
	eruo.mutation.SetEventID(id)
	return eruo
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (eruo *EventRSVPUpdateOne) SetNillableEventID(id *int) *EventRSVPUpdateOne {
	if id != nil {
		eruo = eruo.SetEventID(*id)
	}
	return eruo
}

// SetEvent sets the "event" edge to the Event entity.
func (eruo *EventRSVPUpdateOne) SetEvent(e *Event) *EventRSVPUpdateOne {
	return eruo.SetEventID(e.ID)
}

// SetInviteeID sets the "invitee" edge to the Invitee entity by ID.
func (eruo *EventRSVPUpdateOne) SetInviteeID(id int) *EventRSVPUpdateOne {
	eruo.mutation.SetInviteeID(id)
	return eruo
}

// SetNillableInviteeID sets the "invitee" edge to the Invitee entity by ID if the given value is not nil.
func (eruo *EventRSVPUpdateOne) SetNillableInviteeID(id *int) *EventRSVPUpdateOne {
	if id != nil {
		eruo = eruo.SetInviteeID(*id)
	}
	return eruo
}

// SetInvitee sets the "invitee" edge to the Invitee entity.
func (eruo *EventRSVPUpdateOne) SetInvitee(i *Invitee) *EventRSVPUpdateOne {
	return eruo.SetInviteeID(i.ID)
}

// Mutation returns the EventRSVPMutation object of the builder.
func (eruo *EventRSVPUpdateOne) Mutation() *EventRSVPMutation {
	return eruo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (eruo *EventRSVPUpdateOne) ClearEvent() *EventRSVPUpdateOne {
	eruo.mutation.ClearEvent()
	return eruo
}

// ClearInvitee clears the "invitee" edge to the Invitee entity.
func (eruo *EventRSVPUpdateOne) ClearInvitee() *EventRSVPUpdateOne {
	eruo.mutation.ClearInvitee()
	return eruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eruo *EventRSVPUpdateOne) Select(field string, fields ...string) *EventRSVPUpdateOne {
	eruo.fields = append([]string{field}, fields...)
	return eruo
}

// Save executes the query and returns the updated EventRSVP entity.
func (eruo *EventRSVPUpdateOne) Save(ctx context.Context) (*EventRSVP, error) {
	var (
		err  error
		node *EventRSVP
	)
	if len(eruo.hooks) == 0 {
		node, err = eruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventRSVPMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eruo.mutation = mutation
			node, err = eruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eruo.hooks) - 1; i >= 0; i-- {
			mut = eruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (eruo *EventRSVPUpdateOne) SaveX(ctx context.Context) *EventRSVP {
	node, err := eruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eruo *EventRSVPUpdateOne) Exec(ctx context.Context) error {
	_, err := eruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eruo *EventRSVPUpdateOne) ExecX(ctx context.Context) {
	if err := eruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eruo *EventRSVPUpdateOne) sqlSave(ctx context.Context) (_node *EventRSVP, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventrsvp.Table,
			Columns: eventrsvp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventrsvp.FieldID,
			},
		},
	}
	id, ok := eruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing EventRSVP.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := eruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventrsvp.FieldID)
		for _, f := range fields {
			if !eventrsvp.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eventrsvp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if eruo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventrsvp.EventTable,
			Columns: []string{eventrsvp.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eruo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventrsvp.EventTable,
			Columns: []string{eventrsvp.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eruo.mutation.InviteeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventrsvp.InviteeTable,
			Columns: []string{eventrsvp.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: invitee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eruo.mutation.InviteeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventrsvp.InviteeTable,
			Columns: []string{eventrsvp.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: invitee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EventRSVP{config: eruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventrsvp.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
