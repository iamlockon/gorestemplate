package example_service

import (
	"encoding/json"

	"github.com/iamlockon/gorestemplate/models"
	"github.com/iamlockon/gorestemplate/pkg/gredis"
	"github.com/iamlockon/gorestemplate/pkg/logging"
	"github.com/iamlockon/gorestemplate/service/cache_service"
)

type Example struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (e *Example) ExistByID() (bool, error) {
	return models.ExistExampleByID(e.ID)
}

func (e *Example) ExistByName() (bool, error) {
	return models.ExistExamplByName(e.Name)
}

func (e *Example) Add() error {
	return models.AddExample(e.Name, e.State, e.CreatedBy)
}

func (e *Example) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = e.ModifiedBy
	data["name"] = e.Name
	if e.State >= 0 {
		data["state"] = e.State
	}

	return models.EditExample(e.ID, data)
}

func (e *Example) Delete() error {
	return models.DeleteExample(e.ID)
}

func (e *Example) Count() (int, error) {
	return models.GetExampleTotal(e.getMaps())
}

func (e *Example) GetAll() ([]models.Example, error) {
	var (
		examples, cacheExamples []models.Example
	)

	cache := cache_service.Example{
		State: e.State,

		PageNum:  e.PageNum,
		PageSize: e.PageSize,
	}
	key := cache.GetExamplesKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheExamples)
			return cacheExamples, nil
		}
	}

	examples, err := models.GetExamples(e.PageNum, e.PageSize, e.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, examples, 3600)
	return examples, nil
}

func (e *Example) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if e.Name != "" {
		maps["name"] = e.Name
	}
	if e.State >= 0 {
		maps["state"] = e.State
	}

	return maps
}
