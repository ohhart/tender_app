package migrations

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB, migrationsDir string) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			path := filepath.Join(migrationsDir, file.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			log.Printf("Applying migration: %s", file.Name())
			_, err = db.Exec(string(content))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
