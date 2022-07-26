// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/predicate"
	"github.com/dtekcth/dtek-api/model"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeLunchMenu = "LunchMenu"
)

// LunchMenuMutation represents an operation that mutates the LunchMenu nodes in the graph.
type LunchMenuMutation struct {
	config
	op            Op
	typ           string
	id            *int
	update_time   *time.Time
	resturant     *string
	date          *time.Time
	language      *lunchmenu.Language
	menu          *[]model.LunchMenuItem
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*LunchMenu, error)
	predicates    []predicate.LunchMenu
}

var _ ent.Mutation = (*LunchMenuMutation)(nil)

// lunchmenuOption allows management of the mutation configuration using functional options.
type lunchmenuOption func(*LunchMenuMutation)

// newLunchMenuMutation creates new mutation for the LunchMenu entity.
func newLunchMenuMutation(c config, op Op, opts ...lunchmenuOption) *LunchMenuMutation {
	m := &LunchMenuMutation{
		config:        c,
		op:            op,
		typ:           TypeLunchMenu,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLunchMenuID sets the ID field of the mutation.
func withLunchMenuID(id int) lunchmenuOption {
	return func(m *LunchMenuMutation) {
		var (
			err   error
			once  sync.Once
			value *LunchMenu
		)
		m.oldValue = func(ctx context.Context) (*LunchMenu, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().LunchMenu.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLunchMenu sets the old LunchMenu of the mutation.
func withLunchMenu(node *LunchMenu) lunchmenuOption {
	return func(m *LunchMenuMutation) {
		m.oldValue = func(context.Context) (*LunchMenu, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LunchMenuMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LunchMenuMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LunchMenuMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LunchMenuMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().LunchMenu.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUpdateTime sets the "update_time" field.
func (m *LunchMenuMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *LunchMenuMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the LunchMenu entity.
// If the LunchMenu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LunchMenuMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *LunchMenuMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetResturant sets the "resturant" field.
func (m *LunchMenuMutation) SetResturant(s string) {
	m.resturant = &s
}

// Resturant returns the value of the "resturant" field in the mutation.
func (m *LunchMenuMutation) Resturant() (r string, exists bool) {
	v := m.resturant
	if v == nil {
		return
	}
	return *v, true
}

// OldResturant returns the old "resturant" field's value of the LunchMenu entity.
// If the LunchMenu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LunchMenuMutation) OldResturant(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldResturant is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldResturant requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldResturant: %w", err)
	}
	return oldValue.Resturant, nil
}

// ResetResturant resets all changes to the "resturant" field.
func (m *LunchMenuMutation) ResetResturant() {
	m.resturant = nil
}

// SetDate sets the "date" field.
func (m *LunchMenuMutation) SetDate(t time.Time) {
	m.date = &t
}

// Date returns the value of the "date" field in the mutation.
func (m *LunchMenuMutation) Date() (r time.Time, exists bool) {
	v := m.date
	if v == nil {
		return
	}
	return *v, true
}

// OldDate returns the old "date" field's value of the LunchMenu entity.
// If the LunchMenu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LunchMenuMutation) OldDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDate: %w", err)
	}
	return oldValue.Date, nil
}

// ResetDate resets all changes to the "date" field.
func (m *LunchMenuMutation) ResetDate() {
	m.date = nil
}

// SetLanguage sets the "language" field.
func (m *LunchMenuMutation) SetLanguage(l lunchmenu.Language) {
	m.language = &l
}

// Language returns the value of the "language" field in the mutation.
func (m *LunchMenuMutation) Language() (r lunchmenu.Language, exists bool) {
	v := m.language
	if v == nil {
		return
	}
	return *v, true
}

// OldLanguage returns the old "language" field's value of the LunchMenu entity.
// If the LunchMenu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LunchMenuMutation) OldLanguage(ctx context.Context) (v lunchmenu.Language, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLanguage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLanguage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLanguage: %w", err)
	}
	return oldValue.Language, nil
}

// ClearLanguage clears the value of the "language" field.
func (m *LunchMenuMutation) ClearLanguage() {
	m.language = nil
	m.clearedFields[lunchmenu.FieldLanguage] = struct{}{}
}

// LanguageCleared returns if the "language" field was cleared in this mutation.
func (m *LunchMenuMutation) LanguageCleared() bool {
	_, ok := m.clearedFields[lunchmenu.FieldLanguage]
	return ok
}

// ResetLanguage resets all changes to the "language" field.
func (m *LunchMenuMutation) ResetLanguage() {
	m.language = nil
	delete(m.clearedFields, lunchmenu.FieldLanguage)
}

// SetMenu sets the "menu" field.
func (m *LunchMenuMutation) SetMenu(mmi []model.LunchMenuItem) {
	m.menu = &mmi
}

// Menu returns the value of the "menu" field in the mutation.
func (m *LunchMenuMutation) Menu() (r []model.LunchMenuItem, exists bool) {
	v := m.menu
	if v == nil {
		return
	}
	return *v, true
}

// OldMenu returns the old "menu" field's value of the LunchMenu entity.
// If the LunchMenu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LunchMenuMutation) OldMenu(ctx context.Context) (v []model.LunchMenuItem, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMenu is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMenu requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMenu: %w", err)
	}
	return oldValue.Menu, nil
}

// ResetMenu resets all changes to the "menu" field.
func (m *LunchMenuMutation) ResetMenu() {
	m.menu = nil
}

// Where appends a list predicates to the LunchMenuMutation builder.
func (m *LunchMenuMutation) Where(ps ...predicate.LunchMenu) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *LunchMenuMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (LunchMenu).
func (m *LunchMenuMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LunchMenuMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.update_time != nil {
		fields = append(fields, lunchmenu.FieldUpdateTime)
	}
	if m.resturant != nil {
		fields = append(fields, lunchmenu.FieldResturant)
	}
	if m.date != nil {
		fields = append(fields, lunchmenu.FieldDate)
	}
	if m.language != nil {
		fields = append(fields, lunchmenu.FieldLanguage)
	}
	if m.menu != nil {
		fields = append(fields, lunchmenu.FieldMenu)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LunchMenuMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case lunchmenu.FieldUpdateTime:
		return m.UpdateTime()
	case lunchmenu.FieldResturant:
		return m.Resturant()
	case lunchmenu.FieldDate:
		return m.Date()
	case lunchmenu.FieldLanguage:
		return m.Language()
	case lunchmenu.FieldMenu:
		return m.Menu()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LunchMenuMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case lunchmenu.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case lunchmenu.FieldResturant:
		return m.OldResturant(ctx)
	case lunchmenu.FieldDate:
		return m.OldDate(ctx)
	case lunchmenu.FieldLanguage:
		return m.OldLanguage(ctx)
	case lunchmenu.FieldMenu:
		return m.OldMenu(ctx)
	}
	return nil, fmt.Errorf("unknown LunchMenu field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LunchMenuMutation) SetField(name string, value ent.Value) error {
	switch name {
	case lunchmenu.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case lunchmenu.FieldResturant:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetResturant(v)
		return nil
	case lunchmenu.FieldDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDate(v)
		return nil
	case lunchmenu.FieldLanguage:
		v, ok := value.(lunchmenu.Language)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLanguage(v)
		return nil
	case lunchmenu.FieldMenu:
		v, ok := value.([]model.LunchMenuItem)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMenu(v)
		return nil
	}
	return fmt.Errorf("unknown LunchMenu field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LunchMenuMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LunchMenuMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LunchMenuMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown LunchMenu numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LunchMenuMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(lunchmenu.FieldLanguage) {
		fields = append(fields, lunchmenu.FieldLanguage)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LunchMenuMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LunchMenuMutation) ClearField(name string) error {
	switch name {
	case lunchmenu.FieldLanguage:
		m.ClearLanguage()
		return nil
	}
	return fmt.Errorf("unknown LunchMenu nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LunchMenuMutation) ResetField(name string) error {
	switch name {
	case lunchmenu.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case lunchmenu.FieldResturant:
		m.ResetResturant()
		return nil
	case lunchmenu.FieldDate:
		m.ResetDate()
		return nil
	case lunchmenu.FieldLanguage:
		m.ResetLanguage()
		return nil
	case lunchmenu.FieldMenu:
		m.ResetMenu()
		return nil
	}
	return fmt.Errorf("unknown LunchMenu field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LunchMenuMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LunchMenuMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LunchMenuMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LunchMenuMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LunchMenuMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LunchMenuMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LunchMenuMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown LunchMenu unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LunchMenuMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown LunchMenu edge %s", name)
}
