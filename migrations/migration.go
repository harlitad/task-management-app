package migrations

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dsn string) error {
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalf("failed to get current working directory: %v", err)
	// }
	// migrationsPath := filepath.Join(cwd, "migrations")
	// fmt.Println("dsn: ", dsn)
	// // Split DSN to create a new connection string for golang-migrate
	// m, err := migrate.New("file://"+migrationsPath, dsn)
	// if err != nil {
	// 	return fmt.Errorf("failed to create migrate instance: %w", err)
	// }

	// // if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// // 	return fmt.Errorf("failed to apply migrations: %w", err)
	// // }

	// fmt.Println(">> SUCCESS MIGRATE")

	return nil
}
