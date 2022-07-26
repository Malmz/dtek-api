// Code generated by ent, DO NOT EDIT.

package lunchmenu

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the lunchmenu type in the database.
	Label = "lunch_menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldResturant holds the string denoting the resturant field in the database.
	FieldResturant = "resturant"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMenu holds the string denoting the menu field in the database.
	FieldMenu = "menu"
	// Table holds the table name of the lunchmenu in the database.
	Table = "lunch_menus"
)

// Columns holds all SQL columns for lunchmenu fields.
var Columns = []string{
	FieldID,
	FieldUpdateTime,
	FieldResturant,
	FieldDate,
	FieldLanguage,
	FieldName,
	FieldMenu,
}

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
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// Language defines the type for the "language" enum field.
type Language string

// Language values.
const (
	LanguageSe Language = "se"
	LanguageEn Language = "en"
)

func (l Language) String() string {
	return string(l)
}

// LanguageValidator is a validator for the "language" field enum values. It is called by the builders before save.
func LanguageValidator(l Language) error {
	switch l {
	case LanguageSe, LanguageEn:
		return nil
	default:
		return fmt.Errorf("lunchmenu: invalid enum value for language field: %q", l)
	}
}
