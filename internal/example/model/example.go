package model

import "github.com/charmingruby/pack/pkg/core"

type Example struct {
	ID   string `json:"id"   db:"id"`
	Name string `json:"name" db:"name"`
}

func NewExample(name string) Example {
	return Example{
		ID:   core.NewID(),
		Name: name,
	}
}
