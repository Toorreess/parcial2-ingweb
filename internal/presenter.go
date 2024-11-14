package internal

import (
	"encoding/json"
	"parcial2-ingweb/internal/model"
)

type IPresenter interface {
	ResponseEntity(entityMap map[string]interface{}) *model.Entity
	ResponseEntities(entityList []map[string]interface{}, limit, offset int) map[string]interface{}
}

type presenter struct{}

func NewPresenter() IPresenter {
	return &presenter{}
}

func (p *presenter) ResponseEntity(entityMap map[string]interface{}) *model.Entity {
	jsonBody, _ := json.Marshal(entityMap)
	entity := model.Entity{}
	json.Unmarshal(jsonBody, &entity)
	return &entity
}

func (p *presenter) ResponseEntities(entityList []map[string]interface{}, limit, offset int) map[string]interface{} {
	resultMap := make(map[string]interface{})
	var results []*model.Entity

	for _, entityMap := range entityList {
		jsonBody, _ := json.Marshal(entityMap)
		entity := model.Entity{}
		json.Unmarshal(jsonBody, &entity)
		results = append(results, &entity)
	}

	resultMap["items"] = results
	if len(results) == limit {
		resultMap["next_offset"] = offset + limit
	} else {
		resultMap["next_offset"] = nil
	}
	if (offset - limit) < 0 {
		resultMap["previous_offset"] = nil
	} else {
		resultMap["previous_offset"] = offset - limit
	}

	resultMap["offset"] = offset
	resultMap["limit"] = limit
	resultMap["total"] = len(results)
	return resultMap
}
