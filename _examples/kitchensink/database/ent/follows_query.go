// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lrstanley/entrest/_examples/kitchensink/database/ent/follows"
	"github.com/lrstanley/entrest/_examples/kitchensink/database/ent/pet"
	"github.com/lrstanley/entrest/_examples/kitchensink/database/ent/predicate"
	"github.com/lrstanley/entrest/_examples/kitchensink/database/ent/user"
)

// FollowsQuery is the builder for querying Follows entities.
type FollowsQuery struct {
	config
	ctx        *QueryContext
	order      []follows.OrderOption
	inters     []Interceptor
	predicates []predicate.Follows
	withUser   *UserQuery
	withPet    *PetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FollowsQuery builder.
func (fq *FollowsQuery) Where(ps ...predicate.Follows) *FollowsQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FollowsQuery) Limit(limit int) *FollowsQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FollowsQuery) Offset(offset int) *FollowsQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FollowsQuery) Unique(unique bool) *FollowsQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FollowsQuery) Order(o ...follows.OrderOption) *FollowsQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryUser chains the current query on the "user" edge.
func (fq *FollowsQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(follows.Table, follows.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, follows.UserTable, follows.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPet chains the current query on the "pet" edge.
func (fq *FollowsQuery) QueryPet() *PetQuery {
	query := (&PetClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(follows.Table, follows.PetColumn, selector),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, follows.PetTable, follows.PetColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Follows entity from the query.
// Returns a *NotFoundError when no Follows was found.
func (fq *FollowsQuery) First(ctx context.Context) (*Follows, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{follows.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FollowsQuery) FirstX(ctx context.Context) *Follows {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single Follows entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Follows entity is found.
// Returns a *NotFoundError when no Follows entities are found.
func (fq *FollowsQuery) Only(ctx context.Context) (*Follows, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{follows.Label}
	default:
		return nil, &NotSingularError{follows.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FollowsQuery) OnlyX(ctx context.Context) *Follows {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of FollowsSlice.
func (fq *FollowsQuery) All(ctx context.Context) ([]*Follows, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Follows, *FollowsQuery]()
	return withInterceptors[[]*Follows](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FollowsQuery) AllX(ctx context.Context) []*Follows {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (fq *FollowsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FollowsQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FollowsQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FollowsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, "Exist")
	switch _, err := fq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FollowsQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FollowsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FollowsQuery) Clone() *FollowsQuery {
	if fq == nil {
		return nil
	}
	return &FollowsQuery{
		config:     fq.config,
		ctx:        fq.ctx.Clone(),
		order:      append([]follows.OrderOption{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.Follows{}, fq.predicates...),
		withUser:   fq.withUser.Clone(),
		withPet:    fq.withPet.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FollowsQuery) WithUser(opts ...func(*UserQuery)) *FollowsQuery {
	query := (&UserClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withUser = query
	return fq
}

// WithPet tells the query-builder to eager-load the nodes that are connected to
// the "pet" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FollowsQuery) WithPet(opts ...func(*PetQuery)) *FollowsQuery {
	query := (&PetClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withPet = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FollowedAt time.Time `json:"followed_at"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Follows.Query().
//		GroupBy(follows.FieldFollowedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FollowsQuery) GroupBy(field string, fields ...string) *FollowsGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FollowsGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = follows.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FollowedAt time.Time `json:"followed_at"`
//	}
//
//	client.Follows.Query().
//		Select(follows.FieldFollowedAt).
//		Scan(ctx, &v)
func (fq *FollowsQuery) Select(fields ...string) *FollowsSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FollowsSelect{FollowsQuery: fq}
	sbuild.label = follows.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FollowsSelect configured with the given aggregations.
func (fq *FollowsQuery) Aggregate(fns ...AggregateFunc) *FollowsSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FollowsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !follows.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FollowsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Follows, error) {
	var (
		nodes       = []*Follows{}
		_spec       = fq.querySpec()
		loadedTypes = [2]bool{
			fq.withUser != nil,
			fq.withPet != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Follows).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Follows{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withUser; query != nil {
		if err := fq.loadUser(ctx, query, nodes, nil,
			func(n *Follows, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withPet; query != nil {
		if err := fq.loadPet(ctx, query, nodes, nil,
			func(n *Follows, e *Pet) { n.Edges.Pet = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FollowsQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Follows, init func(*Follows), assign func(*Follows, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Follows)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *FollowsQuery) loadPet(ctx context.Context, query *PetQuery, nodes []*Follows, init func(*Follows), assign func(*Follows, *Pet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Follows)
	for i := range nodes {
		fk := nodes[i].PetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(pet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "pet_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fq *FollowsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FollowsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(follows.Table, follows.Columns, nil)
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if fq.withUser != nil {
			_spec.Node.AddColumnOnce(follows.FieldUserID)
		}
		if fq.withPet != nil {
			_spec.Node.AddColumnOnce(follows.FieldPetID)
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FollowsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(follows.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = follows.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FollowsGroupBy is the group-by builder for Follows entities.
type FollowsGroupBy struct {
	selector
	build *FollowsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FollowsGroupBy) Aggregate(fns ...AggregateFunc) *FollowsGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FollowsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FollowsQuery, *FollowsGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FollowsGroupBy) sqlScan(ctx context.Context, root *FollowsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FollowsSelect is the builder for selecting fields of Follows entities.
type FollowsSelect struct {
	*FollowsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FollowsSelect) Aggregate(fns ...AggregateFunc) *FollowsSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FollowsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FollowsQuery, *FollowsSelect](ctx, fs.FollowsQuery, fs, fs.inters, v)
}

func (fs *FollowsSelect) sqlScan(ctx context.Context, root *FollowsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
