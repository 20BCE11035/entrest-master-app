// Code generated by ent, DO NOT EDIT.

package settings

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lrstanley/entrest/_examples/kitchensink/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Settings {
	return predicate.Settings(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Settings {
	return predicate.Settings(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Settings {
	return predicate.Settings(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Settings {
	return predicate.Settings(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Settings {
	return predicate.Settings(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Settings {
	return predicate.Settings(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Settings {
	return predicate.Settings(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldUpdatedAt, v))
}

// GlobalBanner applies equality check predicate on the "global_banner" field. It's identical to GlobalBannerEQ.
func GlobalBanner(v string) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldGlobalBanner, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldLTE(FieldUpdatedAt, v))
}

// GlobalBannerEQ applies the EQ predicate on the "global_banner" field.
func GlobalBannerEQ(v string) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldGlobalBanner, v))
}

// GlobalBannerNEQ applies the NEQ predicate on the "global_banner" field.
func GlobalBannerNEQ(v string) predicate.Settings {
	return predicate.Settings(sql.FieldNEQ(FieldGlobalBanner, v))
}

// GlobalBannerIn applies the In predicate on the "global_banner" field.
func GlobalBannerIn(vs ...string) predicate.Settings {
	return predicate.Settings(sql.FieldIn(FieldGlobalBanner, vs...))
}

// GlobalBannerNotIn applies the NotIn predicate on the "global_banner" field.
func GlobalBannerNotIn(vs ...string) predicate.Settings {
	return predicate.Settings(sql.FieldNotIn(FieldGlobalBanner, vs...))
}

// GlobalBannerGT applies the GT predicate on the "global_banner" field.
func GlobalBannerGT(v string) predicate.Settings {
	return predicate.Settings(sql.FieldGT(FieldGlobalBanner, v))
}

// GlobalBannerGTE applies the GTE predicate on the "global_banner" field.
func GlobalBannerGTE(v string) predicate.Settings {
	return predicate.Settings(sql.FieldGTE(FieldGlobalBanner, v))
}

// GlobalBannerLT applies the LT predicate on the "global_banner" field.
func GlobalBannerLT(v string) predicate.Settings {
	return predicate.Settings(sql.FieldLT(FieldGlobalBanner, v))
}

// GlobalBannerLTE applies the LTE predicate on the "global_banner" field.
func GlobalBannerLTE(v string) predicate.Settings {
	return predicate.Settings(sql.FieldLTE(FieldGlobalBanner, v))
}

// GlobalBannerContains applies the Contains predicate on the "global_banner" field.
func GlobalBannerContains(v string) predicate.Settings {
	return predicate.Settings(sql.FieldContains(FieldGlobalBanner, v))
}

// GlobalBannerHasPrefix applies the HasPrefix predicate on the "global_banner" field.
func GlobalBannerHasPrefix(v string) predicate.Settings {
	return predicate.Settings(sql.FieldHasPrefix(FieldGlobalBanner, v))
}

// GlobalBannerHasSuffix applies the HasSuffix predicate on the "global_banner" field.
func GlobalBannerHasSuffix(v string) predicate.Settings {
	return predicate.Settings(sql.FieldHasSuffix(FieldGlobalBanner, v))
}

// GlobalBannerIsNil applies the IsNil predicate on the "global_banner" field.
func GlobalBannerIsNil() predicate.Settings {
	return predicate.Settings(sql.FieldIsNull(FieldGlobalBanner))
}

// GlobalBannerNotNil applies the NotNil predicate on the "global_banner" field.
func GlobalBannerNotNil() predicate.Settings {
	return predicate.Settings(sql.FieldNotNull(FieldGlobalBanner))
}

// GlobalBannerEqualFold applies the EqualFold predicate on the "global_banner" field.
func GlobalBannerEqualFold(v string) predicate.Settings {
	return predicate.Settings(sql.FieldEqualFold(FieldGlobalBanner, v))
}

// GlobalBannerContainsFold applies the ContainsFold predicate on the "global_banner" field.
func GlobalBannerContainsFold(v string) predicate.Settings {
	return predicate.Settings(sql.FieldContainsFold(FieldGlobalBanner, v))
}

// HasAdmins applies the HasEdge predicate on the "admins" edge.
func HasAdmins() predicate.Settings {
	return predicate.Settings(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AdminsTable, AdminsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdminsWith applies the HasEdge predicate on the "admins" edge with a given conditions (other predicates).
func HasAdminsWith(preds ...predicate.User) predicate.Settings {
	return predicate.Settings(func(s *sql.Selector) {
		step := newAdminsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Settings) predicate.Settings {
	return predicate.Settings(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Settings) predicate.Settings {
	return predicate.Settings(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Settings) predicate.Settings {
	return predicate.Settings(sql.NotPredicates(p))
}