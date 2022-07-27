package mysql

import (
    "testing"
)

func TestConnectByGorm(t *testing.T) {
    db, err := ConnectByGorm()
    defer db.Close()
    if err != nil {
        t.Error(err)
    }
    err = db.Exec("SELECT CURRENT_TIMESTAMP();").Error
    if err != nil {
        t.Error(err)
    }
}

func TestConnectByRaw(t *testing.T) {
    db, err := ConnectByRaw()
    defer db.Close()
    if err != nil {
        t.Error(err)
    }
    _, err = db.Exec("SELECT CURRENT_TIMESTAMP();")
    if err != nil {
        t.Error(err)
    }
}
