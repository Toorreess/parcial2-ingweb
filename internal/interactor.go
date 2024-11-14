package internal

import "parcial2-ingweb/internal/model"

type IInteractor interface {
	Create(entity *model.Entity) (*model.Entity, error)
	Get(id string) (*model.Entity, error)
	Update(id string, updates map[string]interface{}) (*model.Entity, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error)
}

type interactor struct {
	Repository IRepository
	Presenter  IPresenter
}

func NewInteractor(r IRepository, p IPresenter) IInteractor {
	return &interactor{
		Repository: r,
		Presenter:  p,
	}
}

func (i *interactor) Create(entity *model.Entity) (*model.Entity, error) {
	result, err := i.Repository.Create(entity)
	if err != nil {
		return nil, err
	}
	return i.Presenter.ResponseEntity(result), nil
}

func (i *interactor) Get(id string) (*model.Entity, error) {
	entity, err := i.Repository.Get(id)
	if err != nil {
		return nil, err
	}
	return i.Presenter.ResponseEntity(entity), nil
}

func (i *interactor) Update(id string, updates map[string]interface{}) (*model.Entity, error) {
	entity, err := i.Repository.Update(id, updates)
	if err != nil {
		return nil, err
	}
	return i.Presenter.ResponseEntity(entity), nil
}

func (i *interactor) Delete(id string) error {
	return i.Repository.Delete(id)
}

func (i *interactor) List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error) {
	result, err := i.Repository.List(query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}

	return i.Presenter.ResponseEntities(result, limit, offset), nil
}
