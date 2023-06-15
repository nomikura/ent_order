package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("university").
			Annotations(
				entgql.OrderField("UNIVERSITY"),
			),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organization", Organization.Type).
			Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.MultiOrder(),
	}
}
