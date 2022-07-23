// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPermission holds the string denoting the permission field in the database.
	FieldPermission = "permission"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the account in the database.
	Table = "account"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldAccount,
	FieldPassword,
	FieldPermission,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
