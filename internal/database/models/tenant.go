package models

import (
	"github.com/benjaminrae/authentication/pkg/keyify"
)

type Tenant struct {
	*Model
	Name string `gorm:"unique" json:"name"`
	Key  string `gorm:"unique" json:"key"`
	// Clients   []Client
	// Providers []Provider
}

func CreateTenant(name string) (Tenant, error) {
	key, err := keyify.Keyify(32, "t_")

	if err != nil {
		return Tenant{}, err
	}

	model := CreateModel()

	return Tenant{
		Name:  name,
		Key:   key,
		Model: &model,
	}, nil
}
