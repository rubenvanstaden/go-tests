package main

import (
	// "os"
	// "context"
	// "fmt"
	// "path/filepath"

	"go-tests/app"
	"go-tests/jobs"
	// "go-tests/storage"
)

func main() {

    client := jobs.NewClient()

    a := app.App{
        JobClient: client,
    }

    a.Submit("new job submitted")

	//     path, err := os.MkdirTemp("", "dir")
	// s, err := storage.Open(path)
	// if err != nil {
	//         panic(err)
	// }
	//     defer s.Close()
	//
	//     err = s.WriteData("hello.txt", "hello, world")

    // _ = storage.NewTestStorage()
    // defer s.Close()

    // // fname := filepath.Join("", "file1")
    // data := []byte("hello, world")
    // // _ = os.WriteFile(fname, data, 0666)
    //
    // err := s.WriteAll(context.Background(), "awe.txt", data, nil)
    // if err != nil {
    //     fmt.Println(err)
    // }
}
