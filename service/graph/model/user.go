package model

import (
	"fmt"

	"github.com/wwwwshwww/spot-sandbox/entity"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserFromEntity(e *entity.User) *User {
	return &User{
		ID:   fmt.Sprintf("%d", e.ID),
		Name: e.Name,
	}
}
