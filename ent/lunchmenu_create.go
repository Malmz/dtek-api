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
	"github.com/dtekcth/dtek-api/ent/schema"
)

// LunchMenuCreate is the builder for creating a LunchMenu entity.
type LunchMenuCreate struct {
	config
	mutation *LunchMenuMutation
	hooks    []Hook
}

// SetUpdateTime sets the "update_time" field.
func (lmc *LunchMenuCreate) SetUpdateTime(t time.Time) *LunchMenuCreate {
	lmc.mutation.SetUpdateTime(t)
	return lmc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (lmc *LunchMenuCreate) SetNillableUpdateTime(t *time.Time) *LunchMenuCreate {
	if t != nil {
		lmc.SetUpdateTime(*t)
	}
	return lmc
}

// SetResturant sets the "resturant" field.
func (lmc *LunchMenuCreate) SetResturant(s string) *LunchMenuCreate {
	lmc.mutation.SetResturant(s)
	return lmc
}

// SetDate sets the "date" field.
func (lmc *LunchMenuCreate) SetDate(t time.Time) *LunchMenuCreate {
	lmc.mutation.SetDate(t)
	return lmc
}

// SetLanguage sets the "language" field.
func (lmc *LunchMenuCreate) SetLanguage(l lunchmenu.Language) *LunchMenuCreate {
	lmc.mutation.SetLanguage(l)
	return lmc
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (lmc *LunchMenuCreate) SetNillableLanguage(l *lunchmenu.Language) *LunchMenuCreate {
	if l != nil {
		lmc.SetLanguage(*l)
	}
	return lmc
}

// SetName sets the "name" field.
func (lmc *LunchMenuCreate) SetName(s string) *LunchMenuCreate {
	lmc.mutation.SetName(s)
	return lmc
}

// SetMenu sets the "menu" field.
func (lmc *LunchMenuCreate) SetMenu(smi []schema.LunchMenuItem) *LunchMenuCreate {
	lmc.mutation.SetMenu(smi)
	return lmc
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
	lmc.defaults()
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

// defaults sets the default values of the builder before save.
func (lmc *LunchMenuCreate) defaults() {
	if _, ok := lmc.mutation.UpdateTime(); !ok {
		v := lunchmenu.DefaultUpdateTime()
		lmc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmc *LunchMenuCreate) check() error {
	if _, ok := lmc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "LunchMenu.update_time"`)}
	}
	if _, ok := lmc.mutation.Resturant(); !ok {
		return &ValidationError{Name: "resturant", err: errors.New(`ent: missing required field "LunchMenu.resturant"`)}
	}
	if _, ok := lmc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "LunchMenu.date"`)}
	}
	if v, ok := lmc.mutation.Language(); ok {
		if err := lunchmenu.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "LunchMenu.language": %w`, err)}
		}
	}
	if _, ok := lmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "LunchMenu.name"`)}
	}
	if _, ok := lmc.mutation.Menu(); !ok {
		return &ValidationError{Name: "menu", err: errors.New(`ent: missing required field "LunchMenu.menu"`)}
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
	if value, ok := lmc.mutation.UpdateTime(); ok {
		_spec.SetField(lunchmenu.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := lmc.mutation.Resturant(); ok {
		_spec.SetField(lunchmenu.FieldResturant, field.TypeString, value)
		_node.Resturant = value
	}
	if value, ok := lmc.mutation.Date(); ok {
		_spec.SetField(lunchmenu.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := lmc.mutation.Language(); ok {
		_spec.SetField(lunchmenu.FieldLanguage, field.TypeEnum, value)
		_node.Language = value
	}
	if value, ok := lmc.mutation.Name(); ok {
		_spec.SetField(lunchmenu.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := lmc.mutation.Menu(); ok {
		_spec.SetField(lunchmenu.FieldMenu, field.TypeJSON, value)
		_node.Menu = value
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
			builder.defaults()
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
