// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"

	"github.com/facebook/ent/dialect/sql"
)

// Invitee is the model entity for the Invitee schema.
type Invitee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InviteeQuery when eager-loading is set.
	Edges                  InviteeEdges `json:"edges"`
	invitee_party_invitees *int
}

// InviteeEdges holds the relations/edges for other nodes in the graph.
type InviteeEdges struct {
	// Party holds the value of the party edge.
	Party *InviteeParty
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PartyOrErr returns the Party value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InviteeEdges) PartyOrErr() (*InviteeParty, error) {
	if e.loadedTypes[0] {
		if e.Party == nil {
			// The edge party was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: inviteeparty.Label}
		}
		return e.Party, nil
	}
	return nil, &NotLoadedError{edge: "party"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Invitee) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case invitee.FieldID:
			values[i] = &sql.NullInt64{}
		case invitee.FieldName:
			values[i] = &sql.NullString{}
		case invitee.ForeignKeys[0]: // invitee_party_invitees
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Invitee", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Invitee fields.
func (i *Invitee) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case invitee.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case invitee.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case invitee.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field invitee_party_invitees", value)
			} else if value.Valid {
				i.invitee_party_invitees = new(int)
				*i.invitee_party_invitees = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParty queries the "party" edge of the Invitee entity.
func (i *Invitee) QueryParty() *InviteePartyQuery {
	return (&InviteeClient{config: i.config}).QueryParty(i)
}

// Update returns a builder for updating this Invitee.
// Note that you need to call Invitee.Unwrap() before calling this method if this Invitee
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Invitee) Update() *InviteeUpdateOne {
	return (&InviteeClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Invitee entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Invitee) Unwrap() *Invitee {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Invitee is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Invitee) String() string {
	var builder strings.Builder
	builder.WriteString("Invitee(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", name=")
	builder.WriteString(i.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Invitees is a parsable slice of Invitee.
type Invitees []*Invitee

func (i Invitees) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}