// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// InviteesColumns holds the columns for the "invitees" table.
	InviteesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "invitee_party_invitees", Type: field.TypeInt, Nullable: true},
	}
	// InviteesTable holds the schema information for the "invitees" table.
	InviteesTable = &schema.Table{
		Name:       "invitees",
		Columns:    InviteesColumns,
		PrimaryKey: []*schema.Column{InviteesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "invitees_invitee_parties_invitees",
				Columns: []*schema.Column{InviteesColumns[2]},

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
		InviteesTable,
		InviteePartiesTable,
	}
)

func init() {
	InviteesTable.ForeignKeys[0].RefTable = InviteePartiesTable
}