package models

func MigrationModels() error {
	DB.AutoMigrate(&Channel{})
	return nil
}
