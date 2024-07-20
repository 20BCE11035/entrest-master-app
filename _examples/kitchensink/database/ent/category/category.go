// Code generated by ent, DO NOT EDIT.

package category

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldReadonly holds the string denoting the readonly field in the database.
	FieldReadonly = "readonly"
	// FieldSkipInSpec holds the string denoting the skip_in_spec field in the database.
	FieldSkipInSpec = "skip_in_spec"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// PetsTable is the table that holds the pets relation/edge. The primary key declared below.
	PetsTable = "category_pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldReadonly,
	FieldSkipInSpec,
}

var (
	// PetsPrimaryKey and PetsColumn2 are the table columns denoting the
	// primary key for the pets relation (M2M).
	PetsPrimaryKey = []string{"category_id", "pet_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Category queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByReadonly orders the results by the readonly field.
func ByReadonly(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReadonly, opts...).ToFunc()
}

// BySkipInSpec orders the results by the skip_in_spec field.
func BySkipInSpec(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSkipInSpec, opts...).ToFunc()
}

// ByPetsCount orders the results by pets count.
func ByPetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPetsStep(), opts...)
	}
}

// ByPets orders the results by pets terms.
func ByPets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PetsTable, PetsPrimaryKey...),
	)
}
