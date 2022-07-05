package user

type User struct {
    id int
    Name string
}

func (u *User) Save() error {
    if u.id == 0 {
        return u.create()
    }
    return u.update()
}

func (u *User) create() error {
    return nil
}

func (u *User) update() error {
    return nil
}
