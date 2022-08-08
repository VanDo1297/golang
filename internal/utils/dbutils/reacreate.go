package dbutils

import (
	"demo/internal/database"
	"demo/internal/models/dbmodels"
	"strings"
)

type tableScan struct {
	TableName string `gorm:"column:table_name"`
}

func executeStatements(db database.IDatabaseConnector, tableNames []*tableScan) {
	for _, tableName := range tableNames {
		if len(strings.TrimSpace(tableName.TableName)) > 0 {
			err := db.GetConnection().Migrator().DropTable(tableName.TableName)
			if err != nil {
				panic("Could not drop: " + tableName.TableName)
			}
		}
	}
}

func RecreateDatabase(db database.IDatabaseConnector) {
	dropAndCreate(db, dbmodels.Demo{})
}

func UpgradeDB(db database.IDatabaseConnector) {

}

type IndexData struct {
	Name    string
	Columns []string
}

func dropAndCreate(db database.IDatabaseConnector, table interface{}) {
	err := db.GetConnection().Migrator().DropTable(table)
	if err != nil {
		panic(err)
	}
	err = db.GetConnection().Migrator().CreateTable(table)
	if err != nil {
		panic(err)
	}
}

func IsDuplicateError(err error, key string) bool {
	if err != nil {
		errMessage := strings.ToLower(err.Error())
		keyVal := strings.ToLower(key)
		if strings.Contains(errMessage, "duplicate") && strings.Contains(errMessage, keyVal) {
			return true
		}
	}
	return false
}
