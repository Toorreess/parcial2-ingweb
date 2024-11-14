package entity2

import (
	"parcial2-ingweb/internal/database"
	"parcial2-ingweb/internal/models"
)

type IE2Repository interface {
	Create(entity *models.Entity2) (map[string]interface{}, error)
	Get(id string) (map[string]interface{}, error)
	Update(id string, updates map[string]interface{}) (map[string]interface{}, error)
	Delete(id string) error

	List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
}

type e2Repository struct {
	db *database.Connection
}

func NewE2Repository(db *database.Connection) IE2Repository {
	return &e2Repository{db: db}
}

const ENTITY_INDEX_NAME = "Entry"

func (er *e2Repository) Create(entity *models.Entity2) (map[string]interface{}, error) {
	emr, err := er.db.Create(ENTITY_INDEX_NAME, entity)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *e2Repository) Get(id string) (map[string]interface{}, error) {
	emr, err := er.db.Get(ENTITY_INDEX_NAME, id, models.Entity1{})
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *e2Repository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	emr, err := er.db.Update(ENTITY_INDEX_NAME, id, models.Entity1{}, updates)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *e2Repository) Delete(id string) error {
	err := er.db.Delete(ENTITY_INDEX_NAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (er *e2Repository) List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error) {
	emr, err := er.db.List(ENTITY_INDEX_NAME, query, limit, offset, orderBy, order, models.Entity1{})
	if err != nil {
		return nil, err
	}
	return emr, nil
}
