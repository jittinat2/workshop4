package db

import (
	"database/sql"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init opens sqlite DB at path. If an existing `users` table uses an integer
// primary key (legacy schema), it renames it to `users_old` so AutoMigrate can
// create the new schema (with string LBK IDs).
func Init(path string) error {
	d, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return err
	}

	// inspect existing users table schema
	var rows *sql.Rows
	rows, err = d.Raw("PRAGMA table_info('users')").Rows()
	if err == nil {
		defer rows.Close()
		var hasID bool
		var idType string
		for rows.Next() {
			// cid, name, type, notnull, dflt_value, pk
			var cid int
			var name string
			var typ string
			var notnull int
			var dflt sql.NullString
			var pk int
			if err := rows.Scan(&cid, &name, &typ, &notnull, &dflt, &pk); err == nil {
				if name == "id" {
					hasID = true
					idType = typ
					break
				}
			}
		}
		if hasID {
			// SQLite reports types case-insensitively, check if id is integer -> legacy
			if idType != "text" && idType != "varchar" && idType != "varchar(16)" {
				log.Printf("detected legacy users.id type=%s; renaming table to users_old to recreate schema", idType)
				if err := d.Exec("ALTER TABLE users RENAME TO users_old").Error; err != nil {
					log.Printf("failed to rename legacy users table: %v", err)
				}
			}
		}
	}

	DB = d
	log.Printf("db initialized: %s", path)
	return nil
}
