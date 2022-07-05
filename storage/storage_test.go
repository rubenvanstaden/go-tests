package storage_test

import (
	"os"
	"testing"

	. "go-tests/storage"
	. "go-tests/utils"
)

func TestStorage_Data(t *testing.T) {

    path, err := os.MkdirTemp("", "dir")
    defer os.RemoveAll(path)

	s, err := Open(path)
	if err != nil {
        panic(err)
	}
    defer s.Close()

    err = s.WriteData("hello.txt", "hello, world")
    Ok(t, err)
}
