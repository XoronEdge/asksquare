package initial

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

//CleanUpHook cleans the database
func CleanUpHook(db *gorm.DB) func() {
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "postgres_cleanup_hook"
	// Setup the onCreate Hook
	db.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
		entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
	})
	return func() {
		// Remove the hook once we're done
		defer db.Callback().Create().Remove(hookName)
		// Find out if the current db object is already a transaction
		_, inTransaction := db.CommonDB().(*sql.Tx)
		tx := db
		if !inTransaction {
			tx = db.Begin()
		}
		// Loop from the end. It is important that we delete the entries in the
		// reverse order of their insertion
		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]
			tx.Table(entry.table).Where(entry.keyname+" = ?", entry.key).Delete("")
		}
		if !inTransaction {
			tx.Commit()
		}
	}
}
