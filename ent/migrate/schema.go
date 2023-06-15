// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "priority", Type: field.TypeInt},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "university", Type: field.TypeString},
		{Name: "organization_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_organizations_organization",
				Columns:    []*schema.Column{UsersColumns[2]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrganizationsTable,
		UsersTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = OrganizationsTable
}
