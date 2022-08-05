package database

import (
	"database/sql"
	"errors"
	"log"
	"time"

	c "github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"
)

func (dba *Adapter) CreateCat(cat *c.Cat) error {
	statement, err := dba.conn.Prepare("INSERT INTO Cats (name, breed, dob, humanYears, catFact, CreatedAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		cat.Name,
		cat.Breed,
		cat.DOB,
		cat.HumanYears,
		cat.CatFact,
		time.Now().UTC())
	if err != nil {
		return err
	}

	return nil
}

func (dba *Adapter) RetrieveCat(catName string) (*c.Cat, error) {
	query := `SELECT name, breed, dob, humanYears, catFact FROM Cats WHERE name = ? ORDER BY CreatedAt DESC LIMIT 1;`
	row := dba.conn.QueryRow(query, catName)

	var name string
	var breed string
	var dob string
	var humanYears int64
	var catFact string

	switch err := row.Scan(
		&name,
		&breed,
		&dob,
		&humanYears,
		&catFact); err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		return nil, errors.New("empty result")
	case nil:
		log.Println("Cat found!")
	default:
		return nil, err
	}

	cat := c.NewCat()
	cat.Name = name
	cat.Breed = breed
	cat.DOB = dob
	cat.HumanYears = humanYears
	cat.CatFact = catFact

	return cat, nil
}

func (dba *Adapter) UpdateCat(cat *c.Cat) error {
	statement, err := dba.conn.Prepare("update Cats set breed=?, dob=?, humanYears=?, catFact=?, UpdatedAt=? where name=?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		cat.Breed,
		cat.DOB,
		cat.HumanYears,
		cat.CatFact,
		time.Now().UTC(),
		cat.Name)
	if err != nil {
		return err
	}

	return nil
}

func (dba *Adapter) DeleteCat(catName string) error {
	statement, err := dba.conn.Prepare("delete from Cats where name=?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(catName)
	if err != nil {
		return err
	}
	return nil
}
