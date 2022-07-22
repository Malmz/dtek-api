// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/lunchmenuitem"
	"github.com/dtekcth/dtek-api/ent/predicate"
	"github.com/dtekcth/dtek-api/ent/resturant"
)

// LunchMenuUpdate is the builder for updating LunchMenu entities.
type LunchMenuUpdate struct {
	config
	hooks    []Hook
	mutation *LunchMenuMutation
}

// Where appends a list predicates to the LunchMenuUpdate builder.
func (lmu *LunchMenuUpdate) Where(ps ...predicate.LunchMenu) *LunchMenuUpdate {
	lmu.mutation.Where(ps...)
	return lmu
}

// SetDate sets the "date" field.
func (lmu *LunchMenuUpdate) SetDate(t time.Time) *LunchMenuUpdate {
	lmu.mutation.SetDate(t)
	return lmu
}

// AddItemIDs adds the "items" edge to the LunchMenuItem entity by IDs.
func (lmu *LunchMenuUpdate) AddItemIDs(ids ...int) *LunchMenuUpdate {
	lmu.mutation.AddItemIDs(ids...)
	return lmu
}

// AddItems adds the "items" edges to the LunchMenuItem entity.
func (lmu *LunchMenuUpdate) AddItems(l ...*LunchMenuItem) *LunchMenuUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmu.AddItemIDs(ids...)
}

// SetResturantID sets the "resturant" edge to the Resturant entity by ID.
func (lmu *LunchMenuUpdate) SetResturantID(id int) *LunchMenuUpdate {
	lmu.mutation.SetResturantID(id)
	return lmu
}

// SetResturant sets the "resturant" edge to the Resturant entity.
func (lmu *LunchMenuUpdate) SetResturant(r *Resturant) *LunchMenuUpdate {
	return lmu.SetResturantID(r.ID)
}

// Mutation returns the LunchMenuMutation object of the builder.
func (lmu *LunchMenuUpdate) Mutation() *LunchMenuMutation {
	return lmu.mutation
}

// ClearItems clears all "items" edges to the LunchMenuItem entity.
func (lmu *LunchMenuUpdate) ClearItems() *LunchMenuUpdate {
	lmu.mutation.ClearItems()
	return lmu
}

// RemoveItemIDs removes the "items" edge to LunchMenuItem entities by IDs.
func (lmu *LunchMenuUpdate) RemoveItemIDs(ids ...int) *LunchMenuUpdate {
	lmu.mutation.RemoveItemIDs(ids...)
	return lmu
}

// RemoveItems removes "items" edges to LunchMenuItem entities.
func (lmu *LunchMenuUpdate) RemoveItems(l ...*LunchMenuItem) *LunchMenuUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmu.RemoveItemIDs(ids...)
}

// ClearResturant clears the "resturant" edge to the Resturant entity.
func (lmu *LunchMenuUpdate) ClearResturant() *LunchMenuUpdate {
	lmu.mutation.ClearResturant()
	return lmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lmu *LunchMenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lmu.hooks) == 0 {
		if err = lmu.check(); err != nil {
			return 0, err
		}
		affected, err = lmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LunchMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmu.check(); err != nil {
				return 0, err
			}
			lmu.mutation = mutation
			affected, err = lmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lmu.hooks) - 1; i >= 0; i-- {
			if lmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lmu *LunchMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := lmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lmu *LunchMenuUpdate) Exec(ctx context.Context) error {
	_, err := lmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmu *LunchMenuUpdate) ExecX(ctx context.Context) {
	if err := lmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmu *LunchMenuUpdate) check() error {
	if _, ok := lmu.mutation.ResturantID(); lmu.mutation.ResturantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LunchMenu.resturant"`)
	}
	return nil
}

