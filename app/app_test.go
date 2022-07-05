package app_test

import (
    "testing"

    . "go-tests/utils"
    . "go-tests/app"
)

type TestJobClient struct {
    AddToQueueFunc func(string) error
}

func (self *TestJobClient) AddToQueue(data string) error {
    return self.AddToQueueFunc(data)
}

func TestMyApplication_SendYo(t *testing.T) {

    c := &TestJobClient{}

    // Mock our send function to capture the argument.
    var data string
    c.AddToQueueFunc = func(s string) error {
        data = s
        return nil
    }

    a := &App{JobClient: c}

    err := a.Submit("job submitted")
    Ok(t, err)
    Equals(t, "job submitted", data)
}
