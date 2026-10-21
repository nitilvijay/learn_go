package main

import "fmt"

func (e email) cost() int {
	if !e.isSubscribed {
		return int(0.05* float64(len(e.body)))
	}
	else {
		return int(0.02* float64(len(e.body)))
	}
}

func (e email) format() string {
	// ?
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
