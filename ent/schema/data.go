package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Data holds the schema definition for the Data entity.
type Data struct {
	ent.Schema
}

// Fields of the Data.
func (Data) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Optional().Nillable(),
		field.String("text").
			Optional().Nillable(),
	}
}

// Edges of the Data.
func (Data) Edges() []ent.Edge {
	return nil
}
