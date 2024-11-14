package entity1

import (
	"encoding/json"
	"parcial2-ingweb/internal/models"
)

type IE1Presenter interface {
	ResponseEntity(entityMap map[string]interface{}) *models.Entity1
	ResponseEntities(entityList []map[string]interface{}, limit, offset int) map[string]interface{}
}

type e1Presenter struct{}

func NewE1Presenter() IE1Presenter {
	return &e1Presenter{}
}

func (ep *e1Presenter) ResponseEntity(entityMap map[string]interface{}) *models.Entity1 {
	jsonBody, _ := json.Marshal(entityMap)
	entity := models.Entity1{}
	json.Unmarshal(jsonBody, &entity)
	return &entity
}

func (ep *e1Presenter) ResponseEntities(entityList []map[string]interface{}, limit, offset int) map[string]interface{} {
	resultMap := make(map[string]interface{})
	var results []*models.Entity1

	for _, entityMap := range entityList {
		jsonBody, _ := json.Marshal(entityMap)
		entity := models.Entity1{}
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
