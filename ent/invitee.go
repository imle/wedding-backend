// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"

	"entgo.io/ent/dialect/sql"
)

// Invitee is the model entity for the Invitee schema.
type Invitee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsChild holds the value of the "is_child" field.
	IsChild *bool `json:"is_child,omitempty"`
	// HasPlusOne holds the value of the "has_plus_one" field.
	HasPlusOne bool `json:"has_plus_one,omitempty"`
	// IsBridesmaid holds the value of the "is_bridesmaid" field.
	IsBridesmaid bool `json:"is_bridesmaid,omitempty"`
	// IsGroomsman holds the value of the "is_groomsman" field.
	IsGroomsman bool `json:"is_groomsman,omitempty"`
	// PlusOneName holds the value of the "plus_one_name" field.
	PlusOneName *string `json:"plus_one_name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone *string `json:"phone,omitempty"`
	// Email holds the value of the "email" field.
	Email *string `json:"email,omitempty"`
	// AddressLine1 holds the value of the "address_line_1" field.
	AddressLine1 *string `json:"address_line_1,omitempty"`
	// AddressLine2 holds the value of the "address_line_2" field.
	AddressLine2 *string `json:"address_line_2,omitempty"`
	// AddressCity holds the value of the "address_city" field.
	AddressCity *string `json:"address_city,omitempty"`
	// AddressState holds the value of the "address_state" field.
	AddressState *string `json:"address_state,omitempty"`
	// AddressPostalCode holds the value of the "address_postal_code" field.
	AddressPostalCode *string `json:"address_postal_code,omitempty"`
	// AddressCountry holds the value of the "address_country" field.
	AddressCountry *string `json:"address_country,omitempty"`
	// RsvpResponse holds the value of the "rsvp_response" field.
	RsvpResponse *bool `json:"rsvp_response,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InviteeQuery when eager-loading is set.
	Edges                  InviteeEdges `json:"edges"`
	invitee_party_invitees *int
}

// InviteeEdges holds the relations/edges for other nodes in the graph.
type InviteeEdges struct {
	// Events holds the value of the events edge.
	Events []*EventRSVP `json:"events,omitempty"`
	// Party holds the value of the party edge.
	Party *InviteeParty `json:"party,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e InviteeEdges) EventsOrErr() ([]*EventRSVP, error) {
	if e.loadedTypes[0] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// PartyOrErr returns the Party value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InviteeEdges) PartyOrErr() (*InviteeParty, error) {
	if e.loadedTypes[1] {
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
		case invitee.FieldIsChild, invitee.FieldHasPlusOne, invitee.FieldIsBridesmaid, invitee.FieldIsGroomsman, invitee.FieldRsvpResponse:
			values[i] = new(sql.NullBool)
		case invitee.FieldID:
			values[i] = new(sql.NullInt64)
		case invitee.FieldName, invitee.FieldPlusOneName, invitee.FieldPhone, invitee.FieldEmail, invitee.FieldAddressLine1, invitee.FieldAddressLine2, invitee.FieldAddressCity, invitee.FieldAddressState, invitee.FieldAddressPostalCode, invitee.FieldAddressCountry:
			values[i] = new(sql.NullString)
		case invitee.ForeignKeys[0]: // invitee_party_invitees
			values[i] = new(sql.NullInt64)
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
		case invitee.FieldIsChild:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_child", values[j])
			} else if value.Valid {
				i.IsChild = new(bool)
				*i.IsChild = value.Bool
			}
		case invitee.FieldHasPlusOne:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_plus_one", values[j])
			} else if value.Valid {
				i.HasPlusOne = value.Bool
			}
		case invitee.FieldIsBridesmaid:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_bridesmaid", values[j])
			} else if value.Valid {
				i.IsBridesmaid = value.Bool
			}
		case invitee.FieldIsGroomsman:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_groomsman", values[j])
			} else if value.Valid {
				i.IsGroomsman = value.Bool
			}
		case invitee.FieldPlusOneName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plus_one_name", values[j])
			} else if value.Valid {
				i.PlusOneName = new(string)
				*i.PlusOneName = value.String
			}
		case invitee.FieldPhone:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[j])
			} else if value.Valid {
				i.Phone = new(string)
				*i.Phone = value.String
			}
		case invitee.FieldEmail:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[j])
			} else if value.Valid {
				i.Email = new(string)
				*i.Email = value.String
			}
		case invitee.FieldAddressLine1:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line_1", values[j])
			} else if value.Valid {
				i.AddressLine1 = new(string)
				*i.AddressLine1 = value.String
			}
		case invitee.FieldAddressLine2:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line_2", values[j])
			} else if value.Valid {
				i.AddressLine2 = new(string)
				*i.AddressLine2 = value.String
			}
		case invitee.FieldAddressCity:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_city", values[j])
			} else if value.Valid {
				i.AddressCity = new(string)
				*i.AddressCity = value.String
			}
		case invitee.FieldAddressState:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_state", values[j])
			} else if value.Valid {
				i.AddressState = new(string)
				*i.AddressState = value.String
			}
		case invitee.FieldAddressPostalCode:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_postal_code", values[j])
			} else if value.Valid {
				i.AddressPostalCode = new(string)
				*i.AddressPostalCode = value.String
			}
		case invitee.FieldAddressCountry:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_country", values[j])
			} else if value.Valid {
				i.AddressCountry = new(string)
				*i.AddressCountry = value.String
			}
		case invitee.FieldRsvpResponse:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field rsvp_response", values[j])
			} else if value.Valid {
				i.RsvpResponse = new(bool)
				*i.RsvpResponse = value.Bool
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

// QueryEvents queries the "events" edge of the Invitee entity.
func (i *Invitee) QueryEvents() *EventRSVPQuery {
	return (&InviteeClient{config: i.config}).QueryEvents(i)
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
	if v := i.IsChild; v != nil {
		builder.WriteString(", is_child=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", has_plus_one=")
	builder.WriteString(fmt.Sprintf("%v", i.HasPlusOne))
	builder.WriteString(", is_bridesmaid=")
	builder.WriteString(fmt.Sprintf("%v", i.IsBridesmaid))
	builder.WriteString(", is_groomsman=")
	builder.WriteString(fmt.Sprintf("%v", i.IsGroomsman))
	if v := i.PlusOneName; v != nil {
		builder.WriteString(", plus_one_name=")
		builder.WriteString(*v)
	}
	if v := i.Phone; v != nil {
		builder.WriteString(", phone=")
		builder.WriteString(*v)
	}
	if v := i.Email; v != nil {
		builder.WriteString(", email=")
		builder.WriteString(*v)
	}
	if v := i.AddressLine1; v != nil {
		builder.WriteString(", address_line_1=")
		builder.WriteString(*v)
	}
	if v := i.AddressLine2; v != nil {
		builder.WriteString(", address_line_2=")
		builder.WriteString(*v)
	}
	if v := i.AddressCity; v != nil {
		builder.WriteString(", address_city=")
		builder.WriteString(*v)
	}
	if v := i.AddressState; v != nil {
		builder.WriteString(", address_state=")
		builder.WriteString(*v)
	}
	if v := i.AddressPostalCode; v != nil {
		builder.WriteString(", address_postal_code=")
		builder.WriteString(*v)
	}
	if v := i.AddressCountry; v != nil {
		builder.WriteString(", address_country=")
		builder.WriteString(*v)
	}
	if v := i.RsvpResponse; v != nil {
		builder.WriteString(", rsvp_response=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
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
