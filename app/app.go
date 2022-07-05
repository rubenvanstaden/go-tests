package app

import (
    "log"
)

type App struct {
    JobClient interface {
        AddToQueue(string) error
    }
}

func (self *App) Submit(data string) error {
    log.Println("Job submitted")
    return self.JobClient.AddToQueue(data)
}

// func (self *App) List() error {
//     return nil
// }
//
// func (self *App) Status(id string) error {
//     return nil
// }
//
// func (self *App) Delete(id string) error {
//     return nil
// }
