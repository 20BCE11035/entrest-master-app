// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/entrest/_examples/kitchensink/database/ent/predicate"
	"github.com/lrstanley/entrest/_examples/kitchensink/database/ent/skipped"
)

// SkippedUpdate is the builder for updating Skipped entities.
type SkippedUpdate struct {
	config
	hooks    []Hook
	mutation *SkippedMutation
}

// Where appends a list predicates to the SkippedUpdate builder.
func (su *SkippedUpdate) Where(ps ...predicate.Skipped) *SkippedUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SkippedUpdate) SetName(s string) *SkippedUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SkippedUpdate) SetNillableName(s *string) *SkippedUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// Mutation returns the SkippedMutation object of the builder.
func (su *SkippedUpdate) Mutation() *SkippedMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SkippedUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SkippedUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SkippedUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SkippedUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SkippedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(skipped.Table, skipped.Columns, sqlgraph.NewFieldSpec(skipped.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(skipped.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skipped.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SkippedUpdateOne is the builder for updating a single Skipped entity.
type SkippedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkippedMutation
}

// SetName sets the "name" field.
func (suo *SkippedUpdateOne) SetName(s string) *SkippedUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SkippedUpdateOne) SetNillableName(s *string) *SkippedUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// Mutation returns the SkippedMutation object of the builder.
func (suo *SkippedUpdateOne) Mutation() *SkippedMutation {
	return suo.mutation
}

// Where appends a list predicates to the SkippedUpdate builder.
func (suo *SkippedUpdateOne) Where(ps ...predicate.Skipped) *SkippedUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SkippedUpdateOne) Select(field string, fields ...string) *SkippedUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Skipped entity.
func (suo *SkippedUpdateOne) Save(ctx context.Context) (*Skipped, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SkippedUpdateOne) SaveX(ctx context.Context) *Skipped {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SkippedUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SkippedUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SkippedUpdateOne) sqlSave(ctx context.Context) (_node *Skipped, err error) {
	_spec := sqlgraph.NewUpdateSpec(skipped.Table, skipped.Columns, sqlgraph.NewFieldSpec(skipped.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Skipped.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skipped.FieldID)
		for _, f := range fields {
			if !skipped.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skipped.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(skipped.FieldName, field.TypeString, value)
	}
	_node = &Skipped{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skipped.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
