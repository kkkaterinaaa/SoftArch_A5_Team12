package initializers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"twitter/message-service/models"
)

func SyncDatabase() {
	migrateModel(models.Message{})
}

func migrateModel(model interface{}) {
	if !DB.Migrator().HasTable(model) {
		if err := DB.AutoMigrate(model); err != nil {
			fmt.Printf("Database migration for %T failed: %v", model, err)
		}
	} else {
		logrus.Infof("Skipping migration for %T - table already exists", model)
	}
}
