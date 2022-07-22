// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LunchMenusColumns holds the columns for the "lunch_menus" table.
	LunchMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "resturant_menu", Type: field.TypeInt},
	}
	// LunchMenusTable holds the schema information for the "lunch_menus" table.
	LunchMenusTable = &schema.Table{
		Name:       "lunch_menus",
		Columns:    LunchMenusColumns,
		PrimaryKey: []*schema.Column{LunchMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lunch_menus_resturants_menu",
				Columns:    []*schema.Column{LunchMenusColumns[2]},
				RefColumns: []*schema.Column{ResturantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LunchMenuItemsColumns holds the columns for the "lunch_menu_items" table.
	LunchMenuItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "body", Type: field.TypeString},
		{Name: "language", Type: field.TypeEnum, Nullable: true, Enums: []string{"se", "en"}},
		{Name: "preformatted", Type: field.TypeBool, Default: false},
		{Name: "allergens", Type: field.TypeJSON, Nullable: true},
		{Name: "emission", Type: field.TypeFloat64, Nullable: true},
		{Name: "price", Type: field.TypeString, Nullable: true},
		{Name: "lunch_menu_items", Type: field.TypeInt},
	}
	// LunchMenuItemsTable holds the schema information for the "lunch_menu_items" table.
	LunchMenuItemsTable = &schema.Table{
		Name:       "lunch_menu_items",
		Columns:    LunchMenuItemsColumns,
		PrimaryKey: []*schema.Column{LunchMenuItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lunch_menu_items_lunch_menus_items",
				Columns:    []*schema.Column{LunchMenuItemsColumns[8]},
				RefColumns: []*schema.Column{LunchMenusColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ResturantsColumns holds the columns for the "resturants" table.
	ResturantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "campus", Type: field.TypeEnum, Enums: []string{"lindholmen", "johanneberg"}},
	}
	// ResturantsTable holds the schema information for the "resturants" table.
	ResturantsTable = &schema.Table{
		Name:       "resturants",
		Columns:    ResturantsColumns,
		PrimaryKey: []*schema.Column{ResturantsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LunchMenusTable,
		LunchMenuItemsTable,
		ResturantsTable,
	}
)

func init() {
	LunchMenusTable.ForeignKeys[0].RefTable = ResturantsTable
	LunchMenuItemsTable.ForeignKeys[0].RefTable = LunchMenusTable
}
