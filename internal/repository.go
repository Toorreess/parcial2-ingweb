package internal

import (
	"parcial2-ingweb/internal/database"
	"parcial2-ingweb/internal/model"
)

type IRepository interface {
	Create(entity *model.Entity) (map[string]interface{}, error)
	Get(id string) (map[string]interface{}, error)
	Update(id string, updates map[string]interface{}) (map[string]interface{}, error)
	Delete(id string) error

	List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
}

type repository struct {
	db *database.Connection
}

func NewRepository(db *database.Connection) IRepository {
	return &repository{db: db}
}

const ENTITY_INDEX_NAME = "Wiki"

func (r *repository) Create(entity *model.Entity) (map[string]interface{}, error) {
	emr, err := r.db.Create(ENTITY_INDEX_NAME, entity)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (r *repository) Get(id string) (map[string]interface{}, error) {
	emr, err := r.db.Get(ENTITY_INDEX_NAME, id, model.Entity{})
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (r *repository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	emr, err := r.db.Update(ENTITY_INDEX_NAME, id, model.Entity{}, updates)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (r *repository) Delete(id string) error {
	err := r.db.Delete(ENTITY_INDEX_NAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error) {
	emr, err := r.db.List(ENTITY_INDEX_NAME, query, limit, offset, orderBy, order, model.Entity{})
	if err != nil {
		return nil, err
	}
	return emr, nil
}
