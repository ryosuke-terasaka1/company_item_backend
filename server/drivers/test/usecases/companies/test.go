package companies_test

import (
	"item_management/server/drivers/database"
	"item_management/server/entities"
)

type ComoanyRepository struct {
	database.SQLHandler
}

var (
	DetailTable = []struct {
		testname string
		id       int
		Company  entities.Company
		err      error
	}{
		{
			"success",
			30,
			entities.Company{
				ID:    30,
				Name:  "Terasaka Co",
				Email: "teasaka@test.jp",
				Image: "image",
			},
			nil,
		},
	}
)
