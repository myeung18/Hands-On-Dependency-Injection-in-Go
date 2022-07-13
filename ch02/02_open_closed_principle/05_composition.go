package ocp

import (
	"database/sql"
)

type rowConverter struct {
}

// populate the supplied Person from *sql.Row or *sql.Rows object
func (d *rowConverter) populate(in *Person, scan func(dest ...interface{}) error) error {
	return scan(in.Name, in.Email)
}

func (d *rowConverter) testtest(in *Person, doAction func() int) error  {

	doAction()

	return nil

}

type LoadPerson struct {
	// compose the row converter into this loader
	rowConverter
}

func (loader *LoadPerson) ByID(id int) (Person, error) {
	row := loader.loadFromDB(id)

	person := Person{}
	// call the composed "abstract class"
	err := loader.populate(&person, row.Scan)

	return person, err
}

func (loader *LoadPerson) loadFromDB(id int) *sql.Row {
	// TODO: implement
	return nil
}

//overriding
func (loader *LoadPerson) populate(in *Person, scan func(dest ...interface{}) error) error {
	return scan(in.Name, "")
}

type LoadAll struct {
	// compose the row converter into this loader
	rowConverter
}

func (loader *LoadPerson) All() ([]Person, error) {
	rows := loader.loadAllFromDB()
	defer rows.Close()

	var output []Person
	for rows.Next() {
		person := Person{}

		// call the composed "abstract class"
		err := loader.populate(&person, rows.Scan)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

func (loader *LoadPerson) loadAllFromDB() *sql.Rows {
	// TODO: implement
	return nil
}
