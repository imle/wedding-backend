// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// InviteeCreate is the builder for creating a Invitee entity.
type InviteeCreate struct {
	config
	mutation *InviteeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *InviteeCreate) SetName(s string) *InviteeCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetPartyID sets the "party" edge to the InviteeParty entity by ID.
func (ic *InviteeCreate) SetPartyID(id int) *InviteeCreate {
	ic.mutation.SetPartyID(id)
	return ic
}

// SetNillablePartyID sets the "party" edge to the InviteeParty entity by ID if the given value is not nil.
func (ic *InviteeCreate) SetNillablePartyID(id *int) *InviteeCreate {
	if id != nil {
		ic = ic.SetPartyID(*id)
	}
	return ic
}

// SetParty sets the "party" edge to the InviteeParty entity.
func (ic *InviteeCreate) SetParty(i *InviteeParty) *InviteeCreate {
	return ic.SetPartyID(i.ID)
}

// Mutation returns the InviteeMutation object of the builder.
func (ic *InviteeCreate) Mutation() *InviteeMutation {
	return ic.mutation
}

// Save creates the Invitee in the database.
func (ic *InviteeCreate) Save(ctx context.Context) (*Invitee, error) {
	var (
		err  error
		node *Invitee
	)
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			node, err = ic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InviteeCreate) SaveX(ctx context.Context) *Invitee {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ic *InviteeCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := ic.mutation.Name(); ok {
		if err := invitee.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (ic *InviteeCreate) sqlSave(ctx context.Context) (*Invitee, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *InviteeCreate) createSpec() (*Invitee, *sqlgraph.CreateSpec) {
	var (
		_node = &Invitee{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: invitee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invitee.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldName,
		})
		_node.Name = value
	}
	if nodes := ic.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitee.PartyTable,
			Columns: []string{invitee.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inviteeparty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InviteeCreateBulk is the builder for creating many Invitee entities in bulk.
type InviteeCreateBulk struct {
	config
	builders []*InviteeCreate
}

// Save creates the Invitee entities in the database.
func (icb *InviteeCreateBulk) Save(ctx context.Context) ([]*Invitee, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invitee, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InviteeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InviteeCreateBulk) SaveX(ctx context.Context) []*Invitee {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}