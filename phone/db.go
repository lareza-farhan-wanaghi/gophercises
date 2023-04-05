package phone

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

type PhoneModel struct {
	Id          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
}

// getJsonTag gets the json tag of the field
func (p *PhoneModel) getJsonTag(fieldIndex int) string {
	return reflect.ValueOf(*p).Type().Field(fieldIndex).Tag.Get("json")
}

type phoneDatabase struct {
	db        *sql.DB
	tableName string
}

// initDb handles the basic setup needed to interact with the database
func (p *phoneDatabase) initDb() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	samplePhone := PhoneModel{}

	syntax := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s serial4 NOT NULL,
			%s varchar NOT NULL,
			CONSTRAINT phone_numbers_pk PRIMARY KEY (id)
		);
	`, p.tableName, samplePhone.getJsonTag(0), samplePhone.getJsonTag(1))

	_, err := p.db.ExecContext(ctx, syntax)
	return err
}

// FlushDb removes data in the used table
func (p *phoneDatabase) FlushDb() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	syntax := fmt.Sprintf(`truncate %s RESTART IDENTITY;`, p.tableName)

	_, err := p.db.ExecContext(ctx, syntax)
	return err
}

// InsertPhone inserts the phone into the database
func (p *phoneDatabase) InsertPhone(phone *PhoneModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	syntax := fmt.Sprintf(`
		insert into %s (%s) values ('%s') returning id;
	`, p.tableName, phone.getJsonTag(1), phone.PhoneNumber)

	var lastInsertId int
	err := p.db.QueryRowContext(ctx, syntax).Scan(&lastInsertId)
	if err != nil {
		return -1, err
	}

	return lastInsertId, nil
}

// UpdatePhone updates the phone from the database
func (p *phoneDatabase) UpdatePhone(phone *PhoneModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	syntax := fmt.Sprintf(`
		update %s set %s='%s' where %s=%d;
	`, p.tableName, (*phone).getJsonTag(1),
		phone.PhoneNumber, (*phone).getJsonTag(0), phone.Id)

	_, err := p.db.ExecContext(ctx, syntax)
	if err != nil {
		return err
	}

	return nil
}

// DeletePhoneById deletes the phone from the database by its id
func (p *phoneDatabase) DeletePhoneById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	samplePhone := PhoneModel{}
	syntax := fmt.Sprintf(`
		delete from %s where %s=%d;
	`, p.tableName, samplePhone.getJsonTag(0), id)

	_, err := p.db.ExecContext(ctx, syntax)
	if err != nil {
		return err
	}

	return nil
}

// GetPhoneById returns a phone by id
func (p *phoneDatabase) GetPhoneById(id int) (*PhoneModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	phone := PhoneModel{}
	syntax := fmt.Sprintf(`
		select * from %s where %s=%d;
	`, p.tableName, phone.getJsonTag(0), id)

	err := p.db.QueryRowContext(ctx, syntax).Scan(&phone.Id, &phone.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &phone, nil
}

// GetAllPhones returns all phone records in the database
func (p *phoneDatabase) GetAllPhones() ([]*PhoneModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	syntax := fmt.Sprintf(`
		select * from %s order by id;
	`, p.tableName)

	rows, err := p.db.QueryContext(ctx, syntax)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	phones := []*PhoneModel{}
	for rows.Next() {
		var phone PhoneModel
		err = rows.Scan(&phone.Id, &phone.PhoneNumber)
		if err != nil {
			return nil, err
		}
		phones = append(phones, &phone)
	}

	return phones, nil
}

// NormalizePhoneData normalizes phone numbers in the database to contain numbers only and removes duplicate phone numbers
func (p *phoneDatabase) NormalizePhoneData() error {
	phones, err := p.GetAllPhones()
	if err != nil {
		return err
	}

	pnMap := map[string]struct{}{}
	re := regexp.MustCompile("[^0-9]+")
	for _, phone := range phones {
		normalizedPn := re.ReplaceAllString(phone.PhoneNumber, "")
		_, ok := pnMap[normalizedPn]
		if !ok {
			pnMap[normalizedPn] = struct{}{}
			if phone.PhoneNumber != normalizedPn {
				phone.PhoneNumber = normalizedPn
				err = p.UpdatePhone(phone)
				if err != nil {
					return err
				}
			}
		} else {
			err = p.DeletePhoneById(phone.Id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewPhoneDatabase creates a new phoneDatabase with its default configurations
func NewPhoneDatabase(db *sql.DB) (*phoneDatabase, error) {
	result := &phoneDatabase{
		db:        db,
		tableName: "phone_numbers",
	}

	err := result.initDb()
	if err != nil {
		return nil, err
	}

	return result, nil
}
