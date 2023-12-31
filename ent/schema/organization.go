package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Int("priority").
			Annotations(
				entgql.OrderField("PRIORITY"),
			),
	}
}

func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.MultiOrder(),
	}
}
