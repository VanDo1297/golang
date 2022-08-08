package database

import (
	"demo/internal/logging"

	"github.com/go-errors/errors"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDatabaseConnector interface {
	ConnectToDB() error
	GetConnection() *gorm.DB
}

type ServerDBConnector struct {
	connection *gorm.DB
	Logger     logging.Logger
}

func (connector ServerDBConnector) isDBConnected() bool {
	if connector.connection == nil {
		return false
	}
	innerDB, connErr := connector.connection.DB()
	if connErr != nil {
		return false
	}
	connErr = innerDB.Ping()
	return connErr == nil
}

func (connector *ServerDBConnector) ConnectToDB() error {
	maxIdleConnections := 1
	maxOpenConnections := 10
	dvr := "mysql"
	source := "doift:infinitech@tcp(demodb.cvz2bbev6g4f.us-west-1.rds.amazonaws.com:3306)/demo?parseTime=True"

	connector.Logger.Info("Connecting to Database....", "driver", dvr, "source", source)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: source,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	if err == nil {
		innerDB, err := db.DB()
		if err != nil {
			return errors.Wrap(err, 0)
		}
		err = innerDB.Ping()
		if err != nil {
			return errors.Wrap(err, 0)
		}
		if err != nil {
			return errors.Wrap(err, 0)
		}

		innerDB.SetMaxIdleConns(maxIdleConnections)
		innerDB.SetMaxOpenConns(maxOpenConnections)
		connector.connection = db
		connector.Logger.Info("Connected to database!")
	} else {
		return errors.Wrap(err, 0)
	}
	return nil
}

func (connector *ServerDBConnector) GetConnection() *gorm.DB {
	if !connector.isDBConnected() {
		err := connector.ConnectToDB()
		if err != nil {
			connector.Logger.Error("could not get connection to db", "error", err)
			return nil
		}
	}

	return connector.connection
}
