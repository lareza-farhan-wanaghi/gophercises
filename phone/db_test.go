package phone

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// TestNormalizePhoneData tests the normalizeData functions of the phoneDatabase struct
func TestNormalizePhoneData(t *testing.T) {
	for k, v := range testTable.normalizePhoneData {
		err := testPhoneDb.FlushDb()
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		splits := strings.Split(k, ",")
		err = parseTestKey(splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		err = testPhoneDb.NormalizePhoneData()
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		splits = strings.Split(v, ",")
		err = compareData(splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}
	}
}

// TestComplexPhoneTransac tests the all db interaction functions of the phoneDatabase struct
func TestComplexPhoneTransac(t *testing.T) {
	for k, v := range testTable.complexPhone {
		err := testPhoneDb.FlushDb()
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		splits := strings.Split(k, ",")
		err = parseTestKey(splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}

		splits = strings.Split(v, ",")
		err = compareData(splits)
		if err != nil {
			t.Fatalf("%v. k: %s, v:%s", err, k, v)
		}
	}
}

// parseTestKey reads key splits and do things based on the command specified
func parseTestKey(splits []string) error {
	var err error
	for i, m, c := 0, 0, ""; i < len(splits); {
		if c == "" {
			c = splits[i]
			i += 1
		} else {
			switch c {
			case "#i":
				m, err = strconv.Atoi(splits[i])
				if err != nil {
					return err
				}

				i, err = batchInsertPhone(m, i+1, splits)
				if err != nil {
					return err
				}
			case "#u":
				m, err = strconv.Atoi(splits[i])
				if err != nil {
					return err
				}

				m *= 2
				i, err = batchUpdatePhone(m, i+1, splits)
				if err != nil {
					return err
				}
			case "#d":
				m, err = strconv.Atoi(splits[i])
				if err != nil {
					return err
				}

				i, err = batchDeletePhone(m, i+1, splits)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("command expected: #i, #u, #d but got %v", c)
			}

			c = ""
		}
	}

	return nil
}

// batchInsertPhone inserts several phone based on the splits. Reusable helper function to setup the test
func batchInsertPhone(m, indexOfN int, splits []string) (int, error) {
	for i := 0; i < m; i++ {
		phone := &PhoneModel{
			PhoneNumber: splits[indexOfN+i],
		}
		_, err := testPhoneDb.InsertPhone(phone)
		if err != nil {
			return -1, err
		}
	}
	return indexOfN + m, nil
}

// batchUpdatePhone updates several phones based on the splits. Reusable helper function to setup the test
func batchUpdatePhone(m, indexOfN int, splits []string) (int, error) {
	for i := 0; i < m; i++ {
		id, err := strconv.Atoi(splits[indexOfN+i])
		if err != nil {
			return -1, err
		}

		phone, err := testPhoneDb.GetPhoneById(id)
		if err != nil {
			return -1, err
		}

		i += 1
		phone.PhoneNumber = splits[indexOfN+i]

		err = testPhoneDb.UpdatePhone(phone)
		if err != nil {
			return -1, err
		}
	}
	return indexOfN + m, nil
}

// batchDeletePhone deletes several phones based on the splits. Reusable helper function to setup the test
func batchDeletePhone(m, indexOfN int, splits []string) (int, error) {

	for i := 0; i < m; i++ {
		id, err := strconv.Atoi(splits[indexOfN+i])
		if err != nil {
			return -1, err
		}
		err = testPhoneDb.DeletePhoneById(id)
		if err != nil {
			return -1, err
		}
	}
	return indexOfN + m, nil
}

// compareData compares stored active phones and the data from the splits. Reusable helper function to setup a test
func compareData(splits []string) error {
	data, err := testPhoneDb.GetAllPhones()
	if err != nil {
		return err
	}

	if len(data) != len(splits) {
		return fmt.Errorf("expected %d datas but got %d datas", len(splits), len(data))
	}

	for i := 0; i < len(splits); i++ {
		if data[i].PhoneNumber != splits[i] {
			return fmt.Errorf("expected %s but got %s.", splits[i], data[i].PhoneNumber)
		}
	}

	return nil
}
