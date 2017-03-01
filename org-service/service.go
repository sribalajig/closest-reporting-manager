package service

import (
	"closest-reporting-manager/models"
)

func Create(name string) models.Org {
	org := models.Org{
		Name: name,
	}

	return org
}
