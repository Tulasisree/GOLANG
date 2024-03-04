package models

import (
	"errors"
	"fmt"
)

//fields that start with lower case are package internal
type User struct {
	Id        int
	FirstName string
	LastName  string
}

//var block
var (
	users  []*User //slice of pointers to user
	nextID = 1     //at apackage level u don't need colon
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.Id != 0 {
		return User{}, errors.New(" new user must not include id")
	}
	u.Id = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}
	//Errorf allows us to return formmated string
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, v := range users {
		if v.Id == u.Id {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with Id '%v' not found", u.Id)
}

func RemoveUserById(id int) error {
	for i, v := range users {
		if v.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID %v not found", id)
}
