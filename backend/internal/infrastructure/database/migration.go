package database

import (
	"log"

	authEntity "celestara.com/hris-api/internal/module/auth/entity"
	employeeEntity "celestara.com/hris-api/internal/module/employee/entity"
	tenantEntity "celestara.com/hris-api/internal/module/tenant/entity"
	"gorm.io/gorm"
)

// RunDevMigration
func RunDevMigration(db *gorm.DB, env string) {
	if env != "development" {
		return
	}

	err := db.AutoMigrate(
		&tenantEntity.Company{},
		&authEntity.User{},
		&employeeEntity.Employee{},
	)

	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	log.Println("Database schema migrated sucessfully (DEV MODE).")
}