package repositories

import (
	"usms/db"
	"usms/db/models"
)

type defaultRepository struct {
}

//FilterField is used for where clause
type FilterField struct {
	Column    string
	Value     interface{}
	And       interface{}
	Or        interface{}
	condition string // =, !=, >=, <=, >, <
}

var defaultRepo defaultRepository

func New() *defaultRepository {
	return &defaultRepo
}

//save
func (repo *defaultRepository) save(data interface {
	models.ModelCredential
	models.ModelMetadata
}) error {
	return db.Save(data.GetTableName(), data)
}

//getById
func (repo *defaultRepository) getById(data interface {
	models.ModelMetadata
	models.ModelCache
}, id int) error {
	return db.GetById(id, data.GetTableName(), data)
}

//update
func (repo *defaultRepository) update(data interface {
	models.ModelCredential
	models.ModelMetadata
}, filter ...FilterField) error {
	var conditions map[string]interface{}
	if len(filter) > 0 {
		for _, val := range filter {
			key := val.Column
			conditions[key] = val.Value
		}

	}
	return db.Update(data.GetTableName(), data, conditions)
}

//filter
func (repo *defaultRepository) filter(data interface {
	models.ModelMetadata
	models.ModelCache
}, entities interface{}, filter ...FilterField) error {
	var conditions = make(map[string]interface{})
	if len(filter) > 0 {
		for _, val := range filter {
			key := val.Column
			conditions[key] = val.Value
		}
	}
	return db.Filter(data.GetTableName(), entities, conditions)
}

func logDBErr(err error, template string) {

}
