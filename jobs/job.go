package jobs

import (
    "log"
)

type Job struct {
    ID string
    Status string
}

type Client struct {}

func NewClient() *Client {
    return &Client{}
}

func (self *Client) AddToQueue(data string) error {
    log.Println("Job added to queue")
    return nil
}

func (self *Client) GetListFromDB() error {
    return nil
}

func (self *Client) GetStatusFromDB(id string) error {
    return nil
}

func (self *Client) RemoveFromDB(id string) error {
    return nil
}
