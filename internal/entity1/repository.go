package entity1

import (
	"parcial2-ingweb/internal/database"
	"parcial2-ingweb/internal/models"
)

type IE1Repository interface {
	Create(entity *models.Entity1) (map[string]interface{}, error)
	Get(id string) (map[string]interface{}, error)
	Update(id string, updates map[string]interface{}) (map[string]interface{}, error)
	Delete(id string) error

	List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
}

type e1Repository struct {
	db *database.Connection
}

func NewE1Repository(db *database.Connection) IE1Repository {
	return &e1Repository{db: db}
}

const ENTITY_INDEX_NAME = "Wiki"

func (er *e1Repository) Create(entity *models.Entity1) (map[string]interface{}, error) {
	emr, err := er.db.Create(ENTITY_INDEX_NAME, entity)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *e1Repository) Get(id string) (map[string]interface{}, error) {
	emr, err := er.db.Get(ENTITY_INDEX_NAME, id, models.Entity1{})
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *e1Repository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	emr, err := er.db.Update(ENTITY_INDEX_NAME, id, models.Entity1{}, updates)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *e1Repository) Delete(id string) error {
	err := er.db.Delete(ENTITY_INDEX_NAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (er *e1Repository) List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error) {
	emr, err := er.db.List(ENTITY_INDEX_NAME, query, limit, offset, orderBy, order, models.Entity1{})
	if err != nil {
		return nil, err
	}
	return emr, nil
}
