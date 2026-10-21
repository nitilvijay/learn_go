package main

import (
	"fmt"
	"errors"
)

// Complete the deleteIfNecessary function.

// If the user doesn't exist in the map, return
// the error not found.

// If they exist but aren't scheduled for
// deletion, return deleted as false with
// no errors.

// If they exist and are scheduled for deletion,
// return deleted as true with no errors
// and delete their record from the map.

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	if _, ok := users[name]; !ok{
		return false, errors.New("Name not found")
	} else if !users[name].scheduledForDeletion {
		return false, nil
	} else {
		return true, nil
	}
}

type user struct {
	name string
	number int
	scheduledForDeletion bool
}

func main() {
	users := map[string]user{
		"alice": {name: "alice", number: 1, scheduledForDeletion: true},
		"bob":   {name: "bob", number: 2, scheduledForDeletion: false},
	}

	deleted, err := deleteIfNecessary(users, "alice")
	fmt.Println(deleted, err) // should print true <nil>

	deleted, err = deleteIfNecessary(users, "bob")
	fmt.Println(deleted, err) // should print false <nil>

	deleted, err = deleteIfNecessary(users, "charlie")
	fmt.Println(deleted, err) // should print false Name not found
}