func (lmu *LunchMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lunchmenu.Table,
			Columns: lunchmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchmenu.FieldID,
			},
		},
	}
	if ps := lmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lunchmenu.FieldDate,
		})
	}
	if lmu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lunchmenu.ItemsTable,
			Columns: []string{lunchmenu.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenuitem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !lmu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lunchmenu.ItemsTable,
			Columns: []string{lunchmenu.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenuitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lunchmenu.ItemsTable,
			Columns: []string{lunchmenu.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenuitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmu.mutation.ResturantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchmenu.ResturantTable,
			Columns: []string{lunchmenu.ResturantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resturant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.ResturantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchmenu.ResturantTable,
			Columns: []string{lunchmenu.ResturantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resturant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lunchmenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LunchMenuUpdateOne is the builder for updating a single LunchMenu entity.
type LunchMenuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LunchMenuMutation
}

// SetDate sets the "date" field.
func (lmuo *LunchMenuUpdateOne) SetDate(t time.Time) *LunchMenuUpdateOne {
	lmuo.mutation.SetDate(t)
	return lmuo
}

// AddItemIDs adds the "items" edge to the LunchMenuItem entity by IDs.
func (lmuo *LunchMenuUpdateOne) AddItemIDs(ids ...int) *LunchMenuUpdateOne {
	lmuo.mutation.AddItemIDs(ids...)
	return lmuo
}

// AddItems adds the "items" edges to the LunchMenuItem entity.
func (lmuo *LunchMenuUpdateOne) AddItems(l ...*LunchMenuItem) *LunchMenuUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmuo.AddItemIDs(ids...)
}

// SetResturantID sets the "resturant" edge to the Resturant entity by ID.
func (lmuo *LunchMenuUpdateOne) SetResturantID(id int) *LunchMenuUpdateOne {
	lmuo.mutation.SetResturantID(id)
	return lmuo
}

// SetResturant sets the "resturant" edge to the Resturant entity.
func (lmuo *LunchMenuUpdateOne) SetResturant(r *Resturant) *LunchMenuUpdateOne {
	return lmuo.SetResturantID(r.ID)
}

// Mutation returns the LunchMenuMutation object of the builder.
func (lmuo *LunchMenuUpdateOne) Mutation() *LunchMenuMutation {
	return lmuo.mutation
}

// ClearItems clears all "items" edges to the LunchMenuItem entity.
func (lmuo *LunchMenuUpdateOne) ClearItems() *LunchMenuUpdateOne {
	lmuo.mutation.ClearItems()
	return lmuo
}

// RemoveItemIDs removes the "items" edge to LunchMenuItem entities by IDs.
func (lmuo *LunchMenuUpdateOne) RemoveItemIDs(ids ...int) *LunchMenuUpdateOne {
	lmuo.mutation.RemoveItemIDs(ids...)
	return lmuo
}

// RemoveItems removes "items" edges to LunchMenuItem entities.
func (lmuo *LunchMenuUpdateOne) RemoveItems(l ...*LunchMenuItem) *LunchMenuUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmuo.RemoveItemIDs(ids...)
}

// ClearResturant clears the "resturant" edge to the Resturant entity.
func (lmuo *LunchMenuUpdateOne) ClearResturant() *LunchMenuUpdateOne {
	lmuo.mutation.ClearResturant()
	return lmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lmuo *LunchMenuUpdateOne) Select(field string, fields ...string) *LunchMenuUpdateOne {
	lmuo.fields = append([]string{field}, fields...)
	return lmuo
}

// Save executes the query and returns the updated LunchMenu entity.
func (lmuo *LunchMenuUpdateOne) Save(ctx context.Context) (*LunchMenu, error) {
	var (
		err  error
		node *LunchMenu
	)
	if len(lmuo.hooks) == 0 {
		if err = lmuo.check(); err != nil {
			return nil, err
		}
		node, err = lmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LunchMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmuo.check(); err != nil {
				return nil, err
			}
			lmuo.mutation = mutation
			node, err = lmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lmuo.hooks) - 1; i >= 0; i-- {
			if lmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LunchMenu)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LunchMenuMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lmuo *LunchMenuUpdateOne) SaveX(ctx context.Context) *LunchMenu {
	node, err := lmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lmuo *LunchMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := lmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmuo *LunchMenuUpdateOne) ExecX(ctx context.Context) {
	if err := lmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmuo *LunchMenuUpdateOne) check() error {
	if _, ok := lmuo.mutation.ResturantID(); lmuo.mutation.ResturantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LunchMenu.resturant"`)
	}
	return nil
}

func (lmuo *LunchMenuUpdateOne) sqlSave(ctx context.Context) (_node *LunchMenu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lunchmenu.Table,
			Columns: lunchmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchmenu.FieldID,
			},
		},
	}
	id, ok := lmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LunchMenu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lunchmenu.FieldID)
		for _, f := range fields {
			if !lunchmenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lunchmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmuo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lunchmenu.FieldDate,
		})
	}
	if lmuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lunchmenu.ItemsTable,
			Columns: []string{lunchmenu.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenuitem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !lmuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lunchmenu.ItemsTable,
			Columns: []string{lunchmenu.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenuitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lunchmenu.ItemsTable,
			Columns: []string{lunchmenu.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenuitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmuo.mutation.ResturantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchmenu.ResturantTable,
			Columns: []string{lunchmenu.ResturantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resturant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.ResturantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchmenu.ResturantTable,
			Columns: []string{lunchmenu.ResturantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resturant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LunchMenu{config: lmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lunchmenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
