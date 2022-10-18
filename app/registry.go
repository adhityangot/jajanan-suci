package app

import "github.com/adhityangot/js-be/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.UserCashier{}},
		{Model: models.Products{}},
	}
}
