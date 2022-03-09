package repositories

import (
	"fmt"
	"usms/db/models"
)

type UserAttributeRepository interface {
	GetUserAttribute(userId int) (*models.UserAttribute, error)
}

func (repo defaultRepository) GetUserAttribute(userId int) (*models.UserAttribute, error) {
	var entity models.UserAttribute
	err := repo.getById(&entity, userId)
	if err != nil {
		logDBErr(err, fmt.Sprintf("get by id %v", userId))
		return nil, err
	}
	return &entity, nil
}
