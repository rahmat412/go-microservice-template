package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

// Users represents the users table
var Users = &struct {
	postgres.Table
	ID        postgres.ColumnInteger
	Username  postgres.ColumnString
	FirstName postgres.ColumnString
	LastName  postgres.ColumnString
	BirthDate postgres.ColumnDate
	Email     postgres.ColumnString
	Password  postgres.ColumnString
	IsActive  postgres.ColumnBool
	CreatedAt postgres.ColumnTimestamp
	// Add other columns as needed
}{
	Table:     postgres.NewTable("public", "users", "users"),
	ID:        postgres.IntegerColumn("id"),
	Username:  postgres.StringColumn("username"),
	FirstName: postgres.StringColumn("first_name"),
	LastName:  postgres.StringColumn("last_name"),
	BirthDate: postgres.DateColumn("birth_date"),
	Email:     postgres.StringColumn("email"),
	Password:  postgres.StringColumn("password"),
	IsActive:  postgres.BoolColumn("is_active"),
	CreatedAt: postgres.TimestampColumn("created_at"),
	// Add other column definitions as needed
}
