package entity2

import "parcial2-ingweb/internal/models"

type IE2Interactor interface {
	Create(entity *models.Entity2) (*models.Entity2, error)
	Get(id string) (*models.Entity2, error)
	Update(id string, updates map[string]interface{}) (*models.Entity2, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error)
}

type e2Interactor struct {
	E2Repository IE2Repository
	E2Presenter  IE2Presenter
}

func NewE2Interactor(r IE2Repository, p IE2Presenter) IE2Interactor {
	return &e2Interactor{
		E2Repository: r,
		E2Presenter:  p,
	}
}

func (ei *e2Interactor) Create(entity *models.Entity2) (*models.Entity2, error) {
	result, err := ei.E2Repository.Create(entity)
	if err != nil {
		return nil, err
	}
	return ei.E2Presenter.ResponseEntity(result), nil
}

func (ei *e2Interactor) Get(id string) (*models.Entity2, error) {
	entity, err := ei.E2Repository.Get(id)
	if err != nil {
		return nil, err
	}
	return ei.E2Presenter.ResponseEntity(entity), nil
}

func (ei *e2Interactor) Update(id string, updates map[string]interface{}) (*models.Entity2, error) {
	entity, err := ei.E2Repository.Update(id, updates)
	if err != nil {
		return nil, err
	}
	return ei.E2Presenter.ResponseEntity(entity), nil
}

func (ei *e2Interactor) Delete(id string) error {
	return ei.E2Repository.Delete(id)
}

func (ei *e2Interactor) List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error) {
	result, err := ei.E2Repository.List(query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}

	return ei.E2Presenter.ResponseEntities(result, limit, offset), nil
}
