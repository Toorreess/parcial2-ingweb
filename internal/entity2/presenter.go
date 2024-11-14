package entity2

import (
	"encoding/json"
	"parcial2-ingweb/internal/models"
)

type IE2Presenter interface {
	ResponseEntity(entityMap map[string]interface{}) *models.Entity2
	ResponseEntities(entityList []map[string]interface{}, limit, offset int) map[string]interface{}
}

type e2Presenter struct{}

func NewE2Presenter() IE2Presenter {
	return &e2Presenter{}
}

func (ep *e2Presenter) ResponseEntity(entityMap map[string]interface{}) *models.Entity2 {
	jsonBody, _ := json.Marshal(entityMap)
	entity := models.Entity2{}
	json.Unmarshal(jsonBody, &entity)
	return &entity
}

func (ep *e2Presenter) ResponseEntities(entityList []map[string]interface{}, limit, offset int) map[string]interface{} {
	resultMap := make(map[string]interface{})
	var results []*models.Entity2

	for _, entityMap := range entityList {
		jsonBody, _ := json.Marshal(entityMap)
		entity := models.Entity2{}
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
