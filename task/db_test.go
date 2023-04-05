package task

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// TestAddTask tests the addTask function of the taskDatabase struct
func TestAddTask(t *testing.T) {
	for k, v := range testTable.addTask {
		for _, taskTitle := range strings.Split(k, ",") {
			err := taskDb.addTask(taskTitle)
			if err != nil {
				t.Fatalf("%v. k: %s, v:%s", err, k, v)
			}
		}

		data, err := taskDb.listTask()
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		expectedData := strings.Split(v, ",")
		for i, d := range data {
			if d.Title != expectedData[i] {
				t.Fatalf("expected %s but got %s. in %s out: %s", expectedData[i], d.Title, k, v)
			}
		}

		err = taskDb.flushDb()
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}
	}
}

// TestDoTask tests the doTask function of the taskDatabase struct
func TestDoTask(t *testing.T) {
	for k, v := range testTable.doTask {
		splits := strings.Split(k, ",")
		nextIndex, err := batchAddTask(0, splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		j, err := strconv.Atoi(splits[nextIndex])
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		data, err := taskDb.doTask(j)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		if data == nil {
			if v != "" {
				t.Fatalf("expected %s but got %s. in: %s, out: %s", v, data.Title, k, v)
			}
		} else if data.Title != v {
			t.Fatalf("expected %s but got %s. in: %s, out: %s", v, data.Title, k, v)
		}

		err = taskDb.flushDb()
		if err != nil {
			t.Fatal(err)
		}
	}
}

// TestDeleteTask tests the deleteTask function of the taskDatabase struct
func TestDeleteTask(t *testing.T) {
	for k, v := range testTable.deleteTask {
		splits := strings.Split(k, ",")
		nextIndex, err := batchAddTask(0, splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		j, err := strconv.Atoi(splits[nextIndex])
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		data, err := taskDb.deleteTask(j)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		if data == nil {
			if v != "" {
				t.Fatalf("expected %s but got %s. in: %s, out: %s", v, data.Title, k, v)
			}
		} else if data.Title != v {
			t.Fatalf("expected %s but got %s. in: %s, out: %s", v, data.Title, k, v)
		}

		err = taskDb.flushDb()
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}
	}
}

// TestComplexTask tests the all db interaction functions of the taskDatabase struct
func TestComplexTask(t *testing.T) {
	for k, v := range testTable.complexTask {
		splits := strings.Split(k, ",")
		nextIndex, err := batchAddTask(0, splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}
		nextIndex, err = batchDoTask(nextIndex, splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		_, err = batchDeleteTask(nextIndex, splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		splits = strings.Split(v, ",")
		nextIndex, err = compareCompletedTasks(0, splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}
		_, err = compareActiveTasks(nextIndex, splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		err = taskDb.flushDb()
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}
	}
}

// batchAddTask adds several tasks based on the splits. Reusable helper function to setup a test
func batchAddTask(indexOfN int, splits []string) (int, error) {
	m, err := strconv.Atoi(splits[indexOfN])
	if err != nil {
		return -1, err
	}

	for i := 0; i < m; i++ {
		err = taskDb.addTask(splits[indexOfN+1+i])
		if err != nil {
			return -1, err
		}
	}
	return indexOfN + 1 + m, nil
}

// batchDoTask marks several tasks as completed based on the splits. Reusable helper function to setup a test
func batchDoTask(indexOfN int, splits []string) (int, error) {
	m, err := strconv.Atoi(splits[indexOfN])
	if err != nil {
		return -1, err
	}

	for i := 0; i < m; i++ {
		id, err := strconv.Atoi(splits[indexOfN+1+i])
		if err != nil {
			return -1, err
		}
		_, err = taskDb.doTask(id)
		if err != nil {
			return -1, err
		}
	}
	return indexOfN + 1 + m, nil
}

// batchDeleteTask deletes several tasks based on the splits. Reusable helper function to setup the test
func batchDeleteTask(indexOfN int, splits []string) (int, error) {
	m, err := strconv.Atoi(splits[indexOfN])
	if err != nil {
		return -1, err
	}

	for i := 0; i < m; i++ {
		id, err := strconv.Atoi(splits[indexOfN+1+i])
		if err != nil {
			return -1, err
		}
		_, err = taskDb.deleteTask(id)
		if err != nil {
			return -1, err
		}
	}
	return indexOfN + 1 + m, nil
}

// compareCompletedTasks compares stored completed tasks and the data from the splits. Reusable helper function to setup the test
func compareCompletedTasks(indexOfN int, splits []string) (int, error) {
	m, err := strconv.Atoi(splits[indexOfN])
	if err != nil {
		return -1, err
	}

	data, err := taskDb.completedTask()
	if err != nil {
		return -1, err
	}

	if len(data) != m {
		return -1, fmt.Errorf("expected %d datas but only got %d datas", m, len(data))
	}

	for i := 0; i < m; i++ {
		if data[i].Title != splits[indexOfN+1+i] {
			return -1, fmt.Errorf("expected %s but got %s.", splits[indexOfN+1+i], data[i].Title)
		}
	}
	return indexOfN + 1 + m, nil
}

// compareActiveTasks compares stored active tasks and the data from the splits. Reusable helper function to setup the test
func compareActiveTasks(indexOfN int, splits []string) (int, error) {
	m, err := strconv.Atoi(splits[indexOfN])
	if err != nil {
		return -1, err
	}

	data, err := taskDb.listTask()
	if err != nil {
		return -1, err
	}

	if len(data) != m {
		return -1, fmt.Errorf("expected %d datas but got %d datas", m, len(data))
	}

	for i := 0; i < m; i++ {
		if data[i].Title != splits[indexOfN+1+i] {
			return -1, fmt.Errorf("expected %s but got %s.", splits[indexOfN+1+i], data[i].Title)
		}
	}
	return indexOfN + 1 + m, nil
}
