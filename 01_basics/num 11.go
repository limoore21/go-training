package main

import (
	"errors"
	"fmt"
)

type DB struct {
	users map[string]string
}

func NewDB() *DB {
	a := DB{
		users: make(map[string]string),
	}
	a.users["admin"] = "root"
	return &a
}

func (db *DB) GetRole(login string) (string, error) {
	role, isFound := db.users[login]
	if !isFound {
		return "", errors.New(login + " is not found")
	}
	return role, nil
}

func main() {
	db := NewDB()
	role, err := db.GetRole("admin")
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	fmt.Printf("[Success] Роль: %s\n", role)
	_, err = db.GetRole("hacker")
	if err != nil {
		fmt.Println("[ERROR]", err)
	}
}
