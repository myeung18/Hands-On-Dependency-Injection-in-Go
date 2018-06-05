package data

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/common/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/config"
	// import the MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	// default person id (returned on error)
	defaultPersonID = 0
)

var (
	dbInit sync.Once
	dbPool *sql.DB

	// ErrNotFound is returned when the no records where matched by the query
	ErrNotFound = errors.New("not found")
)

func getDB() (*sql.DB, error) {
	if dbPool == nil {
		if config.App == nil {
			return nil, errors.New("config is not initialized")
		}

		dbInit.Do(func() {
			var err error
			dbPool, err = sql.Open("mysql", config.App.DSN)
			if err != nil {
				// if the DB cannot be accessed we are dead
				panic(err.Error())
			}
		})
	}

	return dbPool, nil
}

// Person is the data transfer object (DTO) for this package
// This is an intentional duplication of the Person definition in order to reduce inter-dependence between packages or
// the creation of a "common" package.  Shared packages create pressure on the definition such that the external API
// format resembles the storage format or vice versa.  This pressure makes it harder for these formats to be evolved
// and maintained separately.
type Person struct {
	// ID is the unique ID for this person
	ID int
	// FullName is the name of this person
	FullName string
	// Phone is the phone for this person
	Phone string
	// Currency is the currency this person has paid in
	Currency string
	// Price is the amount (in the above currency) paid by this person
	Price float64
}

// Save will save the supplied person and return the ID of the newly created person or an error.
// Errors returned are caused by the underlying database or our connection to it.
func Save(in *Person) (int, error) {
	db, err := getDB()
	if err != nil {
		logging.Error("failed to get DB connection. err: %s", err)
		return defaultPersonID, err
	}

	// perform DB insert
	query := "INSERT INTO person (fullname, phone, currency, price) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, in.FullName, in.Phone, in.Currency, in.Price)
	if err != nil {
		logging.Error("failed to save person into DB. err: %s", err)
		return defaultPersonID, err
	}

	// retrieve and return the ID of the person created
	id, err := result.LastInsertId()
	if err != nil {
		logging.Error("failed to retrieve id of last saved person. err: %s", err)
		return defaultPersonID, err
	}
	return int(id), nil
}

// LoadAll will attempt to load all people in the database
// It will return ErrNotFound when there are not people in the database
// Any other errors returned are caused by the underlying database or our connection to it.
func LoadAll() ([]*Person, error) {
	db, err := getDB()
	if err != nil {
		logging.Error("failed to get DB connection. err: %s", err)
		return nil, err
	}

	// perform DB select
	query := "SELECT id, fullname, phone, currency, price FROM person"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	out := []*Person{}

	for rows.Next() {
		// retrieve columns and populate the person object
		record, err := populatePerson(rows.Scan)
		if err != nil {
			logging.Error("failed to convert query result. err: %s", err)
			return nil, err
		}

		out = append(out, record)
	}

	if len(out) == 0 {
		logging.Warn("no people found in the database.")
		return nil, ErrNotFound
	}

	return out, nil
}

// Load will attempt to load and return a person.
// It will return ErrNotFound when the requested person does not exist.
// Any other errors returned are caused by the underlying database or our connection to it.
func Load(in int) (*Person, error) {
	db, err := getDB()
	if err != nil {
		logging.Error("failed to get DB connection. err: %s", err)
		return nil, err
	}

	// perform DB select
	query := "SELECT id, fullname, phone, currency, price FROM person WHERE id = ? LIMIT 1"
	row := db.QueryRow(query, in)

	// retrieve columns and populate the person object
	out, err := populatePerson(row.Scan)
	if err != nil {
		if err == sql.ErrNoRows {
			logging.Warn("failed to load requested person '%d'. err: %s", in, err)
			return nil, ErrNotFound
		}

		logging.Error("failed to convert query result. err: %s", err)
		return nil, err
	}
	return out, nil
}

// custom type so we can convert sql results to easily
type scanner func(dest ...interface{}) error

// reduce the duplication (and maintenance) between sql.Row and sql.Rows usage
func populatePerson(scanner scanner) (*Person, error) {
	out := &Person{}
	err := scanner(&out.ID, &out.FullName, &out.Phone, &out.Currency, &out.Price)
	return out, err
}
