// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/lunchmenuitem"
	"github.com/dtekcth/dtek-api/ent/resturant"
)

// LunchMenuCreate is the builder for creating a LunchMenu entity.
type LunchMenuCreate struct {
	config
	mutation *LunchMenuMutation
	hooks    []Hook
}

// SetDate sets the "date" field.
func (lmc *LunchMenuCreate) SetDate(t time.Time) *LunchMenuCreate {
	lmc.mutation.SetDate(t)
	return lmc
}

// AddItemIDs adds the "items" edge to the LunchMenuItem entity by IDs.
func (lmc *LunchMenuCreate) AddItemIDs(ids ...int) *LunchMenuCreate {
	lmc.mutation.AddItemIDs(ids...)
	return lmc
}

// AddItems adds the "items" edges to the LunchMenuItem entity.
func (lmc *LunchMenuCreate) AddItems(l ...*LunchMenuItem) *LunchMenuCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmc.AddItemIDs(ids...)
}

// SetResturantID sets the "resturant" edge to the Resturant entity by ID.
func (lmc *LunchMenuCreate) SetResturantID(id int) *LunchMenuCreate {
	lmc.mutation.SetResturantID(id)
	return lmc
}

// SetResturant sets the "resturant" edge to the Resturant entity.
func (lmc *LunchMenuCreate) SetResturant(r *Resturant) *LunchMenuCreate {
	return lmc.SetResturantID(r.ID)
}

// Mutation returns the LunchMenuMutation object of the builder.
func (lmc *LunchMenuCreate) Mutation() *LunchMenuMutation {
	return lmc.mutation
}

// Save creates the LunchMenu in the database.
func (lmc *LunchMenuCreate) Save(ctx context.Context) (*LunchMenu, error) {
	var (
		err  error
		node *LunchMenu
	)
	if len(lmc.hooks) == 0 {
		if err = lmc.check(); err != nil {
			return nil, err
		}
		node, err = lmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LunchMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmc.check(); err != nil {
				return nil, err
			}
			lmc.mutation = mutation
			if node, err = lmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lmc.hooks) - 1; i >= 0; i-- {
			if lmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lmc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (lmc *LunchMenuCreate) SaveX(ctx context.Context) *LunchMenu {
	v, err := lmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lmc *LunchMenuCreate) Exec(ctx context.Context) error {
	_, err := lmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmc *LunchMenuCreate) ExecX(ctx context.Context) {
	if err := lmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmc *LunchMenuCreate) check() error {
	if _, ok := lmc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "LunchMenu.date"`)}
	}
	if _, ok := lmc.mutation.ResturantID(); !ok {
		return &ValidationError{Name: "resturant", err: errors.New(`ent: missing required edge "LunchMenu.resturant"`)}
	}
	return nil
}

func (lmc *LunchMenuCreate) sqlSave(ctx context.Context) (*LunchMenu, error) {
	_node, _spec := lmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lmc *LunchMenuCreate) createSpec() (*LunchMenu, *sqlgraph.CreateSpec) {
	var (
		_node = &LunchMenu{config: lmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lunchmenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchmenu.FieldID,
			},
		}
	)
	if value, ok := lmc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lunchmenu.FieldDate,
		})
		_node.Date = value
	}
	if nodes := lmc.mutation.ItemsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lmc.mutation.ResturantIDs(); len(nodes) > 0 {
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
		_node.resturant_menu = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LunchMenuCreateBulk is the builder for creating many LunchMenu entities in bulk.
type LunchMenuCreateBulk struct {
	config
	builders []*LunchMenuCreate
}

// Save creates the LunchMenu entities in the database.
func (lmcb *LunchMenuCreateBulk) Save(ctx context.Context) ([]*LunchMenu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lmcb.builders))
	nodes := make([]*LunchMenu, len(lmcb.builders))
	mutators := make([]Mutator, len(lmcb.builders))
	for i := range lmcb.builders {
		func(i int, root context.Context) {
			builder := lmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LunchMenuMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lmcb *LunchMenuCreateBulk) SaveX(ctx context.Context) []*LunchMenu {
	v, err := lmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lmcb *LunchMenuCreateBulk) Exec(ctx context.Context) error {
	_, err := lmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmcb *LunchMenuCreateBulk) ExecX(ctx context.Context) {
	if err := lmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
