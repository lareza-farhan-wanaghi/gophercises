package task

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/boltdb/bolt"
)

type taskModel struct {
	Title      string    `json:"title"`
	FinishedAt time.Time `json:"finished_at"`
}

type taskDatabase struct {
	db    *bolt.DB
	bName []byte
}

// initDb handles the basic setup needed to interact with the database
func (t *taskDatabase) initDb() error {
	tx, err := t.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists(t.bName)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// flushDb removes data the main bucket used in this database
func (t *taskDatabase) flushDb() error {
	tx, err := t.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.DeleteBucket(t.bName)
	if err != nil {
		return err
	}

	_, err = tx.CreateBucket(t.bName)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// addTask adds the task to the list of active tasks
func (t *taskDatabase) addTask(taskTitle string) error {
	return t.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(t.bName)

		tm := taskModel{
			Title:      taskTitle,
			FinishedAt: time.Time{},
		}
		data, err := json.Marshal(tm)
		if err != nil {
			return err
		}

		err = b.Put([]byte(taskTitle), data)
		return err
	})
}

// listTask lists all active tasks
func (t *taskDatabase) listTask() ([]taskModel, error) {
	result := []taskModel{}

	err := t.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(t.bName)
		b.ForEach(func(k, v []byte) error {
			var tm taskModel
			err := json.Unmarshal(v, &tm)
			if err != nil {
				return err
			}

			if tm.FinishedAt.IsZero() {
				result = append(result, tm)
			}
			return nil
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// doTask updates a task with the same id to be marked as completed
func (t *taskDatabase) doTask(id int) (*taskModel, error) {
	var tm *taskModel

	err := t.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(t.bName)

		taskList, err := t.listTask()
		if err != nil {
			return err
		}

		if id < 0 || id >= len(taskList) {
			return nil
		}

		tm = &taskList[id]
		tm.FinishedAt = time.Now()
		data, err := json.Marshal(*tm)
		if err != nil {
			return err
		}

		return b.Put([]byte(tm.Title), data)
	})

	if err != nil {
		return nil, err
	}
	return tm, nil
}

// deleteTask deletes an active task that has the same id as specified
func (t *taskDatabase) deleteTask(id int) (*taskModel, error) {
	var tm *taskModel
	err := t.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(t.bName)

		taskList, err := t.listTask()
		if err != nil {
			return err
		}

		if id < 0 || id >= len(taskList) {
			return nil
		}

		tm = &taskList[id]
		return b.Delete([]byte(tm.Title))
	})

	if err != nil {
		return nil, err
	}
	return tm, nil
}

// completedTask retrieves all completed tasks that were done less than a day in a time-sorted manner
func (t *taskDatabase) completedTask() ([]taskModel, error) {
	result := []taskModel{}

	err := t.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(t.bName)
		b.ForEach(func(k, v []byte) error {
			var tm taskModel
			err := json.Unmarshal(v, &tm)
			if err != nil {
				return err
			}
			current := time.Now()
			yesterday := current.AddDate(0, 0, -1)

			if tm.FinishedAt.After(yesterday) && tm.FinishedAt.Before(current) {
				result = append(result, tm)
			}
			return nil
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].FinishedAt.Before(result[j].FinishedAt)
	})

	return result, nil
}

// newTaskDatabase creates a new taskDatabase with its default configurations
func newTaskDatabase(db *bolt.DB) (*taskDatabase, error) {
	result := &taskDatabase{
		db:    db,
		bName: []byte("task"),
	}

	err := result.initDb()
	if err != nil {
		return nil, err
	}

	return result, nil
}
