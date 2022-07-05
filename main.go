package main

import (
    "go-tests/app"
    "go-tests/jobs"
)

func main() {

    client := jobs.NewClient()

    a := app.App{
        JobClient: client,
    }

    a.Submit("new job submitted")
}
