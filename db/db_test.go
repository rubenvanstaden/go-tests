package db_test

import (
	"testing"

	. "go-tests/utils"
	. "go-tests/db"
)

func TestDB_Insert(t *testing.T) {

	db, err := Open("mongodb://localhost:27017/")
	if err != nil {
        panic(err)
	}

    inputJob := &Job{
        Data: "some data",
    }

    _, err = db.Insert(inputJob)
    Ok(t, err)
}
