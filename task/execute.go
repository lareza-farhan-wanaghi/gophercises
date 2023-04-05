package task

import (
	"time"

	"github.com/boltdb/bolt"
)

var taskDb *taskDatabase

// Execute establishes a connection to the DB and processes the command specified
func Execute() error {
	dbOps := &bolt.Options{Timeout: 1 * time.Second}
	db, err := bolt.Open("my.db", 0600, dbOps)
	if err != nil {
		return err
	}
	defer db.Close()

	taskDb, err = newTaskDatabase(db)
	if err != nil {
		return err
	}

	err = rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
