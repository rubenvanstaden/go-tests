package user_test

import (
    "testing"

    . "go-tests/utils"
    . "go-tests/user"
)

func TestUser_Save(t *testing.T) {
    u := &User{Name: "Susy Queue"}
    Ok(t, u.Save())
}
