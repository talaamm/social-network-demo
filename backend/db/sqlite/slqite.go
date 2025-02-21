package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func ApplyMigrations(db *sql.DB, migrationsPath string) error {
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("Error in reading migration directory:\n%v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			filePath := filepath.Join(migrationsPath, file.Name())
			err := executeSQLFile(db, filePath)
			if err != nil {
				return fmt.Errorf("Error in reading & executing migration file:\n%s:\n%v", file.Name(), err)
			}
			fmt.Printf("✅Successfully executed: %s\n", file.Name())
		}
	}
	return nil
}

// executeSQLFile reads a file path (a migration file) & executes it on db
func executeSQLFile(db *sql.DB, filePath string) error {
	query, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(query))
	return err
}
