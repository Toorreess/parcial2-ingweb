package entity1

import "parcial2-ingweb/internal/models"

type IE1Interactor interface {
	Create(entity *models.Entity1) (*models.Entity1, error)
	Get(id string) (*models.Entity1, error)
	Update(id string, updates map[string]interface{}) (*models.Entity1, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error)
}

type e1Interactor struct {
	E1Repository IE1Repository
	E1Presenter  IE1Presenter
}

func NewE1Interactor(r IE1Repository, p IE1Presenter) IE1Interactor {
	return &e1Interactor{
		E1Repository: r,
		E1Presenter:  p,
	}
}

func (ei *e1Interactor) Create(entity *models.Entity1) (*models.Entity1, error) {
	result, err := ei.E1Repository.Create(entity)
	if err != nil {
		return nil, err
	}
	return ei.E1Presenter.ResponseEntity(result), nil
}

func (ei *e1Interactor) Get(id string) (*models.Entity1, error) {
	entity, err := ei.E1Repository.Get(id)
	if err != nil {
		return nil, err
	}
	return ei.E1Presenter.ResponseEntity(entity), nil
}

func (ei *e1Interactor) Update(id string, updates map[string]interface{}) (*models.Entity1, error) {
	entity, err := ei.E1Repository.Update(id, updates)
	if err != nil {
		return nil, err
	}
	return ei.E1Presenter.ResponseEntity(entity), nil
}

func (ei *e1Interactor) Delete(id string) error {
	return ei.E1Repository.Delete(id)
}

func (ei *e1Interactor) List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error) {
	result, err := ei.E1Repository.List(query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}

	return ei.E1Presenter.ResponseEntities(result, limit, offset), nil
}
