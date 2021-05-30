// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:        "events",
		Columns:     EventsColumns,
		PrimaryKey:  []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EventRsvpsColumns holds the columns for the "event_rsvps" table.
	EventRsvpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event_rsvps", Type: field.TypeInt, Nullable: true},
		{Name: "invitee_events", Type: field.TypeInt, Nullable: true},
	}
	// EventRsvpsTable holds the schema information for the "event_rsvps" table.
	EventRsvpsTable = &schema.Table{
		Name:       "event_rsvps",
		Columns:    EventRsvpsColumns,
		PrimaryKey: []*schema.Column{EventRsvpsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "event_rsvps_events_rsvps",
				Columns:    []*schema.Column{EventRsvpsColumns[1]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "event_rsvps_invitees_events",
				Columns:    []*schema.Column{EventRsvpsColumns[2]},
				RefColumns: []*schema.Column{InviteesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InviteesColumns holds the columns for the "invitees" table.
	InviteesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_child", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "has_plus_one", Type: field.TypeBool, Default: false},
		{Name: "is_bridesmaid", Type: field.TypeBool, Default: false},
		{Name: "is_groomsman", Type: field.TypeBool, Default: false},
		{Name: "plus_one_name", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "address_line_1", Type: field.TypeString, Nullable: true},
		{Name: "address_line_2", Type: field.TypeString, Nullable: true},
		{Name: "address_city", Type: field.TypeString, Nullable: true},
		{Name: "address_state", Type: field.TypeString, Nullable: true},
		{Name: "address_postal_code", Type: field.TypeString, Nullable: true},
		{Name: "address_country", Type: field.TypeString, Nullable: true},
		{Name: "rsvp_response", Type: field.TypeBool, Nullable: true},
		{Name: "invitee_party_invitees", Type: field.TypeInt, Nullable: true},
	}
	// InviteesTable holds the schema information for the "invitees" table.
	InviteesTable = &schema.Table{
		Name:       "invitees",
		Columns:    InviteesColumns,
		PrimaryKey: []*schema.Column{InviteesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invitees_invitee_parties_invitees",
				Columns:    []*schema.Column{InviteesColumns[16]},
				RefColumns: []*schema.Column{InviteePartiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InviteePartiesColumns holds the columns for the "invitee_parties" table.
	InviteePartiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 10},
	}
	// InviteePartiesTable holds the schema information for the "invitee_parties" table.
	InviteePartiesTable = &schema.Table{
		Name:        "invitee_parties",
		Columns:     InviteePartiesColumns,
		PrimaryKey:  []*schema.Column{InviteePartiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		EventRsvpsTable,
		InviteesTable,
		InviteePartiesTable,
	}
)

func init() {
	EventRsvpsTable.ForeignKeys[0].RefTable = EventsTable
	EventRsvpsTable.ForeignKeys[1].RefTable = InviteesTable
	EventRsvpsTable.Annotation = &entsql.Annotation{
		Table: "event_rsvps",
	}
	InviteesTable.ForeignKeys[0].RefTable = InviteePartiesTable
}
