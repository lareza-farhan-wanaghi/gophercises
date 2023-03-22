package task

import (
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
				t.Fatal(err)
			}
		}

		data, err := taskDb.listTask()
		if err != nil {
			t.Fatal(err)
		}

		expectedData := strings.Split(v, ",")
		for i, d := range data {
			if d.Title != expectedData[i] {
				t.Fatalf("expected %s but got %s. in %s out: %s", expectedData[i], d.Title, k, v)
			}
		}

		err = taskDb.flushDb()
		if err != nil {
			t.Fatal(err)
		}
	}
}

// TestDoTask tests the doTask function of the taskDatabase struct
func TestDoTask(t *testing.T) {
	for k, v := range testTable.doTask {
		splits := strings.Split(k, ",")
		m, err := strconv.Atoi(splits[0])
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < m; i++ {
			err = taskDb.addTask(splits[i+1])
			if err != nil {
				t.Fatal(err)
			}
		}

		j, err := strconv.Atoi(splits[m+1])
		if err != nil {
			t.Fatal(err)
		}

		data, err := taskDb.doTask(j)
		if err != nil {
			t.Fatal(err)
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
		m, err := strconv.Atoi(splits[0])
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < m; i++ {
			err = taskDb.addTask(splits[i+1])
			if err != nil {
				t.Fatal(err)
			}
		}

		j, err := strconv.Atoi(splits[m+1])
		if err != nil {
			t.Fatal(err)
		}

		data, err := taskDb.deleteTask(j)
		if err != nil {
			t.Fatal(err)
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

// TestComplexTask tests the all db interaction functions of the taskDatabase struct
func TestComplexTask(t *testing.T) {
	for k, v := range testTable.complexTask {
		var err error
		c := 0
		m := 0

		for _, split := range strings.Split(k, ",") {
			if m == 0 {
				m, err = strconv.Atoi(split)
				if err != nil {
					t.Fatal(err)
				}
				c += 1
			} else {
				if c == 1 {
					err = taskDb.addTask(split)
					if err != nil {
						t.Fatal(err)
					}
				} else if c == 2 {
					j, err := strconv.Atoi(split)
					if err != nil {
						t.Fatal(err)
					}

					_, err = taskDb.doTask(j)
					if err != nil {
						t.Fatal(err)
					}
				} else {
					j, err := strconv.Atoi(split)
					if err != nil {
						t.Fatal(err)
					}

					_, err = taskDb.deleteTask(j)
					if err != nil {
						t.Fatal(err)
					}
				}
				m -= 1
			}
		}

		c = 0
		m = 0
		var data []taskModel

		for _, split := range strings.Split(v, ",") {
			if m == 0 {
				m, err = strconv.Atoi(split)
				if err != nil {
					t.Fatal(err)
				}
				c += 1
				if c == 1 {
					data, err = taskDb.completedTask()
					if err != nil {
						t.Fatal(err)
					}

				} else {
					data, err = taskDb.listTask()
					if err != nil {
						t.Fatal(err)
					}
				}
			} else {
				if data[len(data)-m].Title != split {
					t.Fatalf("expected %s but got %s. in: %s, out: %s", split, data[m-len(data)].Title, k, v)
				}
				m -= 1
			}
		}

		err = taskDb.flushDb()
		if err != nil {
			t.Fatal(err)
		}
	}
}
